// Code generated by ent, DO NOT EDIT.

package group

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldGroupType holds the string denoting the group_type field in the database.
	FieldGroupType = "group_type"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeEvents holds the string denoting the events edge name in mutations.
	EdgeEvents = "events"
	// EdgeForumPosts holds the string denoting the forum_posts edge name in mutations.
	EdgeForumPosts = "forum_posts"
	// EdgeDeadlines holds the string denoting the deadlines edge name in mutations.
	EdgeDeadlines = "deadlines"
	// EdgeGroupUsers holds the string denoting the group_users edge name in mutations.
	EdgeGroupUsers = "group_users"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "group_users"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// EventsTable is the table that holds the events relation/edge.
	EventsTable = "events"
	// EventsInverseTable is the table name for the Event entity.
	// It exists in this package in order to avoid circular dependency with the "event" package.
	EventsInverseTable = "events"
	// EventsColumn is the table column denoting the events relation/edge.
	EventsColumn = "group_events"
	// ForumPostsTable is the table that holds the forum_posts relation/edge.
	ForumPostsTable = "forum_posts"
	// ForumPostsInverseTable is the table name for the ForumPost entity.
	// It exists in this package in order to avoid circular dependency with the "forumpost" package.
	ForumPostsInverseTable = "forum_posts"
	// ForumPostsColumn is the table column denoting the forum_posts relation/edge.
	ForumPostsColumn = "group_forum_posts"
	// DeadlinesTable is the table that holds the deadlines relation/edge.
	DeadlinesTable = "deadlines"
	// DeadlinesInverseTable is the table name for the Deadline entity.
	// It exists in this package in order to avoid circular dependency with the "deadline" package.
	DeadlinesInverseTable = "deadlines"
	// DeadlinesColumn is the table column denoting the deadlines relation/edge.
	DeadlinesColumn = "group_deadlines"
	// GroupUsersTable is the table that holds the group_users relation/edge.
	GroupUsersTable = "group_users"
	// GroupUsersInverseTable is the table name for the GroupUser entity.
	// It exists in this package in order to avoid circular dependency with the "groupuser" package.
	GroupUsersInverseTable = "group_users"
	// GroupUsersColumn is the table column denoting the group_users relation/edge.
	GroupUsersColumn = "group_id"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldPath,
	FieldName,
	FieldDescription,
	FieldGroupType,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"group_id", "user_id"}
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
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// GroupType defines the type for the "group_type" enum field.
type GroupType string

// GroupType values.
const (
	GroupTypeSpecialInterestGroup GroupType = "special_interest_group"
	GroupTypeModuleGroup          GroupType = "module_group"
	GroupTypeCca                  GroupType = "cca"
	GroupTypeMentorClass          GroupType = "mentor_class"
)

func (gt GroupType) String() string {
	return string(gt)
}

// GroupTypeValidator is a validator for the "group_type" field enum values. It is called by the builders before save.
func GroupTypeValidator(gt GroupType) error {
	switch gt {
	case GroupTypeSpecialInterestGroup, GroupTypeModuleGroup, GroupTypeCca, GroupTypeMentorClass:
		return nil
	default:
		return fmt.Errorf("group: invalid enum value for group_type field: %q", gt)
	}
}

// OrderOption defines the ordering options for the Group queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPath orders the results by the path field.
func ByPath(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPath, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByGroupType orders the results by the group_type field.
func ByGroupType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldGroupType, opts...).ToFunc()
}

// ByUsersCount orders the results by users count.
func ByUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUsersStep(), opts...)
	}
}

// ByUsers orders the results by users terms.
func ByUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByEventsCount orders the results by events count.
func ByEventsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEventsStep(), opts...)
	}
}

// ByEvents orders the results by events terms.
func ByEvents(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEventsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByForumPostsCount orders the results by forum_posts count.
func ByForumPostsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newForumPostsStep(), opts...)
	}
}

// ByForumPosts orders the results by forum_posts terms.
func ByForumPosts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newForumPostsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDeadlinesCount orders the results by deadlines count.
func ByDeadlinesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newDeadlinesStep(), opts...)
	}
}

// ByDeadlines orders the results by deadlines terms.
func ByDeadlines(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeadlinesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGroupUsersCount orders the results by group_users count.
func ByGroupUsersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGroupUsersStep(), opts...)
	}
}

// ByGroupUsers orders the results by group_users terms.
func ByGroupUsers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGroupUsersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UsersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, UsersTable, UsersPrimaryKey...),
	)
}
func newEventsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EventsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EventsTable, EventsColumn),
	)
}
func newForumPostsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ForumPostsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ForumPostsTable, ForumPostsColumn),
	)
}
func newDeadlinesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeadlinesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, DeadlinesTable, DeadlinesColumn),
	)
}
func newGroupUsersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GroupUsersInverseTable, GroupUsersColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, GroupUsersTable, GroupUsersColumn),
	)
}
