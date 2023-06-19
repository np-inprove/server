package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/entity/institution"
)

// InstitutionInviteLink holds the schema definition for the InstitutionInviteLink entity.
type InstitutionInviteLink struct {
	ent.Schema
}

func (InstitutionInviteLink) Mixin() []ent.Mixin{
	return []ent.Mixin{
		InviteLinkMixin{},
	}
}

// Fields of the InstitutionInviteLink.
func (InstitutionInviteLink) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("role").
			GoType(institution.Role("")).
			Comment("Role granted in the invite"),
	}
}

// Edges of the InstitutionInviteLink.
func (InstitutionInviteLink) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("institution", Institution.Type).
			Ref("invites").
			Unique().
			Required(),
	}
}
