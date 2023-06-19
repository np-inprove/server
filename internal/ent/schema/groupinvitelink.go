package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/entity/group"
)

// GroupInviteLink holds the schema definition for the GroupInviteLink entity.
type GroupInviteLink struct {
	ent.Schema
}

func (GroupInviteLink) Mixin() []ent.Mixin {
	return []ent.Mixin{
		InviteLinkMixin{},
	}
}

// Fields of the GroupInviteLink.
func (GroupInviteLink) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").
			GoType(group.Role("")).
			Comment("Role granted in the invite"),
	}
}

// Edges of the GroupInviteLink.
func (GroupInviteLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).
			Ref("invites").
			Unique().
			Required(),
	}
}
