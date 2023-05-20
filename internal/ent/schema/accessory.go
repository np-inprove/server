package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Accessory holds the schema definition for the Accessory entity.
type Accessory struct {
	ent.Schema
}

// Mixin included in the Accessory
func (Accessory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		PrizeMixin{},
	}
}

// Fields of the Accessory.
func (Accessory) Fields() []ent.Field {
	return nil
}

// Edges of the Accessory.
func (Accessory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("institution", Institution.Type).
			Ref("accessories").
			Required().
			Unique().
			Comment("Institution that owns this accessory"),
	}
}
