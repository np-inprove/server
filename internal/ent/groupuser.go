// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/groupuser"
	"github.com/np-inprove/server/internal/ent/user"
)

// GroupUser is the model entity for the GroupUser schema.
type GroupUser struct {
	config `json:"-"`
	// ID of the group
	GroupID int `json:"group_id,omitempty"`
	// ID of the user
	UserID int `json:"user_id,omitempty"`
	// Role holds the value of the "role" field.
	Role groupuser.Role `json:"role,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupUserQuery when eager-loading is set.
	Edges        GroupUserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupUserEdges holds the relations/edges for other nodes in the graph.
type GroupUserEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupUserEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[0] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupUserEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupuser.FieldGroupID, groupuser.FieldUserID:
			values[i] = new(sql.NullInt64)
		case groupuser.FieldRole:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupUser fields.
func (gu *GroupUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupuser.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				gu.GroupID = int(value.Int64)
			}
		case groupuser.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				gu.UserID = int(value.Int64)
			}
		case groupuser.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				gu.Role = groupuser.Role(value.String)
			}
		default:
			gu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupUser.
// This includes values selected through modifiers, order, etc.
func (gu *GroupUser) Value(name string) (ent.Value, error) {
	return gu.selectValues.Get(name)
}

// QueryGroup queries the "group" edge of the GroupUser entity.
func (gu *GroupUser) QueryGroup() *GroupQuery {
	return NewGroupUserClient(gu.config).QueryGroup(gu)
}

// QueryUser queries the "user" edge of the GroupUser entity.
func (gu *GroupUser) QueryUser() *UserQuery {
	return NewGroupUserClient(gu.config).QueryUser(gu)
}

// Update returns a builder for updating this GroupUser.
// Note that you need to call GroupUser.Unwrap() before calling this method if this GroupUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (gu *GroupUser) Update() *GroupUserUpdateOne {
	return NewGroupUserClient(gu.config).UpdateOne(gu)
}

// Unwrap unwraps the GroupUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gu *GroupUser) Unwrap() *GroupUser {
	_tx, ok := gu.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupUser is not a transactional entity")
	}
	gu.config.driver = _tx.drv
	return gu
}

// String implements the fmt.Stringer.
func (gu *GroupUser) String() string {
	var builder strings.Builder
	builder.WriteString("GroupUser(")
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", gu.GroupID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", gu.UserID))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", gu.Role))
	builder.WriteByte(')')
	return builder.String()
}

// GroupUsers is a parsable slice of GroupUser.
type GroupUsers []*GroupUser
