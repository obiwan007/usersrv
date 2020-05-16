package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("userid").Default(""),
		field.String("name").Optional().Default(""),
		field.String("mandantid").Optional().Default(""),
		field.String("roles").Optional().Default(""),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
