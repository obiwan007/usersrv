package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// Project holds the schema definition for the Project entity.
type Project struct {
	ent.Schema
}

// Fields of the Project.
func (Project) Fields() []ent.Field {
	return []ent.Field{
		field.String("description").Default(""),
		field.String("name").Default(""),
		field.String("clientid").Optional().Default(""),
		field.String("userid").Optional().Default(""),
		field.String("mandantid").Optional().Default(""),
		field.String("status").Optional().Default(""),
		field.String("tags").Optional().Default(""),
		field.String("team").Optional().Default(""),
	}
}

// Edges of the Project.
func (Project) Edges() []ent.Edge {
	return nil
}
