package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AcademicSchool holds the schema definition for the AcademicSchool entity.
type AcademicSchool struct {
	ent.Schema
}

// Fields of the AcademicSchool.
func (AcademicSchool) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Comment("Name of the academic school"),
		field.String("school_code").
			NotEmpty().
			Comment("Short code of the academic school (example: ICT)"),
	}
}

// Edges of the AcademicSchool.
func (AcademicSchool) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("institution", Institution.Type).
			Ref("academic_schools").
			Unique(),
		edge.To("courses", Course.Type),
	}
}
