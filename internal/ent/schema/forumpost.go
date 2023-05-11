package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ForumPost holds the schema definition for the ForumPost entity.
type ForumPost struct {
	ent.Schema
}

// Fields of the ForumPost.
func (ForumPost) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty().
			Comment("Title of the forum post"),
		field.String("content").
			NotEmpty().
			Comment("Content of the forum post"),
		field.JSON("mentioned_users_json", []int{}).
			Comment("Slice of user IDs that were mentioned"),
	}
}

// Edges of the ForumPost.
func (ForumPost) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).
			Unique().
			Required().
			Comment("Author of the forum post"),
		edge.From("group", Group.Type).
			Ref("forum_posts").
			Unique().
			Required(),
		edge.To("children", ForumPost.Type).
			From("parent").
			Unique(),
		edge.To("reacted_users", User.Type).
			Through("reactions", Reaction.Type),
	}
}
