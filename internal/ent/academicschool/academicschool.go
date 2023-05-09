// Code generated by ent, DO NOT EDIT.

package academicschool

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the academicschool type in the database.
	Label = "academic_school"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSchoolCode holds the string denoting the school_code field in the database.
	FieldSchoolCode = "school_code"
	// EdgeInstitution holds the string denoting the institution edge name in mutations.
	EdgeInstitution = "institution"
	// EdgeCourses holds the string denoting the courses edge name in mutations.
	EdgeCourses = "courses"
	// Table holds the table name of the academicschool in the database.
	Table = "academic_schools"
	// InstitutionTable is the table that holds the institution relation/edge.
	InstitutionTable = "academic_schools"
	// InstitutionInverseTable is the table name for the Institution entity.
	// It exists in this package in order to avoid circular dependency with the "institution" package.
	InstitutionInverseTable = "institutions"
	// InstitutionColumn is the table column denoting the institution relation/edge.
	InstitutionColumn = "institution_academic_schools"
	// CoursesTable is the table that holds the courses relation/edge.
	CoursesTable = "courses"
	// CoursesInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	CoursesInverseTable = "courses"
	// CoursesColumn is the table column denoting the courses relation/edge.
	CoursesColumn = "academic_school_courses"
)

// Columns holds all SQL columns for academicschool fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSchoolCode,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "academic_schools"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"institution_academic_schools",
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
	// SchoolCodeValidator is a validator for the "school_code" field. It is called by the builders before save.
	SchoolCodeValidator func(string) error
)

// OrderOption defines the ordering options for the AcademicSchool queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// BySchoolCode orders the results by the school_code field.
func BySchoolCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSchoolCode, opts...).ToFunc()
}

// ByInstitutionField orders the results by institution field.
func ByInstitutionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInstitutionStep(), sql.OrderByField(field, opts...))
	}
}

// ByCoursesCount orders the results by courses count.
func ByCoursesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCoursesStep(), opts...)
	}
}

// ByCourses orders the results by courses terms.
func ByCourses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCoursesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newInstitutionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InstitutionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, InstitutionTable, InstitutionColumn),
	)
}
func newCoursesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CoursesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CoursesTable, CoursesColumn),
	)
}