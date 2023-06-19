package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StudyPlan holds the schema definition for the StudyPlan entity.
type StudyPlan struct {
	ent.Schema
}

// Fields of the StudyPlan.
func (StudyPlan) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Name of the study plan"),
		field.String("code").
			Optional().
			Comment("Code for sharing the study plan"),
	}
}

// Edges of the StudyPlan.
func (StudyPlan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("author", User.Type).
			Unique().
			Required().
			Comment("Author of the study plan"),
		edge.To("milestones", Milestone.Type).
			Comment("Milestones of the study plan"),
	}
}
