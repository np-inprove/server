package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Voucher holds the schema definition for the Voucher entity.
type Voucher struct {
	ent.Schema
}

// Mixins of the Voucher
func (Voucher) Mixin() []ent.Mixin {
	return []ent.Mixin{
		PrizeMixin{},
	}
}

// Fields of the Voucher.
func (Voucher) Fields() []ent.Field {
	return nil
}

// Edges of the Voucher.
func (Voucher) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("institution", Institution.Type).
			Ref("vouchers").
			Required().
			Unique().
			Comment("Institution that owns this voucher"),
	}
}
