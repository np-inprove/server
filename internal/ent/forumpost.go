// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/np-inprove/server/internal/ent/forumpost"
	entgroup "github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/user"
)

// ForumPost is the model entity for the ForumPost schema.
type ForumPost struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title of the forum post
	Title string `json:"title,omitempty"`
	// Content of the forum post
	Content string `json:"content,omitempty"`
	// Slice of user IDs that were mentioned
	MentionedUsersJSON []int `json:"mentioned_users_json,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ForumPostQuery when eager-loading is set.
	Edges               ForumPostEdges `json:"edges"`
	forum_post_author   *int
	forum_post_children *int
	group_forum_posts   *int
	selectValues        sql.SelectValues
}

// ForumPostEdges holds the relations/edges for other nodes in the graph.
type ForumPostEdges struct {
	// Author of the forum post
	Author *User `json:"author,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *ForumPost `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*ForumPost `json:"children,omitempty"`
	// ReactedUsers holds the value of the reacted_users edge.
	ReactedUsers []*User `json:"reacted_users,omitempty"`
	// Reactions holds the value of the reactions edge.
	Reactions []*Reaction `json:"reactions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ForumPostEdges) AuthorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Author == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ForumPostEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: entgroup.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ForumPostEdges) ParentOrErr() (*ForumPost, error) {
	if e.loadedTypes[2] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: forumpost.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e ForumPostEdges) ChildrenOrErr() ([]*ForumPost, error) {
	if e.loadedTypes[3] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ReactedUsersOrErr returns the ReactedUsers value or an error if the edge
// was not loaded in eager-loading.
func (e ForumPostEdges) ReactedUsersOrErr() ([]*User, error) {
	if e.loadedTypes[4] {
		return e.ReactedUsers, nil
	}
	return nil, &NotLoadedError{edge: "reacted_users"}
}

// ReactionsOrErr returns the Reactions value or an error if the edge
// was not loaded in eager-loading.
func (e ForumPostEdges) ReactionsOrErr() ([]*Reaction, error) {
	if e.loadedTypes[5] {
		return e.Reactions, nil
	}
	return nil, &NotLoadedError{edge: "reactions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ForumPost) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case forumpost.FieldMentionedUsersJSON:
			values[i] = new([]byte)
		case forumpost.FieldID:
			values[i] = new(sql.NullInt64)
		case forumpost.FieldTitle, forumpost.FieldContent:
			values[i] = new(sql.NullString)
		case forumpost.ForeignKeys[0]: // forum_post_author
			values[i] = new(sql.NullInt64)
		case forumpost.ForeignKeys[1]: // forum_post_children
			values[i] = new(sql.NullInt64)
		case forumpost.ForeignKeys[2]: // group_forum_posts
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ForumPost fields.
func (fp *ForumPost) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case forumpost.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fp.ID = int(value.Int64)
		case forumpost.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				fp.Title = value.String
			}
		case forumpost.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				fp.Content = value.String
			}
		case forumpost.FieldMentionedUsersJSON:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field mentioned_users_json", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &fp.MentionedUsersJSON); err != nil {
					return fmt.Errorf("unmarshal field mentioned_users_json: %w", err)
				}
			}
		case forumpost.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field forum_post_author", value)
			} else if value.Valid {
				fp.forum_post_author = new(int)
				*fp.forum_post_author = int(value.Int64)
			}
		case forumpost.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field forum_post_children", value)
			} else if value.Valid {
				fp.forum_post_children = new(int)
				*fp.forum_post_children = int(value.Int64)
			}
		case forumpost.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field group_forum_posts", value)
			} else if value.Valid {
				fp.group_forum_posts = new(int)
				*fp.group_forum_posts = int(value.Int64)
			}
		default:
			fp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ForumPost.
// This includes values selected through modifiers, order, etc.
func (fp *ForumPost) Value(name string) (ent.Value, error) {
	return fp.selectValues.Get(name)
}

// QueryAuthor queries the "author" edge of the ForumPost entity.
func (fp *ForumPost) QueryAuthor() *UserQuery {
	return NewForumPostClient(fp.config).QueryAuthor(fp)
}

// QueryGroup queries the "group" edge of the ForumPost entity.
func (fp *ForumPost) QueryGroup() *GroupQuery {
	return NewForumPostClient(fp.config).QueryGroup(fp)
}

// QueryParent queries the "parent" edge of the ForumPost entity.
func (fp *ForumPost) QueryParent() *ForumPostQuery {
	return NewForumPostClient(fp.config).QueryParent(fp)
}

// QueryChildren queries the "children" edge of the ForumPost entity.
func (fp *ForumPost) QueryChildren() *ForumPostQuery {
	return NewForumPostClient(fp.config).QueryChildren(fp)
}

// QueryReactedUsers queries the "reacted_users" edge of the ForumPost entity.
func (fp *ForumPost) QueryReactedUsers() *UserQuery {
	return NewForumPostClient(fp.config).QueryReactedUsers(fp)
}

// QueryReactions queries the "reactions" edge of the ForumPost entity.
func (fp *ForumPost) QueryReactions() *ReactionQuery {
	return NewForumPostClient(fp.config).QueryReactions(fp)
}

// Update returns a builder for updating this ForumPost.
// Note that you need to call ForumPost.Unwrap() before calling this method if this ForumPost
// was returned from a transaction, and the transaction was committed or rolled back.
func (fp *ForumPost) Update() *ForumPostUpdateOne {
	return NewForumPostClient(fp.config).UpdateOne(fp)
}

// Unwrap unwraps the ForumPost entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fp *ForumPost) Unwrap() *ForumPost {
	_tx, ok := fp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ForumPost is not a transactional entity")
	}
	fp.config.driver = _tx.drv
	return fp
}

// String implements the fmt.Stringer.
func (fp *ForumPost) String() string {
	var builder strings.Builder
	builder.WriteString("ForumPost(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fp.ID))
	builder.WriteString("title=")
	builder.WriteString(fp.Title)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(fp.Content)
	builder.WriteString(", ")
	builder.WriteString("mentioned_users_json=")
	builder.WriteString(fmt.Sprintf("%v", fp.MentionedUsersJSON))
	builder.WriteByte(')')
	return builder.String()
}

// ForumPosts is a parsable slice of ForumPost.
type ForumPosts []*ForumPost
