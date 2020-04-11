package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/obiwan007/usersrv/api/tracing"
	pb "github.com/obiwan007/usersrv/proto"

	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	balancer           = flag.Bool("ectd", false, "Using etcd")
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", ":10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
	zipkin             = flag.String("zipkin", "http://zipkin:9411/api/v1/spans", "Zipkin URL")
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
	fmt.Println("Init CLI User Service")
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatalf("fail to connect to etcd: %v", err)
	}
	defer cli.Close()

	var opts []grpc.DialOption

	// tracer := dapperish.NewTracer("dapperish_tester")

	tracer, collector, err := tracing.NewTracer("userservice client", *zipkin)
	if err != nil {
		panic(err)
	}
	defer collector.Close()
	opts = append(opts, grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())))
	opts = append(opts, grpc.WithStreamInterceptor(
		otgrpc.OpenTracingStreamClientInterceptor(tracer)))

	if *balancer {
		r := &etcdnaming.GRPCResolver{Client: cli}
		b := grpc.RoundRobin(r)
		opts = append(opts, grpc.WithBalancer(b))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	flag.Parse()
	if *tls {
		if *caFile == "" {
			log.Fatalf("No TLS crt file given")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, "")
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		} else {
			log.Println("TLS enabled")
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
		log.Println("Insecure connection!")
	}
	opts = append(opts, grpc.WithTimeout(10*time.Second))

	// opts = append(opts, grpc.WithBlock())
	log.Println("Dial")

	conn, err := grpc.Dial(*serverAddr, opts...)
	// } else {
	// 	conn, err := grpc.Dial("my-service", opts...)
	// }

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	fmt.Println("Dialed ")
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	for i := 0; i < 2; i++ {
		fmt.Println("AddUser")
		user, err := client.AddUser(ctx, &pb.User{Name: "Markus", Password: "test"})
		if err != nil {
			log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
		}
		fmt.Println("Added User", user)
	}

	users, err := client.GetUsers(ctx, new(pb.ListUsers))
	fmt.Println("Users", users)
}
