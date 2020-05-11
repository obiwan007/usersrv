//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	storage "github.com/obiwan007/usersrv/clientsrv/api/storage"
	pb "github.com/obiwan007/usersrv/proto"
	"github.com/obiwan007/usersrv/timersrv/api/storage/ent/timer"
	"google.golang.org/grpc"

	claims "github.com/obiwan007/usersrv/pkg/claims"
)

var sto storage.FileStorage

type routeGuideServer struct {
	Storage    storage.FileStorage
	SigningKey []byte
}

func (s *routeGuideServer) Add(ctx context.Context, entity *pb.Client) (*pb.Client, error) {
	fmt.Println("ADDING Client", timer.Description)
	c, err := s.getClaims(entity.Jwt)
	if err != nil {
		return nil, err
	}
	newuser, err := s.Storage.Add(ctx, entity, c)
	return newuser, err
}
func (s *routeGuideServer) Update(ctx context.Context, entity *pb.Client) (*pb.Client, error) {
	fmt.Println("Update Client", entity.Description)
	c, err := s.getClaims(entity.Jwt)
	if err != nil {
		return nil, err
	}
	newEntity, err := s.Storage.Update(ctx, entity, c)
	return newEntity, err
}
func (s *routeGuideServer) Del(ctx context.Context, entityID *pb.Id) (*pb.Client, error) {
	fmt.Println("Deleting Client", entityID.GetId())
	c, err := s.getClaims(entityID.Jwt)
	if err != nil {
		return nil, err
	}
	deleted, err := s.Storage.Delete(ctx, entityID.GetId(), c)
	return deleted, err
}

func (s *routeGuideServer) Get(ctx context.Context, entityID *pb.Id) (*pb.Client, error) {
	fmt.Println("Get Client", entityID.GetId())
	c, err := s.getClaims(entityID.Jwt)
	if err != nil {
		return nil, err
	}
	newtimer, err := s.Storage.Get(ctx, entityID.GetId(), c)
	return newtimer, err
}

func (s *routeGuideServer) GetAll(ctx context.Context, l *pb.ListClient) (*pb.ClientResponse, error) {
	fmt.Println("Get Clients")
	c, err := s.getClaims(l.Jwt)
	if err != nil {
		return nil, err
	}
	allEntities, err := s.Storage.GetAll(ctx, c)
	if err != nil {
		return nil, err
	}
	u := new(pb.ClientResponse)

	for _, entity := range allEntities {
		conv := entity
		u.Clients = append(u.Clients, conv)
	}
	return u, err
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
