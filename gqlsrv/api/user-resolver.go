package gql

import (
	"context"
	"fmt"

	api "github.com/obiwan007/usersrv/proto"
	// "github.com/pkg/errors"
)

func (r *Resolver) User(ctx context.Context, args struct{ Id *string }) (*UserResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	request := &api.Id{Id: *args.Id, Jwt: token.Raw}

	res, err := r.userSvc.GetUser(ctx, request)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	s := UserResolver{
		R: &User{Name: res.Name, Email: res.Email, ID: res.Id},
	}

	return &s, nil

}
func (r *Resolver) AllUsers(ctx context.Context) (*[]*UserResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	users, err := r.userSvc.GetUsers(ctx, &api.ListUsers{Jwt: token.Raw})
	if err != nil {
		return nil, err
	}
	var userRxs []*UserResolver
	for _, res := range users.Users {
		user := &User{Name: res.Name, Email: res.Email, ID: res.Id}
		s := UserResolver{R: user}
		userRxs = append(userRxs, &s)

	}

	return &userRxs, nil

}

// type userInput struct {
// 	Name     *string
// 	Password *string
// 	Email    *string
// }

func (r *Resolver) CreateUser(ctx context.Context, args struct{ User UserInput }) (*string, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	request := &api.User{Jwt: token.Raw, Name: *args.User.Name, Password: *args.User.Password, Email: *args.User.Email}

	res, err := r.userSvc.RegisterUser(ctx, request)
	if err != nil {
		return nil, err
	}
	name := res.GetName()
	return &name, nil
}
