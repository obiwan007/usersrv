//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	claims "github.com/obiwan007/usersrv/pkg/claims"
	storage "github.com/obiwan007/usersrv/projectsrv/api/storage"
	pb "github.com/obiwan007/usersrv/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var sto storage.FileStorage

type routeGuideServer struct {
	Storage    storage.FileStorage
	SigningKey []byte
}

func Init(storage storage.FileStorage) {
	fmt.Println("Init Project Service")
	sto = storage
}

func (s *routeGuideServer) Add(ctx context.Context, entity *pb.Project) (*pb.Project, error) {
	fmt.Println("ADDING Project", entity.Description)
	c, err := s.getClaims(entity.Jwt)
	if err != nil {
		return nil, err
	}
	newuser, err := s.Storage.Add(ctx, *entity, c)
	return newuser, nil
}
func (s *routeGuideServer) Update(ctx context.Context, entity *pb.Project) (*pb.Project, error) {
	fmt.Println("ADDING Project", entity.Description)
	c, err := s.getClaims(entity.Jwt)
	if err != nil {
		return nil, err
	}
	// No feature was found, return an unnamed feature
	newuser, err := s.Storage.Update(ctx, entity, c)
	return newuser, err
}
func (s *routeGuideServer) Del(ctx context.Context, id *pb.Id) (*pb.Project, error) {
	fmt.Println("Deleting Project", id.GetId())
	c, err := s.getClaims(id.Jwt)
	if err != nil {
		return nil, err
	}
	// No feature was found, return an unnamed feature
	s.Storage.Delete(ctx, id.GetId(), c)
	return nil, nil
}

func (s *routeGuideServer) Get(ctx context.Context, id *pb.Id) (*pb.Project, error) {
	fmt.Println("Get Project  with JWT", id.GetId(), id.GetJwt())
	c, err := s.getClaims(id.GetJwt())
	if err != nil {
		return nil, err
	}
	// No feature was found, return an unnamed feature
	newtimer, err := s.Storage.Get(ctx, id.GetId(), c)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Project not found")
	}
	return newtimer, nil
}

func (s *routeGuideServer) GetAll(ctx context.Context, l *pb.ListProject) (*pb.ProjectResponse, error) {
	fmt.Println("Get Projects")
	c, err := s.getClaims(l.Jwt)
	if err != nil {
		return nil, err
	}
	timers, err := s.Storage.GetAll(ctx, c)
	if err != nil {
		return nil, err
	}
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

func NewServer(signingKey []byte, dbconnection string) *routeGuideServer {
	fs := storage.NewFileStorage(dbconnection)
	s := &routeGuideServer{Storage: *fs, SigningKey: signingKey}
	return s
}

// func toAPITimer(newtimer storage.Timer) *pb.Timer {
// 	return &pb.Timer{Description: newtimer.Description, Id: newtimer.Id}
// }

// func fromAPITimer(newtimer pb.Timer) *storage.Timer {
// 	return &storage.Timer{Description: newtimer.Description, Id: newtimer.Id}
// }
func (s *routeGuideServer) getClaims(jwtstring string) (*claims.MyCustomClaims, error) {

	token, err := jwt.ParseWithClaims(jwtstring, &claims.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return s.SigningKey, nil
	})
	// c, ok := token.Claims.(*claims.MyCustomClaims)

	if c, ok := token.Claims.(*claims.MyCustomClaims); ok && token.Valid {
		log.Printf("CLAIMS: %v %v", c.StandardClaims.Subject, c.Mandant)
		return c, nil
	} else {
		log.Println(err)
		return nil, grpc.Errorf(grpc.Code(jwt.ValidationError{}), "Error %v", err)
	}

}
