//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"fmt"

	pb "github.com/obiwan007/usersrv/proto"
)

type orderServiceType struct {
}

// func Init() {
// 	fmt.Println("Init User Service")
// 	sto = storage
// }

func (s *orderServiceType) CreateOrder(ctx context.Context, o *pb.OrderRequest) (*pb.OrderResponse, error) {
	fmt.Println("ADDING USER", o.GetCustomer(), o.GetStatus())
	// No feature was found, return an unnamed feature
	res := &pb.OrderResponse{}
	return res, nil
}

func NewServer() *orderServiceType {
	s := &orderServiceType{}
	return s
}

// func toAPIUser(newuser storage.User) *pb.User {
// 	return &pb.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}
// }

// func fromAPIUser(newuser pb.User) *storage.User {
// 	return &storage.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}
// }
