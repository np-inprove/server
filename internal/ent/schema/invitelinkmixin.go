package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// InviteLinkMixin holds the schema definition for the InviteLinkMixin entity.
type InviteLinkMixin struct {
	mixin.Schema
}

// Fields of the InviteLinkMixin.
func (InviteLinkMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").
			NotEmpty().
			Unique().
			Comment("Code used in the invite link"),
	}
}
