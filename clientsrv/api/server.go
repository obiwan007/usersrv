//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"fmt"

	storage "github.com/obiwan007/usersrv/clientsrv/api/storage"
	pb "github.com/obiwan007/usersrv/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var sto storage.FileStorage

type routeGuideServer struct {
	storage storage.FileStorage
}

func Init(storage storage.FileStorage) {
	fmt.Println("Init Project Service")
	sto = storage
}

func (s *routeGuideServer) Add(ctx context.Context, timer *pb.Client) (*pb.Client, error) {
	fmt.Println("ADDING Client", timer.Description)
	// No feature was found, return an unnamed feature
	newuser := s.storage.Add(*timer)
	return &newuser, nil
}
func (s *routeGuideServer) Update(ctx context.Context, timer *pb.Client) (*pb.Client, error) {
	fmt.Println("ADDING Client", timer.Description)
	// No feature was found, return an unnamed feature
	newuser := s.storage.Update(*timer)
	return &newuser, nil
}
func (s *routeGuideServer) Del(ctx context.Context, timerId *pb.Id) (*pb.Client, error) {
	fmt.Println("Deleting Client", timerId.GetId())
	// No feature was found, return an unnamed feature
	deleted, err := s.storage.Delete(timerId.GetId())
	return deleted, err
}

func (s *routeGuideServer) Get(ctx context.Context, timerId *pb.Id) (*pb.Client, error) {
	fmt.Println("Get Client", timerId.GetId())
	// No feature was found, return an unnamed feature
	newtimer, err := s.storage.Get(timerId.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Client not found")
	}
	return newtimer, nil
}

func (s *routeGuideServer) GetAll(ctx context.Context, l *pb.ListClient) (*pb.ClientResponse, error) {
	fmt.Println("Get Clients")
	timers := s.storage.GetAll()
	u := new(pb.ClientResponse)

	for _, timer := range timers {
		conv := timer
		u.Clients = append(u.Clients, conv)
	}
	return u, nil
}

// func AddUser(name string, password string) {
// 	fmt.Println("Add User", name, password)
// 	user := storage.User{Name: name, Password: password}
// 	fmt.Println("Added User", user)
// 	sto.AddUser(user)
// }

func NewServer() *routeGuideServer {
	fs := storage.NewFileStorage()
	s := &routeGuideServer{storage: *fs}
	return s
}

// func toAPITimer(newtimer storage.Timer) *pb.Timer {
// 	return &pb.Timer{Description: newtimer.Description, Id: newtimer.Id}
// }

// func fromAPITimer(newtimer pb.Timer) *storage.Timer {
// 	return &storage.Timer{Description: newtimer.Description, Id: newtimer.Id}
// }
