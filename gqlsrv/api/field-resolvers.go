package gql

import (
	"context"

	pb "github.com/obiwan007/usersrv/proto"
	// "github.com/pkg/errors"
)

type TimerResolver struct {
	R    *Timer
	Root *Resolver
}

type ProjectResolver struct {
	R    *Project
	Root *Resolver
}

func (r TimerResolver) Project(ctx context.Context) *ProjectResolver {
	if r.R.Project == nil {
		return nil
	}

	t := &pb.Id{Id: r.R.Project.ID}
	result, err := r.Root.projectSvc.Get(ctx, t)

	if err != nil {
		return nil
	}

	s := ProjectResolver{R: project2Gql(result), Root: r.Root}

	return &s
	// return &ProjectResolver{r.R.Project}
}

func (r TimerResolver) Client(ctx context.Context) *ClientResolver {

	t := &pb.Id{Id: r.R.Client.ID}
	result, err := r.Root.clientSvc.Get(ctx, t)

	if err != nil {
		return nil
	}

	s := ClientResolver{R: client2Gql(result)}

	return &s
	// return &ProjectResolver{r.R.Project}
}

func (r ProjectResolver) Client(ctx context.Context) *ClientResolver {

	t := &pb.Id{Id: r.R.Client.ID}
	result, err := r.Root.clientSvc.Get(ctx, t)

	if err != nil {
		return nil
	}

	s := ClientResolver{R: client2Gql(result)}

	return &s
}
