// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldPoints holds the string denoting the points field in the database.
	FieldPoints = "points"
	// FieldPointsAwardedCount holds the string denoting the points_awarded_count field in the database.
	FieldPointsAwardedCount = "points_awarded_count"
	// FieldPointsAwardedResetTime holds the string denoting the points_awarded_reset_time field in the database.
	FieldPointsAwardedResetTime = "points_awarded_reset_time"
	// FieldSuperuser holds the string denoting the superuser field in the database.
	FieldSuperuser = "superuser"
	// EdgeInstitution holds the string denoting the institution edge name in mutations.
	EdgeInstitution = "institution"
	// Table holds the table name of the user in the database.
	Table = "users"
	// InstitutionTable is the table that holds the institution relation/edge. The primary key declared below.
	InstitutionTable = "institution_admins"
	// InstitutionInverseTable is the table name for the Institution entity.
	// It exists in this package in order to avoid circular dependency with the "institution" package.
	InstitutionInverseTable = "institutions"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldFirstName,
	FieldLastName,
	FieldEmail,
	FieldPasswordHash,
	FieldPoints,
	FieldPointsAwardedCount,
	FieldPointsAwardedResetTime,
	FieldSuperuser,
}

var (
	// InstitutionPrimaryKey and InstitutionColumn2 are the table columns denoting the
	// primary key for the institution relation (M2M).
	InstitutionPrimaryKey = []string{"institution_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// PasswordHashValidator is a validator for the "password_hash" field. It is called by the builders before save.
	PasswordHashValidator func(string) error
	// PointsValidator is a validator for the "points" field. It is called by the builders before save.
	PointsValidator func(int) error
	// PointsAwardedCountValidator is a validator for the "points_awarded_count" field. It is called by the builders before save.
	PointsAwardedCountValidator func(int) error
)

// OrderOption defines the ordering options for the User queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByFirstName orders the results by the first_name field.
func ByFirstName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFirstName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// ByPasswordHash orders the results by the password_hash field.
func ByPasswordHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

// ByPoints orders the results by the points field.
func ByPoints(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPoints, opts...).ToFunc()
}

// ByPointsAwardedCount orders the results by the points_awarded_count field.
func ByPointsAwardedCount(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPointsAwardedCount, opts...).ToFunc()
}

// ByPointsAwardedResetTime orders the results by the points_awarded_reset_time field.
func ByPointsAwardedResetTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPointsAwardedResetTime, opts...).ToFunc()
}

// BySuperuser orders the results by the superuser field.
func BySuperuser(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSuperuser, opts...).ToFunc()
}

// ByInstitutionCount orders the results by institution count.
func ByInstitutionCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newInstitutionStep(), opts...)
	}
}

// ByInstitution orders the results by institution terms.
func ByInstitution(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newInstitutionStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newInstitutionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(InstitutionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, InstitutionTable, InstitutionPrimaryKey...),
	)
}
