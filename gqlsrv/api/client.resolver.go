package gql

import (
	"context"
	"log"

	pb "github.com/obiwan007/usersrv/proto"
)

func (r *Resolver) AllClients(ctx context.Context) (*[]*ClientResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	query := &pb.ListClient{Jwt: token.Raw}
	result, err := r.clientSvc.GetAll(ctx, query)
	if err != nil {
		return nil, err
	}
	var userRxs []*ClientResolver
	for _, res := range result.Clients {
		t := client2Gql(res)
		userRxs = append(userRxs, &ClientResolver{R: t})
	}

	return &userRxs, nil
}

// func (r *Resolver) RunningTimer(ctx context.Context) (*TimerResolver, error) {

// 	test := &Timer{Description: "Hallo Leute", ID: "1"}
// 	s := TimerResolver{R: test}

// 	return &s, nil
// }

func (r *Resolver) CreateClient(ctx context.Context, arg *CreateClientRequest) (*ClientResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("Create", arg.C.Description)
	t := clientGql2pb(&arg.C)
	t.Jwt = token.Raw
	result, err := r.clientSvc.Add(ctx, t)

	if err != nil {
		return nil, err
	}

	s := ClientResolver{R: client2Gql(result)}

	return &s, nil
}

func (r *Resolver) UpdateClient(ctx context.Context, arg *UpdateClientRequest) (*ClientResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("Update", arg.C.ID)

	t := clientGql2pb(&arg.C)
	t.Jwt = token.Raw
	result, err := r.clientSvc.Update(ctx, t)

	if err != nil {
		return nil, err
	}

	s := ClientResolver{R: client2Gql(result)}

	return &s, nil
}

func (r *Resolver) GetClient(ctx context.Context, arg *GetClientRequest) (*ClientResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("ID", *arg.ID)

	t := &pb.Id{Id: *arg.ID, Jwt: token.Raw}
	result, err := r.clientSvc.Get(ctx, t)

	if err != nil {
		return nil, err
	}

	s := ClientResolver{R: client2Gql(result)}

	return &s, nil
}

func (r *Resolver) DeleteClient(ctx context.Context, arg *DeleteClientRequest) (*ClientResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	log.Println("ID", *&arg.ClientId)

	t := &pb.Id{Id: arg.ClientId, Jwt: token.Raw}
	result, err := r.clientSvc.Del(ctx, t)

	if err != nil {
		return nil, err
	}

	s := ClientResolver{R: client2Gql(result)}

	return &s, nil
}

func client2Gql(result *pb.Client) *Client {
	test := &Client{
		Description: result.Description,
		ID:          result.Id,
		Name:        result.Name,
		// Address:     result.Address,
	}
	return test
}

func clientGql2pb(arg *ClientInput) *pb.Client {
	t := &pb.Client{
		Id:          checkNil(arg.ID, ""),
		Description: checkNil(arg.Description, ""),
		Name:        checkNil(arg.Name, ""),
	}
	return t
}
