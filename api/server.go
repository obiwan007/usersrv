//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"fmt"

	storage "github.com/obiwan007/usersrv/api/storage"
	pb "github.com/obiwan007/usersrv/proto"
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

func NewServer() *routeGuideServer {
	s := &routeGuideServer{storage: storage.NewFileStorage()}
	return s
}
