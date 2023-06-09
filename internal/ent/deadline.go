// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/np-inprove/server/internal/ent/deadline"
	entgroup "github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/user"
)

// Deadline is the model entity for the Deadline schema.
type Deadline struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name of the deadline
	Name string `json:"name,omitempty"`
	// Time when the deadline is due.
	// If there is no due time, set to nil
	DueTime *time.Time `json:"due_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeadlineQuery when eager-loading is set.
	Edges           DeadlineEdges `json:"edges"`
	deadline_author *int
	group_deadlines *int
	selectValues    sql.SelectValues
}

// DeadlineEdges holds the relations/edges for other nodes in the graph.
type DeadlineEdges struct {
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// Users who voted for this deadline
	VotedUsers []*User `json:"voted_users,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeadlineEdges) AuthorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Author == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// VotedUsersOrErr returns the VotedUsers value or an error if the edge
// was not loaded in eager-loading.
func (e DeadlineEdges) VotedUsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.VotedUsers, nil
	}
	return nil, &NotLoadedError{edge: "voted_users"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeadlineEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[2] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: entgroup.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Deadline) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case deadline.FieldID:
			values[i] = new(sql.NullInt64)
		case deadline.FieldName:
			values[i] = new(sql.NullString)
		case deadline.FieldDueTime:
			values[i] = new(sql.NullTime)
		case deadline.ForeignKeys[0]: // deadline_author
			values[i] = new(sql.NullInt64)
		case deadline.ForeignKeys[1]: // group_deadlines
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Deadline fields.
func (d *Deadline) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case deadline.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			d.ID = int(value.Int64)
		case deadline.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				d.Name = value.String
			}
		case deadline.FieldDueTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field due_time", values[i])
			} else if value.Valid {
				d.DueTime = new(time.Time)
				*d.DueTime = value.Time
			}
		case deadline.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field deadline_author", value)
			} else if value.Valid {
				d.deadline_author = new(int)
				*d.deadline_author = int(value.Int64)
			}
		case deadline.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_deadlines", value)
			} else if value.Valid {
				d.group_deadlines = new(int)
				*d.group_deadlines = int(value.Int64)
			}
		default:
			d.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Deadline.
// This includes values selected through modifiers, order, etc.
func (d *Deadline) Value(name string) (ent.Value, error) {
	return d.selectValues.Get(name)
}

// QueryAuthor queries the "author" edge of the Deadline entity.
func (d *Deadline) QueryAuthor() *UserQuery {
	return NewDeadlineClient(d.config).QueryAuthor(d)
}

// QueryVotedUsers queries the "voted_users" edge of the Deadline entity.
func (d *Deadline) QueryVotedUsers() *UserQuery {
	return NewDeadlineClient(d.config).QueryVotedUsers(d)
}

// QueryGroup queries the "group" edge of the Deadline entity.
func (d *Deadline) QueryGroup() *GroupQuery {
	return NewDeadlineClient(d.config).QueryGroup(d)
}

// Update returns a builder for updating this Deadline.
// Note that you need to call Deadline.Unwrap() before calling this method if this Deadline
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Deadline) Update() *DeadlineUpdateOne {
	return NewDeadlineClient(d.config).UpdateOne(d)
}

// Unwrap unwraps the Deadline entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Deadline) Unwrap() *Deadline {
	_tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Deadline is not a transactional entity")
	}
	d.config.driver = _tx.drv
	return d
}

// String implements the fmt.Stringer.
func (d *Deadline) String() string {
	var builder strings.Builder
	builder.WriteString("Deadline(")
	builder.WriteString(fmt.Sprintf("id=%v, ", d.ID))
	builder.WriteString("name=")
	builder.WriteString(d.Name)
	builder.WriteString(", ")
	if v := d.DueTime; v != nil {
		builder.WriteString("due_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Deadlines is a parsable slice of Deadline.
type Deadlines []*Deadline
