package types

import (
	"context"
	"log"
	// "github.com/pkg/errors"
)

type TimerResolver struct {
	R *Timer
	// Root *IResolver
}

func (r TimerResolver) Project(ctx context.Context) *ProjectResolver {
	log.Println("Resolvong Project")
	return &ProjectResolver{r.R.Project}
}
