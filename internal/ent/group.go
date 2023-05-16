// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/np-inprove/server/internal/ent/group"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL path of the group (example: csf01-2023)
	Path string `json:"path,omitempty"`
	// Name of the group (example: CSF01 2023)
	Name string `json:"name,omitempty"`
	// Description of the group
	Description string `json:"description,omitempty"`
	// GroupType holds the value of the "group_type" field.
	GroupType group.GroupType `json:"group_type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges        GroupEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Users that are in the group
	Users []*User `json:"users,omitempty"`
	// Events from the group
	Events []*Event `json:"events,omitempty"`
	// Forum posts from the group
	ForumPosts []*ForumPost `json:"forum_posts,omitempty"`
	// Deadlines created by users from the group
	Deadlines []*Deadline `json:"deadlines,omitempty"`
	// GroupUsers holds the value of the group_users edge.
	GroupUsers []*GroupUser `json:"group_users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// EventsOrErr returns the Events value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) EventsOrErr() ([]*Event, error) {
	if e.loadedTypes[1] {
		return e.Events, nil
	}
	return nil, &NotLoadedError{edge: "events"}
}

// ForumPostsOrErr returns the ForumPosts value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) ForumPostsOrErr() ([]*ForumPost, error) {
	if e.loadedTypes[2] {
		return e.ForumPosts, nil
	}
	return nil, &NotLoadedError{edge: "forum_posts"}
}

// DeadlinesOrErr returns the Deadlines value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) DeadlinesOrErr() ([]*Deadline, error) {
	if e.loadedTypes[3] {
		return e.Deadlines, nil
	}
	return nil, &NotLoadedError{edge: "deadlines"}
}

// GroupUsersOrErr returns the GroupUsers value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) GroupUsersOrErr() ([]*GroupUser, error) {
	if e.loadedTypes[4] {
		return e.GroupUsers, nil
	}
	return nil, &NotLoadedError{edge: "group_users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			values[i] = new(sql.NullInt64)
		case group.FieldPath, group.FieldName, group.FieldDescription, group.FieldGroupType:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (gr *Group) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gr.ID = int(value.Int64)
		case group.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				gr.Path = value.String
			}
		case group.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case group.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				gr.Description = value.String
			}
		case group.FieldGroupType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_type", values[i])
			} else if value.Valid {
				gr.GroupType = group.GroupType(value.String)
			}
		default:
			gr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Group.
// This includes values selected through modifiers, order, etc.
func (gr *Group) Value(name string) (ent.Value, error) {
	return gr.selectValues.Get(name)
}

// QueryUsers queries the "users" edge of the Group entity.
func (gr *Group) QueryUsers() *UserQuery {
	return NewGroupClient(gr.config).QueryUsers(gr)
}

// QueryEvents queries the "events" edge of the Group entity.
func (gr *Group) QueryEvents() *EventQuery {
	return NewGroupClient(gr.config).QueryEvents(gr)
}

// QueryForumPosts queries the "forum_posts" edge of the Group entity.
func (gr *Group) QueryForumPosts() *ForumPostQuery {
	return NewGroupClient(gr.config).QueryForumPosts(gr)
}

// QueryDeadlines queries the "deadlines" edge of the Group entity.
func (gr *Group) QueryDeadlines() *DeadlineQuery {
	return NewGroupClient(gr.config).QueryDeadlines(gr)
}

// QueryGroupUsers queries the "group_users" edge of the Group entity.
func (gr *Group) QueryGroupUsers() *GroupUserQuery {
	return NewGroupClient(gr.config).QueryGroupUsers(gr)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Group) Update() *GroupUpdateOne {
	return NewGroupClient(gr.config).UpdateOne(gr)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Group) Unwrap() *Group {
	_tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	gr.config.driver = _tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gr.ID))
	builder.WriteString("path=")
	builder.WriteString(gr.Path)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(gr.Description)
	builder.WriteString(", ")
	builder.WriteString("group_type=")
	builder.WriteString(fmt.Sprintf("%v", gr.GroupType))
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group