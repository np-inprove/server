package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Name of the event"),
		field.Time("start_time").
			Comment("Start time of the event.\nIf event is all-day, set to 00:00:00"),
		field.Time("end_time").
			Optional().
			Nillable().
			Comment("End time of the event.\nIf event is all-day, set to nil"),
		field.String("location").
			Optional().
			Comment("Location of the event"),
		field.String("repeat_pattern").
			Optional().
			Comment("Repeat pattern of the event using cron syntax"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).
			Unique().
			Ref("events").
			Required(),
	}
}
