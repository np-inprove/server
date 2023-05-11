package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GroupUser holds the schema definition for the GroupUser entity.
type GroupUser struct {
	ent.Schema
}

// Annotations of the GroupUser.
func (GroupUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("group_id", "user_id"),
	}
}

// Fields of the GroupUser.
func (GroupUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("group_id").
			Comment("ID of the group"),
		field.Int("user_id").
			Comment("ID of the user"),
		field.Enum("role").
			Values("student", "lecturer"),
	}
}

// Edges of the GroupUser.
func (GroupUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("group", Group.Type).
			Unique().
			Required().
			Field("group_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}
