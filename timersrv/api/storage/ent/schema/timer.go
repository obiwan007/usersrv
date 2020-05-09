package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Timer holds the schema definition for the Timer entity.
type Timer struct {
	ent.Schema
}

// Fields of the Timer.
func (Timer) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").
			Default(""),
		field.String("projectid").
			Optional().
			Default(""),
		field.String("tags").
			Optional().
			Default(""),
		field.Time("timerStart").Optional(),

		field.Time("timerEnd").Optional(),
		field.Int("elapsedSeconds").Optional().
			Default(0),
		field.Bool("IsBilled").Optional().
			Default(false),
		field.Bool("IsRunning").Optional().
			Default(false),
		field.String("userid").Optional().
			Default(""),
		field.String("mandantid").Optional().
			Default(""),
	}
}

// Edges of the Timer.
func (Timer) Edges() []ent.Edge {
	return nil
}
