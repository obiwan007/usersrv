//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"fmt"
	"log"

	"github.com/obiwan007/usersrv/natsutil"
	pb "github.com/obiwan007/usersrv/proto"
	"github.com/obiwan007/usersrv/store"
)

type eventstoreServiceType struct {
	*natsutil.StreamingComponent
}

func (s *eventstoreServiceType) CreateEvent(ctx context.Context, in *pb.Event) (*pb.Response, error) {
	fmt.Println("ADDING Event", in.GetEventId())
	// Persist data into EventStore database
	command := store.EventStore{}
	// Persist events as immutable logs into CockroachDB
	err := command.CreateEvent(in)
	if err != nil {
		return nil, err
	}
	// Publish event on NATS Streaming Server
	go publishEvent(s.StreamingComponent, in)
	res := &pb.Response{IsSuccess: true}
	return res, nil
}

// GetEvents RPC gets events from EventStore by given AggregateId
func (s *eventstoreServiceType) GetEvents(ctx context.Context, in *pb.EventFilter) (*pb.EventResponse, error) {
	eventStore := store.EventStore{}
	events := eventStore.GetEvents(in)
	return &pb.EventResponse{Events: events}, nil
}

// publishEvent publishes an event via NATS Streaming server
func publishEvent(component *natsutil.StreamingComponent, event *pb.Event) {
	sc := component.NATS()
	channel := event.Channel
	eventMsg := []byte(event.EventData)
	// Publish message on subject (channel)
	sc.Publish(channel, eventMsg)
	log.Println("Published message on channel: ", channel)
}

func NewServer(comp *natsutil.StreamingComponent) *eventstoreServiceType {
	s := &eventstoreServiceType{comp}
	return s
}

// func toAPIUser(newuser storage.User) *pb.User {
// 	return &pb.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}
// }

// func fromAPIUser(newuser pb.User) *storage.User {
// 	return &storage.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}
// }
