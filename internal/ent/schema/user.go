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
		field.Bool("god_mode").
			Comment("Superuser of the iNProve platform"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("course", Course.Type).
			Ref("students").
			Unique(),
		edge.From("institution", Institution.Type).
			Ref("admins"),
		edge.From("redemptions", Redemption.Type).
			Ref("user"),
		edge.From("forum_posts", ForumPost.Type).
			Ref("author"),
		edge.From("pet", Pet.Type).
			Ref("owner").
			Through("user_pets", UserPet.Type),
		edge.From("groups", Group.Type).
			Ref("users").
			Through("group_users", GroupUser.Type),
		edge.From("reacted_posts", ForumPost.Type).
			Ref("reacted_users").
			Through("reactions", Reaction.Type),
	}
}
