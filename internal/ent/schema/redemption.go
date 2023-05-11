package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Redemption holds the schema definition for the Redemption entity.
type Redemption struct {
	ent.Schema
}

// Fields of the Redemptions.
func (Redemption) Fields() []ent.Field {
	return []ent.Field{
		field.Time("redeemed_at").
			Comment("Time when the prize was redeemed"),
	}
}

// Edges of the Redemptions.
func (Redemption) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("voucher", Voucher.Type).
			Ref("redemptions").
			Unique(),
		edge.From("accessory", Accessory.Type).
			Ref("redemptions").
			Unique(),
		edge.To("user", User.Type).
			Required().
			Unique(),
	}
}
