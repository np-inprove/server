package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Prize holds the schema definition for the Prize entity.
type Prize struct {
	ent.Schema
}

// Fields of the Prize.
func (Prize) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().
			Comment("Name of the prize"),
		field.String("description").
			Comment("Description of the prize"),
		field.Int("points_required").Positive().
			Comment("Points required to redeem the prize"),
		field.Enum("discriminator").
			Values("voucher", "accessory").
			Comment("Type of prize, can only be voucher or a pet accessory"),
	}
}

// Edges of the Prize.
func (Prize) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("institution", Institution.Type).
			Ref("prizes").
			Unique(),
		edge.To("redemption_users", User.Type).
			Through("prize_redemptions", PrizeRedemptions.Type).
			Comment("Users who have redeemed this prize"),
	}
}
