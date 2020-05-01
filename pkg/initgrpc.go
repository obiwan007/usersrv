package pkg

import (
	"fmt"
	"log"
	"time"

	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/obiwan007/usersrv/usersrv/api/tracing"
	"github.com/opentracing/opentracing-go"
	zipkintracer "github.com/openzipkin/zipkin-go-opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// CommandParam gives all needed information to be able to prepare tracing
type CommandParams struct {
	TlsForGrpc *bool
	Balancer   *bool
	CertFile   *string
	KeyFile    *string
	Port       *int
	Zipkin     *string
	Tracename  string
}

// PrepareOptsTracing creates a new tracing collector, attaches it to GRPC options and returns it to be added toa grpcserver instacne
func PrepareOptsTracing(p *CommandParams) ([]grpc.ServerOption, zipkintracer.Collector) {
	grpc.EnableTracing = true
	var opts []grpc.ServerOption
	// Create the Server
	tracer, collector, err := tracing.NewTracer(p.Tracename, fmt.Sprintf("localhost:%d", *p.Port), *p.Zipkin)
	if err != nil {
		panic(err)
	}
	log.Println("Server Logging to Zipkin:", *p.Zipkin)
	// Add Tracer
	opts = append(opts, grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads())))
	opts = append(opts, grpc.StreamInterceptor(
		otgrpc.OpenTracingStreamServerInterceptor(tracer)))

	if *p.TlsForGrpc {
		if *p.CertFile == "" {
			log.Fatalln("No certfile")
		}
		if *p.KeyFile == "" {
			log.Fatalln("No Keyfile")
		}

		creds, err := credentials.NewServerTLSFromFile(*p.CertFile, *p.KeyFile)

		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		log.Println("Certs loaded", *p.CertFile, *p.KeyFile)
		opts = []grpc.ServerOption{grpc.Creds(creds)}
		log.Printf("Credentials %v", opts)
	}
	return opts, collector
}

func CreateClient(p *CommandParams, destURL string) (*grpc.ClientConn, opentracing.Tracer, error) {
	var opts []grpc.DialOption

	log.Println("Perpare grpc client")
	_, collector, err := tracing.NewTracer(p.Tracename, destURL, *p.Zipkin)
	if err != nil {
		panic(err)
	}
	log.Println("Client Logging to Zipkin:", *p.Zipkin)
	defer collector.Close()
	t := opentracing.GlobalTracer()
	opts = append(opts, grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(t, otgrpc.LogPayloads())))
	opts = append(opts, grpc.WithStreamInterceptor(
		otgrpc.OpenTracingStreamClientInterceptor(t)))

	if *p.TlsForGrpc {
		if *p.CertFile == "" {
			log.Fatalf("No TLS crt file given")
		}
		creds, err := credentials.NewClientTLSFromFile(*p.CertFile, "")
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
	log.Println("Dial", destURL)

	conn, err := grpc.Dial(destURL, opts...)
	// } else {
	// 	conn, err := grpc.Dial("my-service", opts...)
	// }

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
		return nil, nil, err
	}
	return conn, t, nil

}
