//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"fmt"

	storage "github.com/obiwan007/usersrv/projectsrv/api/storage"
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

func (s *routeGuideServer) Add(ctx context.Context, entity *pb.Project) (*pb.Project, error) {
	fmt.Println("ADDING Project", entity.Description)
	// No feature was found, return an unnamed feature
	newuser := s.storage.Add(*entity)
	return newuser, nil
}
func (s *routeGuideServer) Update(ctx context.Context, entity *pb.Project) (*pb.Project, error) {
	fmt.Println("ADDING Project", entity.Description)
	// No feature was found, return an unnamed feature
	newuser := s.storage.Update(*entity)
	return newuser, nil
}
func (s *routeGuideServer) Del(ctx context.Context, id *pb.Id) (*pb.Project, error) {
	fmt.Println("Deleting Project", id.GetId())
	// No feature was found, return an unnamed feature
	s.storage.Delete(id.GetId())
	return nil, nil
}

func (s *routeGuideServer) Get(ctx context.Context, id *pb.Id) (*pb.Project, error) {
	fmt.Println("Get Project", id.GetId())
	// No feature was found, return an unnamed feature
	newtimer, err := s.storage.Get(id.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Project not found")
	}
	return newtimer, nil
}

func (s *routeGuideServer) GetAll(ctx context.Context, l *pb.ListProject) (*pb.ProjectResponse, error) {
	fmt.Println("Get Projects")
	timers := s.storage.GetAll()
	u := new(pb.ProjectResponse)

	for _, timer := range timers {
		conv := timer
		u.Projects = append(u.Projects, conv)
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
