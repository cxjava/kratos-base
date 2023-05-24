package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Beer struct {
	ent.Schema
}

// Fields of the Beer.
func (Beer) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("name"),
	}
}

// Edges of the Beer.
func (Beer) Edges() []ent.Edge {
	return []ent.Edge{}
}
