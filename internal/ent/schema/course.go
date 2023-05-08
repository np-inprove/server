package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Name of the course"),
		field.String("course_id").
			NotEmpty().
			Comment("Short code of the course (example: CSF)"),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("students", User.Type).
			Comment("Students taking this course"),
		edge.From("academic_school", AcademicSchool.Type).
			Ref("courses").
			Unique(),
	}
}
