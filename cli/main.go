package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/clientv3"
	pb "github.com/obiwan007/usersrv/proto"

	etcdnaming "github.com/coreos/etcd/clientv3/naming"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	balancer           = flag.Bool("ectd", false, "Using etcd")
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.youtube.com", "The server name use to verify the hostname returned by TLS handshake")
)

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
