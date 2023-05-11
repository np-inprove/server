package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Deadline holds the schema definition for the Deadline entity.
type Deadline struct {
	ent.Schema
}

// Fields of the Deadline.
func (Deadline) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Name of the deadline"),
		field.Time("due_time").
			Optional().
			Nillable().
			Comment("Time when the deadline is due.\nIf there is no due time, set to nil"),
	}
}

// Edges of the Deadline.
func (Deadline) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).
			Unique().
			Required(),
		edge.To("voted_users", User.Type).
			Comment("Users who voted for this deadline"),
		edge.From("group", Group.Type).
			Ref("deadlines").
			Unique().
			Required(),
	}
}
