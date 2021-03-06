package gql

import (
	"context"

	"github.com/leodotcloud/log"

	pb "github.com/obiwan007/usersrv/proto"
)

func (r *Resolver) AllProjects(ctx context.Context) (*[]*ProjectResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	query := &pb.ListProject{Jwt: token.Raw}
	result, err := r.projectSvc.GetAll(ctx, query)
	if err != nil {
		return nil, err
	}
	var userRxs []*ProjectResolver
	for _, res := range result.Projects {
		t := project2Gql(res)
		userRxs = append(userRxs, &ProjectResolver{R: t, Root: r})
	}

	return &userRxs, nil
}

// func (r *Resolver) RunningTimer(ctx context.Context) (*TimerResolver, error) {

// 	test := &Timer{Description: "Hallo Leute", ID: "1"}
// 	s := TimerResolver{R: test}

// 	return &s, nil
// }

func (r *Resolver) CreateProject(ctx context.Context, arg *CreateProjectRequest) (*ProjectResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("Create %v", arg.P.Description)
	t := projectGql2pb(&arg.P)
	t.Jwt = token.Raw

	result, err := r.projectSvc.Add(ctx, t)

	if err != nil {
		return nil, err
	}

	s := ProjectResolver{R: project2Gql(result), Root: r}

	return &s, nil
}

func (r *Resolver) UpdateProject(ctx context.Context, arg *UpdateProjectRequest) (*ProjectResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("Update %v", arg.P.ID)

	t := projectGql2pb(&arg.P)
	t.Jwt = token.Raw
	result, err := r.projectSvc.Update(ctx, t)

	if err != nil {
		return nil, err
	}

	s := ProjectResolver{R: project2Gql(result), Root: r}

	return &s, nil
}

func (r *Resolver) GetProject(ctx context.Context, arg *GetProjectRequest) (*ProjectResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("ID %v", *arg.ID)

	t := &pb.Id{Id: *arg.ID, Jwt: token.Raw}
	result, err := r.projectSvc.Get(ctx, t)

	if err != nil {
		return nil, err
	}

	s := ProjectResolver{R: project2Gql(result), Root: r}

	return &s, nil
}

func (r *Resolver) DeleteProject(ctx context.Context, arg *DeleteProjectRequest) (*ProjectResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	log.Infof("ID %v", *&arg.ProjectId)

	t := &pb.Id{Id: arg.ProjectId, Jwt: token.Raw}
	result, err := r.projectSvc.Del(ctx, t)

	if err != nil {
		return nil, err
	}

	s := ProjectResolver{R: project2Gql(result), Root: r}

	return &s, nil
}

func project2Gql(result *pb.Project) *Project {
	test := &Project{
		Description: result.Description,
		ID:          result.Id,
		Team:        result.Team,
		Name:        result.Name,
		Status:      result.Status,
		Client:      &Client{ID: result.Client},
	}
	return test
}

func projectGql2pb(arg *ProjectInput) *pb.Project {
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

func checkNilBool(ptr *bool, def bool) bool {
	if ptr == nil {
		return def
	}
	return *ptr
}
