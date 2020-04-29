package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/namsral/flag"

	"github.com/common-nighthawk/go-figure"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"

	"github.com/obiwan007/usersrv/ordersrv/api"
	pb "github.com/obiwan007/usersrv/proto"
	"github.com/obiwan007/usersrv/usersrv/api/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	event     = "order-created"
	aggregate = "order"
	grpcUri   = "localhost:50051"
)

var (
	config = flag.String("config", "", "Using configfile")

	caFile = flag.String("ca_file", "", "The file containing the CA root cert file")

	tlsForGrpc = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	balancer   = flag.Bool("ectd", false, "Using etcd")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	storeAddr  = flag.String("store_addr", "store:10001", "The store server address in the format of host:port")
	port       = flag.Int("port", 10001, "The server port")
	zipkin     = flag.String("zipkin", "http://zipkin:9411/api/v1/spans", "Zipkin URL")
)

func main() {

	flag.Parse()

	myFigure := figure.NewFigure("ORDERSRV", "", true)
	myFigure.Print()

	fmt.Println("Init User Service")
	grpc.EnableTracing = true

	// Create the Server
	tracer, collector, err := tracing.NewTracer("orderservice", "localhost:10001", *zipkin)
	if err != nil {
		panic(err)
	}
	log.Println("Logging to Zipkin:", *zipkin)
	defer collector.Close()
	// tracer := dapperish.NewTracer("dapperish_tester")
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

	if *tlsForGrpc {
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
	conn, client := createStoreClient()
	defer conn.Close()

	pb.RegisterOrderServiceServer(grpcServer, api.NewServer(client))
	grpcServer.Serve(lis)
}

func createStoreClient() (*grpc.ClientConn, pb.EventStoreClient) {
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
	log.Println("Dial", *storeAddr)

	conn, err := grpc.Dial(*storeAddr, opts...)
	// } else {
	// 	conn, err := grpc.Dial("my-service", opts...)
	// }

	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	return conn, pb.NewEventStoreClient(conn)

}

// func initRoutes() *mux.Router {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/api/orders", createOrder).Methods("POST")
// 	return router
// }

// func createOrder(w http.ResponseWriter, r *http.Request) {
// 	var order pb.OrderCreateCommand
// 	err := json.NewDecoder(r.Body).Decode(&order)
// 	if err != nil {
// 		http.Error(w, "Invalid Order Data", 500)
// 		return
// 	}
// 	aggregateID := uuid.NewV4().String()
// 	order.OrderId = aggregateID
// 	order.Status = "Pending"
// 	order.CreatedOn = time.Now().Unix()
// 	err = createOrderRPC(order)
// 	if err != nil {
// 		log.Print(err)
// 		http.Error(w, "Failed to create Order", 500)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusCreated)
// 	j, _ := json.Marshal(order)
// 	w.Write(j)
// }

// // createOrderRPC calls the CreateEvent RPC
// func createOrderRPC(order pb.OrderCreateCommand) error {

// 	conn, err := grpc.Dial(grpcUri, grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Unable to connect: %v", err)
// 	}
// 	defer conn.Close()
// 	client := pb.NewEventStoreClient(conn)
// 	orderJSON, _ := json.Marshal(order)

// 	event := &pb.Event{
// 		EventId:       uuid.NewV4().String(),
// 		EventType:     event,
// 		AggregateId:   order.OrderId,
// 		AggregateType: aggregate,
// 		EventData:     string(orderJSON),
// 		Channel:       event,
// 	}

// 	resp, err := client.CreateEvent(context.Background(), event)
// 	if err != nil {
// 		return errors.Wrap(err, "Error from RPC server")
// 	}
// 	if resp.IsSuccess {
// 		return nil
// 	} else {
// 		return errors.Wrap(err, "Error from RPC server")
// 	}
// }
