package gql

import (
	"context"
	"log"

	jwt "github.com/dgrijalva/jwt-go"
	pb "github.com/obiwan007/usersrv/proto"
)

func (r *Resolver) AllTimer(ctx context.Context, args *AllTimerRequest) (*[]*TimerResolver, error) {
	log.Println(*args.Filter.Dayrange)

	t := ctx.Value("jwt")
	token, ok := t.(*jwt.Token)
	log.Println("AllTimer Token", token, ok)

	query := &pb.ListTimer{}
	result, err := r.timerSvc.GetAll(ctx, query)

	if err != nil {
		return nil, err
	}
	var userRxs []*TimerResolver
	for _, res := range result.Timers {
		t := timerToGql(res)
		userRxs = append(userRxs, &TimerResolver{R: t, Root: r})
	}

	return &userRxs, nil
}

func (r *Resolver) RunningTimer(ctx context.Context) (*TimerResolver, error) {

	test := &Timer{Description: "Hallo Leute", ID: "1"}
	s := TimerResolver{R: test, Root: r}

	return &s, nil
}

func (r *Resolver) GetTimer(ctx context.Context, arg *GetTimerRequest) (*TimerResolver, error) {
	log.Println("ID", *arg.ID)

	t := &pb.Id{Id: *arg.ID}
	result, err := r.timerSvc.Get(ctx, t)

	if err != nil {
		return nil, err
	}

	s := TimerResolver{R: timerToGql(result), Root: r}

	return &s, nil
}

func (r *Resolver) StartTimer(ctx context.Context, arg *StartTimerRequest) (*TimerResolver, error) {
	log.Println("startTimer ID", arg.TimerId)

	t := &pb.Id{Id: *&arg.TimerId}
	result, err := r.timerSvc.Start(ctx, t)
	if err != nil {
		return nil, err
	}
	s := TimerResolver{R: timerToGql(result), Root: r}

	return &s, nil
}

func (r *Resolver) StopTimer(ctx context.Context, arg *StopTimerRequest) (*TimerResolver, error) {
	log.Println("stopTimer ID", arg.TimerId)
	t := &pb.Id{Id: *&arg.TimerId}
	result, err := r.timerSvc.Stop(ctx, t)
	if err != nil {
		return nil, err
	}
	s := TimerResolver{R: timerToGql(result), Root: r}
	return &s, nil
}

func (r *Resolver) CreateTimer(ctx context.Context, arg *CreateTimerRequest) (*TimerResolver, error) {

	t := &pb.Timer{
		Description: checkNil(arg.T.Description, ""),
		Project:     checkNil(arg.T.Project, ""),
	}

	result, err := r.timerSvc.Add(ctx, t)

	if err != nil {
		return nil, err
	}

	s := TimerResolver{R: timerToGql(result), Root: r}

	return &s, nil
}
func (r *Resolver) UpdateTimer(ctx context.Context, arg *UpdateTimerRequest) (*TimerResolver, error) {
	log.Println("Update", arg.T.ID)

	t := timerGql2pb(&arg.T)

	result, err := r.timerSvc.Update(ctx, t)

	if err != nil {
		return nil, err
	}

	s := TimerResolver{R: timerToGql(result), Root: r}

	return &s, nil
}

func timerToGql(result *pb.Timer) *Timer {
	test := &Timer{
		Description:    result.Description,
		Client:         &Client{ID: result.Client},
		Project:        &Project{ID: result.Project},
		ElapsedSeconds: result.ElapsedSeconds,
		ID:             result.Id,
		TimerStart:     result.TimerStart,
		TimerEnd:       result.TimerEnd,
		IsBilled:       result.IsBilled,
		IsRunning:      result.IsRunning}

	return test
}

func timerGql2pb(arg *TimerInput) *pb.Timer {
	t := &pb.Timer{
		Id:          checkNil(arg.ID, ""),
		Description: checkNil(arg.Description, ""),
		Name:        checkNil(arg.Name, ""),
		Project:     checkNil(arg.Project, ""),
	}
	return t
}
