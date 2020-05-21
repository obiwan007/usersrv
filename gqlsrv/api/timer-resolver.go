package gql

import (
	"context"
	"log"
	"strconv"

	pb "github.com/obiwan007/usersrv/proto"
)

func (r *Resolver) AllTimer(ctx context.Context, args *AllTimerRequest) (*[]*TimerResolver, error) {
	log.Println(*args.Filter.Dayrange)
	dayRange, err := strconv.Atoi(*args.Filter.Dayrange)
	if err != nil {
		dayRange = -1
	}
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}

	query := &pb.ListTimer{Jwt: token.Raw, DayRange: int32(dayRange)}
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

	_, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}

	test := &Timer{Description: "Hallo Leute", ID: "1"}
	s := TimerResolver{R: test, Root: r}

	return &s, nil
}

func (r *Resolver) GetTimer(ctx context.Context, arg *GetTimerRequest) (*TimerResolver, error) {
	log.Println("ID", *arg.ID)

	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}

	t := &pb.Id{Id: *arg.ID, Jwt: token.Raw}
	result, err := r.timerSvc.Get(ctx, t)

	if err != nil {
		return nil, err
	}

	s := TimerResolver{R: timerToGql(result), Root: r}

	return &s, nil
}

func (r *Resolver) DeleteTimer(ctx context.Context, arg *DeleteTimerRequest) (*TimerResolver, error) {
	log.Println("ID", arg.TimerId)

	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}

	t := &pb.Id{Id: arg.TimerId, Jwt: token.Raw}
	result, err := r.timerSvc.Del(ctx, t)

	if err != nil {
		return nil, err
	}

	s := TimerResolver{R: timerToGql(result), Root: r}

	return &s, nil
}

func (r *Resolver) StartTimer(ctx context.Context, arg *StartTimerRequest) (*TimerResolver, error) {
	log.Println("startTimer ID", arg.TimerId)
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}

	t := &pb.Id{Id: *&arg.TimerId, Jwt: token.Raw}
	result, err := r.timerSvc.Start(ctx, t)
	if err != nil {
		return nil, err
	}
	s := TimerResolver{R: timerToGql(result), Root: r}

	return &s, nil
}

func (r *Resolver) StopTimer(ctx context.Context, arg *StopTimerRequest) (*TimerResolver, error) {
	log.Println("stopTimer ID", arg.TimerId)
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	t := &pb.Id{Id: *&arg.TimerId, Jwt: token.Raw}
	result, err := r.timerSvc.Stop(ctx, t)
	if err != nil {
		return nil, err
	}
	s := TimerResolver{R: timerToGql(result), Root: r}
	return &s, nil
}

func (r *Resolver) CreateTimer(ctx context.Context, arg *CreateTimerRequest) (*TimerResolver, error) {
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	t := &pb.Timer{
		Jwt:         token.Raw,
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
	token, err := validateToken(ctx)
	if err != nil {
		return nil, err
	}
	t := timerGql2pb(&arg.T)
	t.Jwt = token.Raw

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
		Tags:           result.Tags,
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
		Tags:        checkNil(arg.Tags, ""),
		Project:     checkNil(arg.Project, ""),
		TimerStart:  checkNil(arg.TimerStart, ""),
		TimerEnd:    checkNil(arg.TimerEnd, ""),
		IsBilled:    checkNilBool(arg.IsBilled, false),
	}
	return t
}
