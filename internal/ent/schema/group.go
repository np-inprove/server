package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
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
		field.String("name").
			NotEmpty().
			Comment("Name of the group (example: CSF01 2023)"),
		field.String("short_name").
			NotEmpty().
			Comment("Short name of the group (example: csf01-2023)"),
		field.String("description").
			Comment("Description of the group"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type).
			Through("group_users", GroupUser.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Comment("Users that are in the group"),
		edge.To("events", Event.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Comment("Events from the group"),
		edge.To("forums", Forum.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Comment("Forum from the group"),
		edge.To("deadlines", Deadline.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Comment("Deadlines created by users from the group"),
		edge.To("invites", GroupInviteLink.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)).
			Comment("Invites for the group"),
		edge.From("institution", Institution.Type).
			Ref("groups").
			Required().
			Unique().
			Comment("Institution owning this group"),
	}
}
