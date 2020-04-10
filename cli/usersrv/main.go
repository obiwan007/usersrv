//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/obiwan007/usersrv/api"
	"github.com/obiwan007/usersrv/api/tracing"
	pb "github.com/obiwan007/usersrv/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

// const (
// 	endpoint_url              = "http://localhost:9411/api/v1/spans"
// 	host_url                  = "localhost:5051"
// 	service_name_cache_client = "cache service client"
// 	service_name_call_get     = "callGet"
// )

// func newTracer() (opentracing.Tracer, zipkintracer.Collector, error) {
// 	collector, err := openzipkin.NewHTTPCollector(endpoint_url)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	recorder := openzipkin.NewRecorder(collector, true, host_url, service_name_cache_client)
// 	tracer, err := openzipkin.NewTracer(
// 		recorder,
// 		openzipkin.ClientServerSameSpan(true))

// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	opentracing.SetGlobalTracer(tracer)

// 	return tracer, collector, nil
// }
func main() {
	grpc.EnableTracing = true

	tracer, collector, err := tracing.NewTracer("userservice")
	if err != nil {
		panic(err)
	}
	defer collector.Close()
	// tracer := dapperish.NewTracer("dapperish_tester")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	// Add Tracer
	opts = append(opts, grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())))
	opts = append(opts, grpc.StreamInterceptor(
		otgrpc.OpenTracingStreamServerInterceptor(tracer)))

	if *tls {
		if *certFile == "" {
			log.Fatalln("No certfile")
		}
		if *keyFile == "" {
			log.Fatalln("No Keyfile")
		}

		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)

		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		log.Println("Certs loaded", *certFile, *keyFile)
		opts = []grpc.ServerOption{grpc.Creds(creds)}
		log.Printf("Credentials %v", opts)
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, api.NewServer())
	grpcServer.Serve(lis)
}
