package gql

import (
	"context"
	"log"

	"github.com/obiwan007/usersrv/gqlsrv/api/types"
	pb "github.com/obiwan007/usersrv/proto"
)

func (r *Resolver) AllProjects(ctx context.Context) (*[]*types.ProjectResolver, error) {

	query := &pb.ListProject{}
	result, err := r.projectSvc.GetAll(ctx, query)
	if err != nil {
		return nil, err
	}
	var userRxs []*types.ProjectResolver
	for _, res := range result.Projects {
		t := project2Gql(res)
		userRxs = append(userRxs, &types.ProjectResolver{R: t})
	}

	return &userRxs, nil
}

// func (r *Resolver) RunningTimer(ctx context.Context) (*types.TimerResolver, error) {

// 	test := &types.Timer{Description: "Hallo Leute", ID: "1"}
// 	s := types.TimerResolver{R: test}

// 	return &s, nil
// }

func (r *Resolver) CreateProject(ctx context.Context, arg *types.CreateProjectRequest) (*types.ProjectResolver, error) {
	log.Println("Create", arg.P.Description)
	t := projectGql2pb(&arg.P)
	result, err := r.projectSvc.Add(ctx, t)

	if err != nil {
		return nil, err
	}

	s := types.ProjectResolver{R: project2Gql(result)}

	return &s, nil
}

func (r *Resolver) UpdateProject(ctx context.Context, arg *types.UpdateProjectRequest) (*types.ProjectResolver, error) {
	log.Println("Update", arg.P.ID)

	t := projectGql2pb(&arg.P)

	result, err := r.projectSvc.Update(ctx, t)

	if err != nil {
		return nil, err
	}

	s := types.ProjectResolver{R: project2Gql(result)}

	return &s, nil
}

func (r *Resolver) GetProject(ctx context.Context, arg *types.GetProjectRequest) (*types.ProjectResolver, error) {
	log.Println("ID", *arg.ID)

	t := &pb.Id{Id: *arg.ID}
	result, err := r.projectSvc.Get(ctx, t)

	if err != nil {
		return nil, err
	}

	s := types.ProjectResolver{R: project2Gql(result)}

	return &s, nil
}

func (r *Resolver) DeleteProject(ctx context.Context, arg *types.DeleteProjectRequest) (*types.ProjectResolver, error) {
	log.Println("ID", *&arg.ProjectId)

	t := &pb.Id{Id: arg.ProjectId}
	result, err := r.projectSvc.Del(ctx, t)

	if err != nil {
		return nil, err
	}

	s := types.ProjectResolver{R: project2Gql(result)}

	return &s, nil
}

func project2Gql(result *pb.Project) *types.Project {
	test := &types.Project{
		Description: result.Description,
		ID:          result.Id,
		Team:        result.Team,
		Name:        result.Name,
		Status:      result.Status,
	}
	return test
}

func projectGql2pb(arg *types.ProjectInput) *pb.Project {
	t := &pb.Project{
		Id:          checkNil(arg.ID, ""),
		Description: checkNil(arg.Description, ""),
		Name:        checkNil(arg.Name, ""),
		Team:        checkNil(arg.Team, ""),
		Status:      checkNil(arg.Status, ""),
		Client:      checkNil(arg.Client, ""),
	}
	return t
}

func checkNil(ptr *string, def string) string {
	if ptr == nil {
		return def
	}
	return *ptr
}
