package gql

import (
	"context"
	"fmt"

	api "github.com/obiwan007/usersrv/proto"
	// "github.com/pkg/errors"
)

func (r *Resolver) User(ctx context.Context, args struct{ Id *string }) (*UserResolver, error) {

	request := &api.Id{Id: *args.Id}

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

	request := &api.User{Name: *args.User.Name, Password: *args.User.Password, Email: *args.User.Email}

	res, err := r.userSvc.RegisterUser(ctx, request)
	if err != nil {
		return nil, err
	}
	name := res.GetName()
	return &name, nil
}
