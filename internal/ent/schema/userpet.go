package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// UserPet holds the schema definition for the UserPet entity.
type UserPet struct {
	ent.Schema
}

// Annotations of the UserPet.
func (UserPet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Create composite key using prize_id and user_id
		field.ID("pet_id", "user_id"),
	}
}

// Fields of the UserPet.
func (UserPet) Fields() []ent.Field {
	return []ent.Field{
		field.Float("hunger_percentage").
			Min(0).
			Max(1).
			Comment("Hunger percentage of the pet (0-1)"),
		field.JSON("enabled_svg_group_element_ids", map[string]bool{}).
			Default(map[string]bool{}).
			Comment("Map of enabled SVG group element IDs"),
		field.Int("pet_id").
			Comment("ID of the pet"),
		field.Int("user_id").
			Comment("ID of the user"),
	}
}

// Edges of the UserPet.
func (UserPet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pet", Pet.Type).
			Required().
			Unique().
			Field("pet_id"),
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
	}
}

func (UserPet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").
			Unique(),
	}
}
