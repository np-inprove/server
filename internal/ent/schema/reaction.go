package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Reaction holds the schema definition for the Reaction entity.
type Reaction struct {
	ent.Schema
}

// Annotations of the Reaction.
func (Reaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("forum_post_id", "user_id"),
	}
}

// Fields of the Reaction.
func (Reaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").
			Comment("ID of the user that reacted on the post"),
		field.Int("forum_post_id").
			Comment("ID of the post that was reacted to"),
		field.String("emoji").
			NotEmpty().
			Comment("Emoji used for the reaction"),
	}
}

// Edges of the Reaction.
func (Reaction) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("forum_post", ForumPost.Type).
			Required().
			Unique().
			Field("forum_post_id"),
	}
}
