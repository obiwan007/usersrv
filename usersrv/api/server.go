//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	pb "github.com/obiwan007/usersrv/proto"
	storage "github.com/obiwan007/usersrv/usersrv/api/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	claims "github.com/obiwan007/usersrv/pkg/claims"
)

var sto storage.Storage

type routeGuideServer struct {
	Storage    storage.FileStorage
	SigningKey []byte
}

func Init(storage storage.Storage) {
	fmt.Println("Init User Service")
	sto = storage
}

func (s *routeGuideServer) AddUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	fmt.Println("ADDING USER", user.GetName())
	c, err := s.getClaims(user.Jwt)
	if err != nil {
		return nil, err
	}
	// No feature was found, return an unnamed feature
	newuser, err := s.Storage.AddUser(ctx, user, c)
	return newuser, nil
}

func (s *routeGuideServer) GetUser(ctx context.Context, user *pb.Id) (*pb.User, error) {
	fmt.Println("Get USER", user.GetId())
	c, err := s.getClaims(user.Jwt)
	if err != nil {
		return nil, err
	}
	// No feature was found, return an unnamed feature
	newuser, err := s.Storage.GetUser(ctx, user.GetId(), c)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "User not found")
	}
	return newuser, nil
}

func (s *routeGuideServer) RegisterUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	fmt.Println("Check for user with email", user.GetEmail(), user.GetId())
	// No feature was found, return an unnamed feature
	c, err := s.getClaims(user.Jwt)
	if err != nil {
		return nil, err
	}
	userExisting, err := s.Storage.GetUserFromEmail(ctx, user.GetEmail(), c)
	fmt.Println("Existing User:", userExisting)
	if err != nil {
		// User is not existing, add it
		fmt.Println("User Checked, not existing")
		newuser, err := s.Storage.AddUser(ctx, user, c)
		return newuser, err
	}

	return nil, status.Error(codes.InvalidArgument, "User already existing")
}

func (s *routeGuideServer) CheckUser(ctx context.Context, mail *pb.Email) (*pb.User, error) {
	fmt.Println("Check for user with email", mail.GetEmail())
	// No feature was found, return an unnamed feature
	c, err := s.getClaims(mail.Jwt)
	if err != nil {
		return nil, err
	}
	newuser, err := s.Storage.GetUserFromEmail(ctx, mail.GetEmail(), c)

	if err != nil {
		return nil, err
	}
	return newuser, nil
}

func (s *routeGuideServer) GetUsers(ctx context.Context, l *pb.ListUsers) (*pb.UsersResponse, error) {
	fmt.Println("Get Users")
	c, err := s.getClaims(l.Jwt)
	if err != nil {
		return nil, err
	}
	allEntities, err := s.Storage.ListUser(ctx, c)
	if err != nil {
		return nil, err
	}
	u := new(pb.UsersResponse)

	for _, entity := range allEntities {
		conv := entity
		u.Users = append(u.Users, conv)
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

func toAPIUser(newuser storage.User) *pb.User {
	return &pb.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}
}

func fromAPIUser(newuser pb.User) *storage.User {
	return &storage.User{Name: newuser.Name, Email: newuser.Email, Password: newuser.Password, Id: newuser.Id}
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
