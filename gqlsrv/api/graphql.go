package gql

import (
	"context"
	"fmt"

	api "github.com/obiwan007/usersrv/proto"
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

func (r *Resolver) GetUser(ctx context.Context, args struct{ Id *string }) (*UserResolver, error) {

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

type userInput struct {
	Name     *string
	Password *string
	Email    *string
}

func (r *Resolver) RegisterUser(ctx context.Context, args struct{ User userInput }) (*string, error) {

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
