// Code generated by ent, DO NOT EDIT.

package groupuser

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/np-inprove/server/internal/entity/group"
)

const (
	// Label holds the string label denoting the groupuser type in the database.
	Label = "group_user"
	// FieldGroupID holds the string denoting the group_id field in the database.
	FieldGroupID = "group_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// GroupFieldID holds the string denoting the ID field of the Group.
	GroupFieldID = "id"
	// UserFieldID holds the string denoting the ID field of the User.
	UserFieldID = "id"
	// Table holds the table name of the groupuser in the database.
	Table = "group_users"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "group_users"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "entgroup" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_id"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "group_users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
)

// Columns holds all SQL columns for groupuser fields.
var Columns = []string{
	FieldGroupID,
	FieldUserID,
	FieldRole,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r group.Role) error {
	switch r {
	case "owner", "educator", "member":
		return nil
	default:
		return fmt.Errorf("groupuser: invalid enum value for role field: %q", r)
	}
}

// OrderOption defines the ordering options for the GroupUser queries.
type OrderOption func(*sql.Selector)

// ByGroupID orders the results by the group_id field.
func ByGroupID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroupID, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByGroupField orders the results by group field.
func ByGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupStep(), sql.OrderByField(field, opts...))
	}
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, GroupColumn),
		sqlgraph.To(GroupInverseTable, GroupFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
	)
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, UserColumn),
		sqlgraph.To(UserInverseTable, UserFieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
	)
}
