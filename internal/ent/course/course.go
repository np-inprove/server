// Code generated by ent, DO NOT EDIT.

package course

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the course type in the database.
	Label = "course"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCourseID holds the string denoting the course_id field in the database.
	FieldCourseID = "course_id"
	// EdgeStudents holds the string denoting the students edge name in mutations.
	EdgeStudents = "students"
	// EdgeAcademicSchool holds the string denoting the academic_school edge name in mutations.
	EdgeAcademicSchool = "academic_school"
	// Table holds the table name of the course in the database.
	Table = "courses"
	// StudentsTable is the table that holds the students relation/edge.
	StudentsTable = "users"
	// StudentsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	StudentsInverseTable = "users"
	// StudentsColumn is the table column denoting the students relation/edge.
	StudentsColumn = "course_students"
	// AcademicSchoolTable is the table that holds the academic_school relation/edge.
	AcademicSchoolTable = "courses"
	// AcademicSchoolInverseTable is the table name for the AcademicSchool entity.
	// It exists in this package in order to avoid circular dependency with the "academicschool" package.
	AcademicSchoolInverseTable = "academic_schools"
	// AcademicSchoolColumn is the table column denoting the academic_school relation/edge.
	AcademicSchoolColumn = "academic_school_courses"
)

// Columns holds all SQL columns for course fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCourseID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "courses"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"academic_school_courses",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// CourseIDValidator is a validator for the "course_id" field. It is called by the builders before save.
	CourseIDValidator func(string) error
)

// OrderOption defines the ordering options for the Course queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCourseID orders the results by the course_id field.
func ByCourseID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCourseID, opts...).ToFunc()
}

// ByStudentsCount orders the results by students count.
func ByStudentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStudentsStep(), opts...)
	}
}

// ByStudents orders the results by students terms.
func ByStudents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStudentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAcademicSchoolField orders the results by academic_school field.
func ByAcademicSchoolField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAcademicSchoolStep(), sql.OrderByField(field, opts...))
	}
}
func newStudentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StudentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StudentsTable, StudentsColumn),
	)
}
func newAcademicSchoolStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AcademicSchoolInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AcademicSchoolTable, AcademicSchoolColumn),
	)
}