package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Institution holds the schema definition for the Institution entity.
type Institution struct {
	ent.Schema
}

// Fields of the Institution.
func (Institution) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Name of the institution"),
		field.String("short_name").
			NotEmpty().
			Unique().
			Comment("Short name of the institution (example: np)"),
	}
}

// Edges of the Institution.
func (Institution) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admins", User.Type).
			Comment("Admins of the institution"),
		edge.To("vouchers", Voucher.Type).
			Comment("Prizes (vouchers) available to be redeemed by users of the institution"),
		edge.To("accessories", Accessory.Type).
			Comment("Prizes (accessories) available to be redeemed by users of the institution"),
		edge.To("departments", Department.Type).
			Comment("Departments of the institution"),
	}
}
