package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PrizeRedemptions holds the schema definition for the PrizeRedemptions entity.
type PrizeRedemptions struct {
	ent.Schema
}

// Annotations of the PrizeRedemptions.
func (PrizeRedemptions) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Create composite key using prize_id and user_id
		field.ID("prize_id", "user_id"),
	}
}

// Fields of the PrizeRedemptions.
func (PrizeRedemptions) Fields() []ent.Field {
	return []ent.Field{
		field.Time("redeemed_at").
			Comment("Time when the prize was redeemed"),
		field.Int("prize_id").
			Comment("ID of the prize"),
		field.Int("user_id").
			Comment("ID of the user"),
	}
}

// Edges of the PrizeRedemptions.
func (PrizeRedemptions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("prize", Prize.Type).
			Required().
			Unique().
			Field("prize_id"),
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
	}
}
