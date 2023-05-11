package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// PrizeMixin holds the schema definition for the PrizeMixin entity.
type PrizeMixin struct {
	mixin.Schema
}

// Fields of the PrizeMixin.
func (PrizeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().
			Comment("Name of the prize"),
		field.String("description").
			Comment("Description of the prize"),
		field.Int("points_required").Positive().
			Comment("Points required to redeem the prize"),
	}
}

// Edges of the PrizeMixin.
func (PrizeMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("redemptions", Redemption.Type).
			Comment("Redemptions that involve this prize"),
	}
}
