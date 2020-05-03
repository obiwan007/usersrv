package gql

import (
	"context"
	"log"

	"github.com/obiwan007/usersrv/gqlsrv/api/types"
)

func (r *Resolver) AllTimer(ctx context.Context, args *types.AllTimerRequest) (*[]*types.TimerResolver, error) {
	log.Println(*args.Filter.Dayrange)
	var userRxs []*types.TimerResolver
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

func (r *Resolver) Timer(ctx context.Context, arg *types.TimerRequest) (*types.TimerResolver, error) {
	log.Println("ID", *arg.ID)
	test := &types.Timer{Description: "Hallo Leute", ID: "1"}
	s := types.TimerResolver{R: test}

	return &s, nil
}

func (r *Resolver) StartTimer(ctx context.Context, arg *types.StartTimerRequest) (*types.TimerResolver, error) {
	log.Println("startTimer ID", arg.TimerId)
	test := &types.Timer{Description: "Hallo Leute", ID: "1"}
	s := types.TimerResolver{R: test}

	return &s, nil
}

func (r *Resolver) StopTimer(ctx context.Context, arg *types.StopTimerRequest) (*types.TimerResolver, error) {
	log.Println("stopTimer ID", arg.TimerId)
	test := &types.Timer{Description: "Hallo Leute", ID: "1"}
	s := types.TimerResolver{R: test}

	return &s, nil
}
