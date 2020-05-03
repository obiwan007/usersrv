package gql

import (
	"context"
	"fmt"

	"github.com/obiwan007/usersrv/gqlsrv/api/types"
	api "github.com/obiwan007/usersrv/proto"
	// "github.com/pkg/errors"
)

func (r *Resolver) User(ctx context.Context, args struct{ Id *string }) (*types.UserResolver, error) {

	request := &api.Id{Id: *args.Id}

	res, err := r.userSvc.GetUser(ctx, request)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)
	s := types.UserResolver{
		R: &types.User{Name: res.Name, Email: res.Email, ID: res.Id},
	}

	return &s, nil

}
func (r *Resolver) AllUsers(ctx context.Context) (*[]*types.UserResolver, error) {
	// t := ctx.Value("jwt")

	// token, ok := t.(*jwt.Token)
	// if !ok || !token.Valid {
	// 	return nil, errors.Errorf("Unauthorized")
	// }

	// fmt.Println(t)
	// claims := token.Claims.(*MyCustomClaims)
	// log.Println("Subject:", claims.Subject)
	users, err := r.userSvc.GetUsers(ctx, &api.ListUsers{})
	if err != nil {
		return nil, err
	}
	var userRxs []*types.UserResolver
	for _, res := range users.Users {
		user := &types.User{Name: res.Name, Email: res.Email, ID: res.Id}
		s := types.UserResolver{R: user}
		userRxs = append(userRxs, &s)

	}

	return &userRxs, nil

}

// type userInput struct {
// 	Name     *string
// 	Password *string
// 	Email    *string
// }

func (r *Resolver) CreateUser(ctx context.Context, args struct{ User types.UserInput }) (*string, error) {

	request := &api.User{Name: *args.User.Name, Password: *args.User.Password, Email: *args.User.Email}

	res, err := r.userSvc.RegisterUser(ctx, request)
	if err != nil {
		return nil, err
	}
	name := res.GetName()
	return &name, nil
}
