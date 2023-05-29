package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").
			NotEmpty().
			Unique().
			Comment("URL path of the group (example: csf01-2023)"),
		field.String("name").
			NotEmpty().
			Comment("Name of the group (example: CSF01 2023)"),
		field.String("description").
			Comment("Description of the group"),
		field.Enum("group_type").
			Values("special_interest_group", "module_group", "cca", "mentor_class"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Through("group_users", GroupUser.Type).
			Comment("Users that are in the group"),
		edge.To("events", Event.Type).
			Comment("Events from the group"),
		edge.To("forum_posts", ForumPost.Type).
			Comment("Forum posts from the group"),
		edge.To("deadlines", Deadline.Type).
			Comment("Deadlines created by users from the group"),
		edge.From("institution", Institution.Type).
			Ref("groups").
			Unique().
			Required().
			Comment("Institution owning this group"),
	}
}
