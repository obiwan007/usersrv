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
func checkNil(ptr *string, def string) string {
	if ptr == nil {
		return def
	}
	return *ptr
}
func (r *Resolver) CreateProject(ctx context.Context, arg *types.CreateProjectRequest) (*types.ProjectResolver, error) {
	log.Println("Create", arg.P.Description)

	t := &pb.Project{
		Description: checkNil(arg.P.Description, ""),
		Name:        checkNil(arg.P.Name, ""),
		Team:        checkNil(arg.P.Team, ""),
		Status:      checkNil(arg.P.Status, "new"),
		Client:      checkNil(arg.P.Client, ""),
	}
	result, err := r.projectSvc.Add(ctx, t)

	if err != nil {
		return nil, err
	}

	s := types.ProjectResolver{R: project2Gql(result)}

	return &s, nil
}

func (r *Resolver) UpdateProject(ctx context.Context, arg *types.UpdateProjectRequest) (*types.ProjectResolver, error) {
	log.Println("Update", arg.P.Description)

	t := &pb.Project{
		Id:          checkNil(arg.P.ID, ""),
		Description: checkNil(arg.P.Description, ""),
		Name:        checkNil(arg.P.Name, ""),
		Team:        checkNil(arg.P.Team, ""),
		Status:      checkNil(arg.P.Status, "new"),
		Client:      checkNil(arg.P.Client, ""),
	}
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

// func (r *Resolver) StartTimer(ctx context.Context, arg *types.StartTimerRequest) (*types.TimerResolver, error) {
// 	log.Println("startTimer ID", arg.TimerId)

// 	t := &pb.Id{Id: *&arg.TimerId}
// 	result, err := r.timerSvc.Start(ctx, t)
// 	if err != nil {
// 		return nil, err
// 	}
// 	s := types.TimerResolver{R: toGql(result)}

// 	return &s, nil
// }

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
