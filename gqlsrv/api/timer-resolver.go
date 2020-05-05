package gql

import (
	"context"
	"log"

	"github.com/obiwan007/usersrv/gqlsrv/api/types"
	pb "github.com/obiwan007/usersrv/proto"
)

func (r *Resolver) AllTimer(ctx context.Context, args *types.AllTimerRequest) (*[]*types.TimerResolver, error) {
	log.Println(*args.Filter.Dayrange)
	query := &pb.ListTimer{}
	result, err := r.timerSvc.GetAll(ctx, query)
	if err != nil {
		return nil, err
	}
	var userRxs []*types.TimerResolver
	for _, res := range result.Timers {
		t := toGql(res)
		userRxs = append(userRxs, &types.TimerResolver{R: t})
	}

	test := &types.Timer{Description: "Hallo Leute", ID: "1"}
	s := types.TimerResolver{R: test}
	userRxs = append(userRxs, &s)

	return &userRxs, nil
}

func (r *Resolver) RunningTimer(ctx context.Context) (*types.TimerResolver, error) {

	test := &types.Timer{Description: "Hallo Leute", ID: "1"}
	s := types.TimerResolver{R: test}

	return &s, nil
}

func (r *Resolver) GetTimer(ctx context.Context, arg *types.GetTimerRequest) (*types.TimerResolver, error) {
	log.Println("ID", *arg.ID)

	t := &pb.Id{Id: *arg.ID}
	result, err := r.timerSvc.Get(ctx, t)

	if err != nil {
		return nil, err
	}

	s := types.TimerResolver{R: toGql(result)}

	return &s, nil
}

func (r *Resolver) StartTimer(ctx context.Context, arg *types.StartTimerRequest) (*types.TimerResolver, error) {
	log.Println("startTimer ID", arg.TimerId)

	t := &pb.Id{Id: *&arg.TimerId}
	result, err := r.timerSvc.Start(ctx, t)
	if err != nil {
		return nil, err
	}
	s := types.TimerResolver{R: toGql(result)}

	return &s, nil
}

func (r *Resolver) StopTimer(ctx context.Context, arg *types.StopTimerRequest) (*types.TimerResolver, error) {
	log.Println("stopTimer ID", arg.TimerId)
	t := &pb.Id{Id: *&arg.TimerId}
	result, err := r.timerSvc.Stop(ctx, t)
	if err != nil {
		return nil, err
	}
	s := types.TimerResolver{R: toGql(result)}
	return &s, nil
}

func toGql(result *pb.Timer) *types.Timer {
	test := &types.Timer{
		Description: result.Description,
		// Client:         result.Client,
		ElapsedSeconds: result.ElapsedSeconds,
		ID:             result.Id,
		TimerStart:     result.TimerStart,
		TimerEnd:       result.TimerEnd,
		IsBilled:       result.IsBilled,
		IsRunning:      result.IsRunning}

	return test
}
