package gql

import (
	"context"

	"github.com/obiwan007/usersrv/gqlsrv/api/types"
	api "github.com/obiwan007/usersrv/proto"
	// "github.com/pkg/errors"
)

type Resolver struct {
	types.GqlResolver
	userSvc  api.UserServiceClient
	timerSvc api.TimerServiceClient
}

func NewResolver(userSrvClient api.UserServiceClient, timerSrvClient api.TimerServiceClient) *Resolver {
	return &Resolver{
		userSvc:  userSrvClient,
		timerSvc: timerSrvClient,
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
