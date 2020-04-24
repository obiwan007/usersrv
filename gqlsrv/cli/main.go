package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/namsral/flag"

	"github.com/coreos/etcd/clientv3"
	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/cors"

	"github.com/common-nighthawk/go-figure"
	graphql "github.com/graph-gophers/graphql-go"
	gql "github.com/obiwan007/usersrv/gqlsrv/api"
	pb "github.com/obiwan007/usersrv/proto"
	"github.com/obiwan007/usersrv/usersrv/api/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	config             = flag.String("config", "", "Using configfile")
	balancer           = flag.Bool("ectd", false, "Using etcd")
	tlsForGrpc         = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	gClientID          = flag.String("clientid", "xxxxxxxxxxxxxx", "OAuth2 client Id")
	gClientSecr        = flag.String("clientsecret", "yyyyyyyyyyyyy", "OAuth2 client secret")
	redirect           = flag.String("redirect", "http://localhost:3000/auth/callback", "OAuth2 client redirect callback")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "usersrv:10000", "The server address in the format of host:port")
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
	// Your credentials should be obtained from the Google
	// Developer Console (https://console.developers.google.com).

	myFigure := figure.NewFigure("GQLSRV", "", true)
	myFigure.Print()
	flag.Parse()
	fmt.Println("Init GraphQL Service")
	s, err := getSchema("../schema/schema.graphql")
	if err != nil {
		panic(err)
	}

	tlsCfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

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

	_, collector, err := tracing.NewTracer("gql service", "localhost:10000", *zipkin)
	if err != nil {
		panic(err)
	}
	log.Println("Logging to Zipkin:", *zipkin)
	defer collector.Close()
	t := opentracing.GlobalTracer()
	opts = append(opts, grpc.WithUnaryInterceptor(
		otgrpc.OpenTracingClientInterceptor(t, otgrpc.LogPayloads())))
	opts = append(opts, grpc.WithStreamInterceptor(
		otgrpc.OpenTracingStreamClientInterceptor(t)))

	if *balancer {
		r := &etcdnaming.GRPCResolver{Client: cli}
		b := grpc.RoundRobin(r)
		opts = append(opts, grpc.WithBalancer(b))
	}

	if *tlsForGrpc {
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
	log.Println("Dial", *serverAddr)

	conn, err := grpc.Dial(*serverAddr, opts...)
	// } else {
	// 	conn, err := grpc.Dial("my-service", opts...)
	// }

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	gqlClient := pb.NewUserServiceClient(conn)

	resolver := gql.NewResolver(gqlClient)

	// schema := graphql.MustParseSchema(s, resolver, graphql.UseStringDescriptions(), graphql.Tracer(trace.OpenTracingTracer{}))
	schema := graphql.MustParseSchema(s, resolver, graphql.UseStringDescriptions())
	mux := gql.NewRouter(schema, t, gClientID, gClientSecr, redirect)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.
	// mux2 := cors.Default().Handler(mux)
	mux2 := cors.New(cors.Options{
		// AllowedOrigins:   []string{"http://foo.com", "http://foo.com:8080"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"http://localhost:3000", "https://frontend-obiwan007.cloud.okteto.net"},
		// Enable Debugging for testing, consider disabling in production
		AllowedHeaders: []string{"Authorization", "Content-Type", "X-B3-Sampled", "X-B3-Spanid", "X-B3-Traceid"},
		Debug:          true,
	}).Handler(mux)

	srv := &http.Server{
		Addr:         ":8090",
		Handler:      mux2,
		TLSConfig:    tlsCfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Listening on port 8090")
	log.Fatal(srv.ListenAndServe())
	// log.Fatal(srv.ListenAndServeTLS("server.rsa.crt", "server.rsa.key"))
}

func getSchema(path string) (string, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
