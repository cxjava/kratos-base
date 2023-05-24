package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}
type MetaData struct {
	Name   string
	Adress string
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user"},
		// Adding this annotation to the schema enables
		// comments for the table and all its fields.
		entsql.WithComments(true),
		schema.Comment("Comment that appears in both the schema and the generated code"),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	fields := []ent.Field{
		field.Int64("id"),
		field.Int32("age").
			Comment("年龄"),
		field.String("name").
			Comment("姓名"),
	}
	// fields = append(fields, IdFields()...)
	// fields = append(fields, TimeFields()...)
	// fields = append(fields, VersionField()...)
	return fields
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
