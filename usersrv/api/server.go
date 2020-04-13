//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"fmt"

	pb "github.com/obiwan007/usersrv/proto"
	storage "github.com/obiwan007/usersrv/usersrv/api/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	return toAPIUser(newuser), nil
}

func (s *routeGuideServer) GetUser(ctx context.Context, user *pb.Id) (*pb.User, error) {
	fmt.Println("Get USER", user.GetId())
	// No feature was found, return an unnamed feature
	newuser, err := s.storage.GetUser(user.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "User not found")
	}
	return toAPIUser(newuser), nil
}

func (s *routeGuideServer) RegisterUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	fmt.Println("Check for user with email", user.GetEmail(), user.GetId())
	// No feature was found, return an unnamed feature
	userExisting, err := s.storage.GetUserFromEmail(user.GetEmail())
	fmt.Println("Existing User:", userExisting)
	if err != nil {
		// User is not existing, add it
		fmt.Println("User Checked, not existing")
		newuser := s.storage.AddUser(storage.User{Name: user.GetName(), Email: user.GetEmail(), Password: user.GetPassword()})
		return toAPIUser(newuser), nil
	}

	return nil, status.Error(codes.InvalidArgument, "User already existing")
}

func (s *routeGuideServer) CheckUser(ctx context.Context, mail *pb.Email) (*pb.User, error) {
	fmt.Println("Check for user with email", mail.GetEmail())
	// No feature was found, return an unnamed feature
	newuser, err := s.storage.GetUserFromEmail(mail.GetEmail())

	if err != nil {
		return &pb.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}, nil
	}
	return toAPIUser(newuser), nil
}

func (s *routeGuideServer) GetUsers(ctx context.Context, l *pb.ListUsers) (*pb.UsersResponse, error) {
	fmt.Println("Get USERS")
	users := s.storage.ListUser()
	u := new(pb.UsersResponse)
	// var u = pb.UsersResponse.Users

	for _, user := range users {
		conv := toAPIUser(user)
		u.Users = append(u.Users, conv)
		// fmt.Println(u.Users)
	}
	// fmt.Println(u.Users)
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

func toAPIUser(newuser storage.User) *pb.User {
	return &pb.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}
}

func fromAPIUser(newuser pb.User) *storage.User {
	return &storage.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}
}
