package gql

import (
	"context"
	// "github.com/pkg/errors"
)

type Timer struct {
	Id          string
	Description string
	Teammember  string
	// client: Client
	// project: Project
	Tags           string
	ElapsedSeconds int32
	TimerStart     string
	TimerEnd       string
	IsRunning      bool
	IsBilled       bool
}

type TimerResolver struct{ u *Timer }

func (r *TimerResolver) Description() *string {
	return &r.u.Description
}

func (r *TimerResolver) Id() *string {
	return &r.u.Id
}

func (r *TimerResolver) Teammember() *string {
	return &r.u.Teammember
}

func (r *TimerResolver) Tags() *string {
	return &r.u.Tags
}

func (r *TimerResolver) ElapsedSeconds() *int32 {
	return &r.u.ElapsedSeconds
}

func (r *TimerResolver) TimerStart() *string {
	return &r.u.TimerStart
}

func (r *TimerResolver) TimerEnd() *string {
	return &r.u.TimerEnd
}

func (r *TimerResolver) IsRunning() *bool {
	return &r.u.IsRunning
}
func (r *TimerResolver) IsBilled() *bool {
	return &r.u.IsBilled
}

func (r *Resolver) AllTimer(ctx context.Context) (*[]*TimerResolver, error) {

	var userRxs []*TimerResolver
	test := &Timer{Description: "Hallo Leute", Id: "1"}
	s := TimerResolver{u: test}
	userRxs = append(userRxs, &s)

	return &userRxs, nil
}
