package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").
			NotEmpty().
			Comment("First name of the user"),
		field.String("last_name").
			NotEmpty().
			Comment("Last name of the user"),
		field.String("email").
			NotEmpty().
			Comment("Email of the user"),
		field.String("password_hash").
			NotEmpty().
			Sensitive().
			Comment("Password hash of the user"),
		field.Int("points").
			Positive().
			Comment("Points of the user.\nMust always be positive"),
		field.Int("points_awarded_count").
			Positive().
			Comment("Points awarded by the user after points_awarded_reset_time.\nMust always be positive"),
		field.Time("points_awarded_reset_time").
			Optional().
			Comment("Time when points_awarded_count was last reset to 0"),
		field.Bool("superuser").
			Comment("Superuser of the iNProve platform"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("institution", Institution.Type).
			Ref("admins"),
		edge.From("course", Course.Type).
			Ref("students").
			Unique(),
	}
}
