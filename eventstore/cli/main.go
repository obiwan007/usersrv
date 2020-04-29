package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/common-nighthawk/go-figure"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/nats-io/stan.go"
	"github.com/obiwan007/usersrv/natsutil"

	"github.com/obiwan007/usersrv/eventstore/api"
	pb "github.com/obiwan007/usersrv/proto"
	"github.com/obiwan007/usersrv/usersrv/api/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	event     = "order-created"
	aggregate = "order"
	grpcUri   = "localhost:50051"
	clusterID = "test-cluster"
	clientID  = "event-store-api"
)

var (
	config     = flag.String("config", "", "Using configfile")
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10002, "The server port")
	zipkin     = flag.String("zipkin", "http://zipkin:9411/api/v1/spans", "Zipkin URL")
)

func main() {

	flag.Parse()

	myFigure := figure.NewFigure("EVENTSTORE", "", true)
	myFigure.Print()

	fmt.Println("Init User Service")
	grpc.EnableTracing = true

	// Create the Server
	tracer, collector, err := tracing.NewTracer("eventstore", "localhost:10002", *zipkin)
	if err != nil {
		panic(err)
	}
	log.Println("Logging to Zipkin:", *zipkin)
	defer collector.Close()
	// tracer := dapperish.NewTracer("dapperish_tester")
	log.Println("Listening on port", *port)
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

	comp := natsutil.NewStreamingComponent(clientID)
	// Connect to NATS
	err = comp.ConnectToNATSStreaming(
		clusterID,
		stan.NatsURL(stan.DefaultNatsURL),
	)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterEventStoreServer(grpcServer, api.NewServer(comp))
	grpcServer.Serve(lis)
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
