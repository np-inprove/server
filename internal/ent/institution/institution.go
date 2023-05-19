// Code generated by ent, DO NOT EDIT.

package institution

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the institution type in the database.
	Label = "institution"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldShortName holds the string denoting the short_name field in the database.
	FieldShortName = "short_name"
	// EdgeAdmins holds the string denoting the admins edge name in mutations.
	EdgeAdmins = "admins"
	// EdgeVouchers holds the string denoting the vouchers edge name in mutations.
	EdgeVouchers = "vouchers"
	// EdgeAccessories holds the string denoting the accessories edge name in mutations.
	EdgeAccessories = "accessories"
	// EdgeDepartments holds the string denoting the departments edge name in mutations.
	EdgeDepartments = "departments"
	// Table holds the table name of the institution in the database.
	Table = "institutions"
	// AdminsTable is the table that holds the admins relation/edge. The primary key declared below.
	AdminsTable = "institution_admins"
	// AdminsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AdminsInverseTable = "users"
	// VouchersTable is the table that holds the vouchers relation/edge.
	VouchersTable = "vouchers"
	// VouchersInverseTable is the table name for the Voucher entity.
	// It exists in this package in order to avoid circular dependency with the "voucher" package.
	VouchersInverseTable = "vouchers"
	// VouchersColumn is the table column denoting the vouchers relation/edge.
	VouchersColumn = "institution_vouchers"
	// AccessoriesTable is the table that holds the accessories relation/edge.
	AccessoriesTable = "accessories"
	// AccessoriesInverseTable is the table name for the Accessory entity.
	// It exists in this package in order to avoid circular dependency with the "accessory" package.
	AccessoriesInverseTable = "accessories"
	// AccessoriesColumn is the table column denoting the accessories relation/edge.
	AccessoriesColumn = "institution_accessories"
	// DepartmentsTable is the table that holds the departments relation/edge.
	DepartmentsTable = "departments"
	// DepartmentsInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentsInverseTable = "departments"
	// DepartmentsColumn is the table column denoting the departments relation/edge.
	DepartmentsColumn = "institution_departments"
)

// Columns holds all SQL columns for institution fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldShortName,
}

var (
	// AdminsPrimaryKey and AdminsColumn2 are the table columns denoting the
	// primary key for the admins relation (M2M).
	AdminsPrimaryKey = []string{"institution_id", "user_id"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// ShortNameValidator is a validator for the "short_name" field. It is called by the builders before save.
	ShortNameValidator func(string) error
)

// OrderOption defines the ordering options for the Institution queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByShortName orders the results by the short_name field.
func ByShortName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldShortName, opts...).ToFunc()
}

// ByAdminsCount orders the results by admins count.
func ByAdminsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAdminsStep(), opts...)
	}
}

// ByAdmins orders the results by admins terms.
func ByAdmins(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAdminsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByVouchersCount orders the results by vouchers count.
func ByVouchersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newVouchersStep(), opts...)
	}
}

// ByVouchers orders the results by vouchers terms.
func ByVouchers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newVouchersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAccessoriesCount orders the results by accessories count.
func ByAccessoriesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAccessoriesStep(), opts...)
	}
}

// ByAccessories orders the results by accessories terms.
func ByAccessories(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccessoriesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDepartmentsCount orders the results by departments count.
func ByDepartmentsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDepartmentsStep(), opts...)
	}
}

// ByDepartments orders the results by departments terms.
func ByDepartments(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDepartmentsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAdminsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AdminsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AdminsTable, AdminsPrimaryKey...),
	)
}
func newVouchersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(VouchersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, VouchersTable, VouchersColumn),
	)
}
func newAccessoriesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccessoriesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AccessoriesTable, AccessoriesColumn),
	)
}
func newDepartmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DepartmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DepartmentsTable, DepartmentsColumn),
	)
}
