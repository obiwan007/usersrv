package gql

import (
	"context"

	api "github.com/obiwan007/usersrv/proto"
)

var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	type Query {
		hello(name: String!): String!
	}
	
	type Mutation {
  		registerUser(user: UserInput!):String
	}

	input UserInput {
		name: String
		email: String
		password: String
	  }
`

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
