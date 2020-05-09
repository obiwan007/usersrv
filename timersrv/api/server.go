//go:generate protoc -I ../proto/ ../proto/user.proto --go_out=plugins=grpc:../proto

package api

import (
	"context"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	claims "github.com/obiwan007/usersrv/pkg/claims"
	pb "github.com/obiwan007/usersrv/proto"
	storage "github.com/obiwan007/usersrv/timersrv/api/storage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var sto storage.FileStorage

type routeGuideServer struct {
	Storage    storage.FileStorage
	signingKey []byte
}

func Init(storage storage.FileStorage) {
	fmt.Println("Init User Service")
	sto = storage
}

func (s *routeGuideServer) Add(ctx context.Context, timer *pb.Timer) (*pb.Timer, error) {
	fmt.Println("ADDING TIMER", timer.Description)
	// No feature was found, return an unnamed feature
	c, err := s.getClaims(timer.Jwt)
	if err != nil {
		return nil, err
	}
	newuser, err := s.Storage.Add(ctx, timer, c)
	return newuser, err
}
func (s *routeGuideServer) Update(ctx context.Context, timer *pb.Timer) (*pb.Timer, error) {
	fmt.Println("ADDING TIMER", timer.Description)
	// No feature was found, return an unnamed feature
	newuser := s.Storage.Update(*timer)
	return newuser, nil
}
func (s *routeGuideServer) Del(ctx context.Context, timerId *pb.Id) (*pb.Timer, error) {
	fmt.Println("Deleting TIMER", timerId.GetId())
	// No feature was found, return an unnamed feature
	s.Storage.Delete(timerId.GetId())
	return nil, nil
}

func (s *routeGuideServer) Get(ctx context.Context, timerId *pb.Id) (*pb.Timer, error) {
	fmt.Println("Get TIMER", timerId.GetId())
	// No feature was found, return an unnamed feature
	newtimer, err := s.Storage.Get(timerId.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "User not found")
	}
	return newtimer, nil
}

func (s *routeGuideServer) Start(ctx context.Context, timerId *pb.Id) (*pb.Timer, error) {
	fmt.Println("Get TIMER", timerId.GetId())
	// No feature was found, return an unnamed feature
	t, err := s.Storage.Start(timerId.GetId())
	return &t, err
}

func (s *routeGuideServer) Stop(ctx context.Context, timerId *pb.Id) (*pb.Timer, error) {
	fmt.Println("Get TIMER", timerId.GetId())
	// No feature was found, return an unnamed feature
	t, err := s.Storage.Stop(timerId.GetId())
	return &t, err
}

// func (s *routeGuideServer) RegisterUser(ctx context.Context, user *pb.Timer) (*pb.Timer, error) {
// 	fmt.Println("Check for user with email", user.GetEmail(), user.GetId())
// 	// No feature was found, return an unnamed feature
// 	userExisting, err := s.storage.GetUserFromEmail(user.GetEmail())
// 	fmt.Println("Existing User:", userExisting)
// 	if err != nil {
// 		// User is not existing, add it
// 		fmt.Println("User Checked, not existing")
// 		newuser := s.storage.AddUser(*fromAPITimer(*user))
// 		return toAPITimer(newuser), nil
// 	}

// 	return nil, status.Error(codes.InvalidArgument, "User already existing")
// }

// func (s *routeGuideServer) CheckUser(ctx context.Context, mail *pb.Email) (*pb.User, error) {
// 	fmt.Println("Check for user with email", mail.GetEmail())
// 	// No feature was found, return an unnamed feature
// 	newuser, err := s.storage.GetUserFromEmail(mail.GetEmail())

// 	if err != nil {
// 		return nil, err
// 	}
// 	return toAPITimer(newuser), nil
// }

func (s *routeGuideServer) GetAll(ctx context.Context, l *pb.ListTimer) (*pb.TimerResponse, error) {
	log.Println("Get Timers")

	token, err := jwt.ParseWithClaims(l.Jwt, &claims.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return s.signingKey, nil
	})
	// c, ok := token.Claims.(*claims.MyCustomClaims)

	if c, ok := token.Claims.(*claims.MyCustomClaims); ok && token.Valid {
		log.Printf("CLAIMS: %v %v", c.StandardClaims.Subject, c.Mandant)
	} else {
		log.Println(err)
		return nil, grpc.Errorf(grpc.Code(jwt.ValidationError{}), "Error %v", err)
	}
	// token, ok := t.(*jwt.Token)
	// log.Println("Token", token.Valid, token)
	// log.Println("Claims", token.Claims)
	timers := s.Storage.GetAll()
	u := new(pb.TimerResponse)

	for _, timer := range timers {
		conv := timer
		u.Timers = append(u.Timers, conv)
	}
	return u, nil
}

// func AddUser(name string, password string) {
// 	fmt.Println("Add User", name, password)
// 	user := storage.User{Name: name, Password: password}
// 	fmt.Println("Added User", user)
// 	sto.AddUser(user)
// }

// NewServer initiates a new server object with initialized database storage
func NewServer(signingKey []byte, dbconnection string) *routeGuideServer {
	fs := storage.NewFileStorage(dbconnection)
	s := &routeGuideServer{Storage: *fs, signingKey: signingKey}
	return s
}

func (s *routeGuideServer) getClaims(jwtstring string) (*claims.MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(jwtstring, &claims.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return s.signingKey, nil
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

// func toAPITimer(newtimer storage.Timer) *pb.Timer {
// 	return &pb.Timer{Description: newtimer.Description, Id: newtimer.Id}
// }

// func fromAPITimer(newtimer pb.Timer) *storage.Timer {
// 	return &storage.Timer{Description: newtimer.Description, Id: newtimer.Id}
// }
