//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package main

import (
	"fmt"
	"log"
	"net"

	"github.com/namsral/flag"

	"github.com/common-nighthawk/go-figure"
	"github.com/obiwan007/usersrv/pkg"
	pb "github.com/obiwan007/usersrv/proto"
	"github.com/obiwan007/usersrv/timersrv/api"
	"google.golang.org/grpc"
)

var (
	tlsForGrpc = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	balancer   = flag.Bool("ectd", false, "Using etcd")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10001, "The server port")
	zipkin     = flag.String("zipkin", "http://zipkin:9411/api/v1/spans", "Zipkin URL")
)

var mySigningKey = []byte("captainjacksparrowsayshi")

func main() {
	flag.Parse()

	myFigure := figure.NewFigure("TIMERSRV", "", true)
	myFigure.Print()

	p := &pkg.CommandParams{tlsForGrpc, balancer, certFile, keyFile, port, zipkin, "timerservice"}

	opts, collector := pkg.PrepareOptsTracing(p)
	defer collector.Close()

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterTimerServiceServer(grpcServer, api.NewServer(mySigningKey))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer.Serve(lis)
}
