package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Milestone holds the schema definition for the Milestone entity.
type Milestone struct {
	ent.Schema
}

// Fields of the Milestone.
func (Milestone) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Name of the milestone"),
		field.Time("target_completion_time").
			Comment("Time when the milestone is should be completed"),
	}
}

// Edges of the Milestone.
func (Milestone) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("study_plan", StudyPlan.Type).
			Ref("milestones").
			Required().
			Unique(),
	}
}
