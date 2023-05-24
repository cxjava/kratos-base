package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

func (Address) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "address"},
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		schema.Comment("Host Address"),
	}
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").Comment("主键"),
		field.Int32("street").
			Comment("街道"),
		field.Time("joinTime").
			Comment("加入时间"),
		field.String("host").
			Comment("户主"),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return nil
}
