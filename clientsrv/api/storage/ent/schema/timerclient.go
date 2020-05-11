package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// TimerClient holds the schema definition for the TimerClient entity.
type TimerClient struct {
	ent.Schema
}

// Fields of the TimerClient.
func (TimerClient) Fields() []ent.Field {
	// Fields of the Client.
	return []ent.Field{
		field.String("description").Default(""),
		field.String("name").Default(""),
		field.String("address").Optional().Default(""),
		field.String("userid").Optional().Default(""),
		field.String("mandantid").Optional().Default(""),
	}
}

// Edges of the TimerClient.
func (TimerClient) Edges() []ent.Edge {
	return nil
}
