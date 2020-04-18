package gql

import (
	"context"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	api "github.com/obiwan007/usersrv/proto"

	// "github.com/pkg/errors"
	errors "github.com/graph-gophers/graphql-go/errors"
)

type Resolver struct {
	userSvc api.UserServiceClient
}

func NewResolver(helloClient api.UserServiceClient) *Resolver {
	return &Resolver{
		userSvc: helloClient,
	}
}

func (r *Resolver) Hello(ctx context.Context, args struct{ Name string }) (string, error) {

	request := &api.User{Name: "Markus", Password: "test"}

	res, err := r.userSvc.AddUser(ctx, request)
	if err != nil {
		return "", err
	}

	return res.GetName(), nil
}

type userResponse struct {
	Name  string
	Email string
}

type User struct {
	Id    string
	Name  string
	Email string
}

func (r *Resolver) User(ctx context.Context, args struct{ Id *string }) (*UserResolver, error) {

	request := &api.Id{Id: *args.Id}

	res, err := r.userSvc.GetUser(ctx, request)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	s := UserResolver{
		u: &User{Name: res.Name, Email: res.Email, Id: res.Id},
	}

	return &s, nil

}
func (r *Resolver) AllUsers(ctx context.Context) (*[]*UserResolver, error) {
	t := ctx.Value("jwt")

	token, ok := t.(*jwt.Token)
	if !ok || !token.Valid {

		return nil, errors.Errorf("Unauthorized")

		// return nil, fmt.Errorf("Invalid Token")
	}

	fmt.Println(t)
	claims := token.Claims.(*MyCustomClaims)
	fmt.Println("Email Claim:", claims.Email)
	users, err := r.userSvc.GetUsers(ctx, &api.ListUsers{})
	if err != nil {
		return nil, err
	}
	var userRxs []*UserResolver
	for _, res := range users.Users {
		user := &User{Name: res.Name, Email: res.Email, Id: res.Id}
		s := UserResolver{u: user}
		userRxs = append(userRxs, &s)

	}

	return &userRxs, nil

}

type userInput struct {
	Name     *string
	Password *string
	Email    *string
}

func (r *Resolver) CreateUser(ctx context.Context, args struct{ User userInput }) (*string, error) {

	request := &api.User{Name: *args.User.Name, Password: *args.User.Password, Email: *args.User.Email}

	res, err := r.userSvc.RegisterUser(ctx, request)
	if err != nil {
		return nil, err
	}
	name := res.GetName()
	return &name, nil
}

/*
 * UserResolver
 */

type UserResolver struct{ u *User }

func (r *UserResolver) Name() *string {
	return &r.u.Name
}

func (r *UserResolver) Email() *string {
	return &r.u.Email
}

func (r *UserResolver) Id() *string {
	return &r.u.Id
}
