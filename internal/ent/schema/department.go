package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Name of the department"),
		field.String("short_name").
			NotEmpty().
			Comment("Short name of the department (example: ict).\nMust be unique within an institution"),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("institution", Institution.Type).
			Ref("departments").
			Required().
			Unique(),
		edge.To("children", Department.Type).
			From("parent").
			Unique(),
		edge.To("users", User.Type).
			Comment("Users in this department"),
	}
}
