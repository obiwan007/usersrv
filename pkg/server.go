package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	storage "github.com/obiwan007/usersrv/pkg/storage"
	pb "github.com/obiwan007/usersrv/proto"
	"google.golang.org/grpc"
)

var sto storage.Storage

type routeGuideServer struct {
	storage storage.Storage
}

func Init(storage storage.Storage) {
	fmt.Println("Init User Service")
	sto = storage
}

func (s *routeGuideServer) AddUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	fmt.Println("ADDING USER", user.GetName(), user.GetPassword())
	// No feature was found, return an unnamed feature
	newuser := s.storage.AddUser(storage.User{Name: user.GetName(), Password: user.GetPassword()})
	return &pb.User{Name: newuser.Name, Password: newuser.Password, Id: &pb.Id{Id: int32(newuser.Id)}}, nil
}

func (s *routeGuideServer) GetUser(ctx context.Context, user *pb.Id) (*pb.User, error) {
	fmt.Println("Get USER", user.GetId())
	// No feature was found, return an unnamed feature
	newuser, err := s.storage.GetUser(int(user.GetId()))
	if err != nil {
		return &pb.User{Name: newuser.Name, Password: newuser.Password, Id: &pb.Id{Id: int32(newuser.Id)}}, nil
	}
	return &pb.User{Name: newuser.Name, Password: newuser.Password, Id: &pb.Id{Id: int32(newuser.Id)}}, nil
}

func convertToUser(newuser storage.User) *pb.User {
	return &pb.User{Name: newuser.Name, Password: newuser.Password, Id: &pb.Id{Id: int32(newuser.Id)}}
}

func (s *routeGuideServer) GetUsers(ctx context.Context, l *pb.ListUsers) (*pb.UsersResponse, error) {
	fmt.Println("Get USERS")
	users := s.storage.ListUser()
	u := new(pb.UsersResponse)
	// var u = pb.UsersResponse.Users

	for _, user := range users {
		conv := convertToUser(user)
		u.Users = append(u.Users, conv)
		fmt.Println(u.Users)
	}
	fmt.Println(u.Users)
	return u, nil
}

// func AddUser(name string, password string) {
// 	fmt.Println("Add User", name, password)
// 	user := storage.User{Name: name, Password: password}
// 	fmt.Println("Added User", user)
// 	sto.AddUser(user)
// }

func newServer() *routeGuideServer {
	s := &routeGuideServer{storage: storage.NewFileStorage()}
	return s
}

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "", "The TLS cert file")
	keyFile    = flag.String("key_file", "", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "", "A json file containing a list of features")
	port       = flag.Int("port", 10000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	// if *tls {
	// 	if *certFile == "" {
	// 		*certFile = testdata.Path("server1.pem")
	// 	}
	// 	if *keyFile == "" {
	// 		*keyFile = testdata.Path("server1.key")
	// 	}
	// 	creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	// 	if err != nil {
	// 		log.Fatalf("Failed to generate credentials %v", err)
	// 	}
	// 	opts = []grpc.ServerOption{grpc.Creds(creds)}
	// }
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterUserServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
