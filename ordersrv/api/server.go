//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	pb "github.com/obiwan007/usersrv/proto"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

const (
	event     = "order-created"
	aggregate = "order"
)

type orderServiceType struct {
	eventstore pb.EventStoreClient
}

func NewServer(eventstore pb.EventStoreClient) *orderServiceType {
	s := &orderServiceType{eventstore: eventstore}
	return s
}

func (s *orderServiceType) CreateOrder(ctx context.Context, o *pb.OrderRequest) (*pb.OrderResponse, error) {
	fmt.Println("ADDING USER", o.GetCustomer(), o.GetStatus())
	// No feature was found, return an unnamed feature
	var order pb.OrderCreateCommand
	aggregateID := uuid.NewV4().String()
	order.OrderId = aggregateID
	order.Status = "Pending"
	order.CreatedOn = time.Now().Unix()
	err := s.createOrderRPC(order)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	res := &pb.OrderResponse{}
	return res, nil
}

// createOrderRPC calls the CreateEvent RPC
func (s *orderServiceType) createOrderRPC(order pb.OrderCreateCommand) error {
	log.Println("Submitting to store")
	orderJSON, _ := json.Marshal(order)

	event := &pb.Event{
		EventId:       uuid.NewV4().String(),
		EventType:     event,
		AggregateId:   order.OrderId,
		AggregateType: aggregate,
		EventData:     string(orderJSON),
		Channel:       event,
	}

	resp, err := s.eventstore.CreateEvent(context.Background(), event)
	if err != nil {
		return errors.Wrap(err, "Error from RPC server")
	}
	if resp.IsSuccess {
		return nil
	} else {
		return errors.Wrap(err, "Error from RPC server")
	}
}

// func toAPIUser(newuser storage.User) *pb.User {
// 	return &pb.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}
// }

// func fromAPIUser(newuser pb.User) *storage.User {
// 	return &storage.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}
// }
