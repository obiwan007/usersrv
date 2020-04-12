package gql

import (
	"context"

	api "github.com/obiwan007/usersrv/proto"
)

var Schema = `
	schema {
		query: Query
		# mutation: Mutation
	}
	type Query {
		hello(name: String!): String!
	}
	# type Mutation {}
`

type Resolver struct {
	helloSvc api.UserServiceClient
}

func NewResolver(helloClient api.UserServiceClient) *Resolver {
	return &Resolver{
		helloSvc: helloClient,
	}
}

func (r *Resolver) Hello(ctx context.Context, args struct{ Name string }) (string, error) {

	request := &api.User{Name: "Markus", Password: "test"}

	res, err := r.helloSvc.AddUser(ctx, request)
	if err != nil {
		return "", err
	}

	return res.GetName(), nil
}
