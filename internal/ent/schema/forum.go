package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Forum holds the schema definition for the Forum entity.
type Forum struct {
	ent.Schema
}

// Fields of the Forum.
func (Forum) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Name of the forum (example: General)"),
		field.String("short_name").
			NotEmpty().
			Comment("Short name of the forum (example: general)"),
		field.String("description").
			Comment("Description of the forum"),
	}
}

// Edges of the Forum.
func (Forum) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).
			Ref("forums").
			Required().
			Unique(),
		edge.To("posts", ForumPost.Type).
			Comment("Posts in the forum"),
	}
}
