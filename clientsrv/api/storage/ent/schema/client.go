package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Client holds the schema definition for the Client entity.
type Client struct {
	ent.Schema
}

// Fields of the Client.
func (Client) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Default(""),
		field.String("name").Default(""),
		field.String("address").Optional().Default(""),
		field.String("userid").Optional().Default(""),
		field.String("mandantid").Optional().Default(""),
	}
}

// Edges of the Client.
func (Client) Edges() []ent.Edge {
	return nil
}
