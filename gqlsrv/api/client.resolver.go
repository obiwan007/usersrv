package gql

import (
	"context"
	"log"

	"github.com/obiwan007/usersrv/gqlsrv/api/types"
	pb "github.com/obiwan007/usersrv/proto"
)

func (r *Resolver) AllClients(ctx context.Context) (*[]*types.ClientResolver, error) {

	query := &pb.ListClient{}
	result, err := r.clientSvc.GetAll(ctx, query)
	if err != nil {
		return nil, err
	}
	var userRxs []*types.ClientResolver
	for _, res := range result.Clients {
		t := client2Gql(res)
		userRxs = append(userRxs, &types.ClientResolver{R: t})
	}

	return &userRxs, nil
}

// func (r *Resolver) RunningTimer(ctx context.Context) (*types.TimerResolver, error) {

// 	test := &types.Timer{Description: "Hallo Leute", ID: "1"}
// 	s := types.TimerResolver{R: test}

// 	return &s, nil
// }

func (r *Resolver) CreateClient(ctx context.Context, arg *types.CreateClientRequest) (*types.ClientResolver, error) {
	log.Println("Create", arg.C.Description)
	t := clientGql2pb(&arg.C)
	result, err := r.clientSvc.Add(ctx, t)

	if err != nil {
		return nil, err
	}

	s := types.ClientResolver{R: client2Gql(result)}

	return &s, nil
}

func (r *Resolver) UpdateClient(ctx context.Context, arg *types.UpdateClientRequest) (*types.ClientResolver, error) {
	log.Println("Update", arg.C.ID)

	t := clientGql2pb(&arg.C)

	result, err := r.clientSvc.Update(ctx, t)

	if err != nil {
		return nil, err
	}

	s := types.ClientResolver{R: client2Gql(result)}

	return &s, nil
}

func (r *Resolver) GetClient(ctx context.Context, arg *types.GetClientRequest) (*types.ClientResolver, error) {
	log.Println("ID", *arg.ID)

	t := &pb.Id{Id: *arg.ID}
	result, err := r.clientSvc.Get(ctx, t)

	if err != nil {
		return nil, err
	}

	s := types.ClientResolver{R: client2Gql(result)}

	return &s, nil
}

func (r *Resolver) DeleteClient(ctx context.Context, arg *types.DeleteClientRequest) (*types.ClientResolver, error) {
	log.Println("ID", *&arg.ClientId)

	t := &pb.Id{Id: arg.ClientId}
	result, err := r.clientSvc.Del(ctx, t)

	if err != nil {
		return nil, err
	}

	s := types.ClientResolver{R: client2Gql(result)}

	return &s, nil
}

func client2Gql(result *pb.Client) *types.Client {
	test := &types.Client{
		Description: result.Description,
		ID:          result.Id,
		Name:        result.Name,
		// Address:     result.Address,
	}
	return test
}

func clientGql2pb(arg *types.ClientInput) *pb.Client {
	t := &pb.Client{
		Id:          checkNil(arg.ID, ""),
		Description: checkNil(arg.Description, ""),
		Name:        checkNil(arg.Name, ""),
	}
	return t
}
