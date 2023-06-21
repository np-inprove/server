// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/forum"
	"github.com/np-inprove/server/internal/ent/forumpost"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/reaction"
	"github.com/np-inprove/server/internal/ent/user"
)

// ForumPostQuery is the builder for querying ForumPost entities.
type ForumPostQuery struct {
	config
	ctx              *QueryContext
	order            []forumpost.OrderOption
	inters           []Interceptor
	predicates       []predicate.ForumPost
	withAuthor       *UserQuery
	withForum        *ForumQuery
	withParent       *ForumPostQuery
	withChildren     *ForumPostQuery
	withReactedUsers *UserQuery
	withReactions    *ReactionQuery
	withFKs          bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ForumPostQuery builder.
func (fpq *ForumPostQuery) Where(ps ...predicate.ForumPost) *ForumPostQuery {
	fpq.predicates = append(fpq.predicates, ps...)
	return fpq
}

// Limit the number of records to be returned by this query.
func (fpq *ForumPostQuery) Limit(limit int) *ForumPostQuery {
	fpq.ctx.Limit = &limit
	return fpq
}

// Offset to start from.
func (fpq *ForumPostQuery) Offset(offset int) *ForumPostQuery {
	fpq.ctx.Offset = &offset
	return fpq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fpq *ForumPostQuery) Unique(unique bool) *ForumPostQuery {
	fpq.ctx.Unique = &unique
	return fpq
}

// Order specifies how the records should be ordered.
func (fpq *ForumPostQuery) Order(o ...forumpost.OrderOption) *ForumPostQuery {
	fpq.order = append(fpq.order, o...)
	return fpq
}

// QueryAuthor chains the current query on the "author" edge.
func (fpq *ForumPostQuery) QueryAuthor() *UserQuery {
	query := (&UserClient{config: fpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(forumpost.Table, forumpost.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, forumpost.AuthorTable, forumpost.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryForum chains the current query on the "forum" edge.
func (fpq *ForumPostQuery) QueryForum() *ForumQuery {
	query := (&ForumClient{config: fpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(forumpost.Table, forumpost.FieldID, selector),
			sqlgraph.To(forum.Table, forum.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, forumpost.ForumTable, forumpost.ForumColumn),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParent chains the current query on the "parent" edge.
func (fpq *ForumPostQuery) QueryParent() *ForumPostQuery {
	query := (&ForumPostClient{config: fpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(forumpost.Table, forumpost.FieldID, selector),
			sqlgraph.To(forumpost.Table, forumpost.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, forumpost.ParentTable, forumpost.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (fpq *ForumPostQuery) QueryChildren() *ForumPostQuery {
	query := (&ForumPostClient{config: fpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(forumpost.Table, forumpost.FieldID, selector),
			sqlgraph.To(forumpost.Table, forumpost.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, forumpost.ChildrenTable, forumpost.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReactedUsers chains the current query on the "reacted_users" edge.
func (fpq *ForumPostQuery) QueryReactedUsers() *UserQuery {
	query := (&UserClient{config: fpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(forumpost.Table, forumpost.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, forumpost.ReactedUsersTable, forumpost.ReactedUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReactions chains the current query on the "reactions" edge.
func (fpq *ForumPostQuery) QueryReactions() *ReactionQuery {
	query := (&ReactionClient{config: fpq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fpq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fpq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(forumpost.Table, forumpost.FieldID, selector),
			sqlgraph.To(reaction.Table, reaction.ForumPostColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, forumpost.ReactionsTable, forumpost.ReactionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fpq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ForumPost entity from the query.
// Returns a *NotFoundError when no ForumPost was found.
func (fpq *ForumPostQuery) First(ctx context.Context) (*ForumPost, error) {
	nodes, err := fpq.Limit(1).All(setContextOp(ctx, fpq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{forumpost.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fpq *ForumPostQuery) FirstX(ctx context.Context) *ForumPost {
	node, err := fpq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ForumPost ID from the query.
// Returns a *NotFoundError when no ForumPost ID was found.
func (fpq *ForumPostQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fpq.Limit(1).IDs(setContextOp(ctx, fpq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{forumpost.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fpq *ForumPostQuery) FirstIDX(ctx context.Context) int {
	id, err := fpq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ForumPost entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ForumPost entity is found.
// Returns a *NotFoundError when no ForumPost entities are found.
func (fpq *ForumPostQuery) Only(ctx context.Context) (*ForumPost, error) {
	nodes, err := fpq.Limit(2).All(setContextOp(ctx, fpq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{forumpost.Label}
	default:
		return nil, &NotSingularError{forumpost.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fpq *ForumPostQuery) OnlyX(ctx context.Context) *ForumPost {
	node, err := fpq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ForumPost ID in the query.
// Returns a *NotSingularError when more than one ForumPost ID is found.
// Returns a *NotFoundError when no entities are found.
func (fpq *ForumPostQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fpq.Limit(2).IDs(setContextOp(ctx, fpq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{forumpost.Label}
	default:
		err = &NotSingularError{forumpost.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fpq *ForumPostQuery) OnlyIDX(ctx context.Context) int {
	id, err := fpq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ForumPosts.
func (fpq *ForumPostQuery) All(ctx context.Context) ([]*ForumPost, error) {
	ctx = setContextOp(ctx, fpq.ctx, "All")
	if err := fpq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ForumPost, *ForumPostQuery]()
	return withInterceptors[[]*ForumPost](ctx, fpq, qr, fpq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fpq *ForumPostQuery) AllX(ctx context.Context) []*ForumPost {
	nodes, err := fpq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ForumPost IDs.
func (fpq *ForumPostQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fpq.ctx.Unique == nil && fpq.path != nil {
		fpq.Unique(true)
	}
	ctx = setContextOp(ctx, fpq.ctx, "IDs")
	if err = fpq.Select(forumpost.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fpq *ForumPostQuery) IDsX(ctx context.Context) []int {
	ids, err := fpq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fpq *ForumPostQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fpq.ctx, "Count")
	if err := fpq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fpq, querierCount[*ForumPostQuery](), fpq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fpq *ForumPostQuery) CountX(ctx context.Context) int {
	count, err := fpq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fpq *ForumPostQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fpq.ctx, "Exist")
	switch _, err := fpq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fpq *ForumPostQuery) ExistX(ctx context.Context) bool {
	exist, err := fpq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ForumPostQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fpq *ForumPostQuery) Clone() *ForumPostQuery {
	if fpq == nil {
		return nil
	}
	return &ForumPostQuery{
		config:           fpq.config,
		ctx:              fpq.ctx.Clone(),
		order:            append([]forumpost.OrderOption{}, fpq.order...),
		inters:           append([]Interceptor{}, fpq.inters...),
		predicates:       append([]predicate.ForumPost{}, fpq.predicates...),
		withAuthor:       fpq.withAuthor.Clone(),
		withForum:        fpq.withForum.Clone(),
		withParent:       fpq.withParent.Clone(),
		withChildren:     fpq.withChildren.Clone(),
		withReactedUsers: fpq.withReactedUsers.Clone(),
		withReactions:    fpq.withReactions.Clone(),
		// clone intermediate query.
		sql:  fpq.sql.Clone(),
		path: fpq.path,
	}
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *ForumPostQuery) WithAuthor(opts ...func(*UserQuery)) *ForumPostQuery {
	query := (&UserClient{config: fpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fpq.withAuthor = query
	return fpq
}

// WithForum tells the query-builder to eager-load the nodes that are connected to
// the "forum" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *ForumPostQuery) WithForum(opts ...func(*ForumQuery)) *ForumPostQuery {
	query := (&ForumClient{config: fpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fpq.withForum = query
	return fpq
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *ForumPostQuery) WithParent(opts ...func(*ForumPostQuery)) *ForumPostQuery {
	query := (&ForumPostClient{config: fpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fpq.withParent = query
	return fpq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *ForumPostQuery) WithChildren(opts ...func(*ForumPostQuery)) *ForumPostQuery {
	query := (&ForumPostClient{config: fpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fpq.withChildren = query
	return fpq
}

// WithReactedUsers tells the query-builder to eager-load the nodes that are connected to
// the "reacted_users" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *ForumPostQuery) WithReactedUsers(opts ...func(*UserQuery)) *ForumPostQuery {
	query := (&UserClient{config: fpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fpq.withReactedUsers = query
	return fpq
}

// WithReactions tells the query-builder to eager-load the nodes that are connected to
// the "reactions" edge. The optional arguments are used to configure the query builder of the edge.
func (fpq *ForumPostQuery) WithReactions(opts ...func(*ReactionQuery)) *ForumPostQuery {
	query := (&ReactionClient{config: fpq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fpq.withReactions = query
	return fpq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ForumPost.Query().
//		GroupBy(forumpost.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fpq *ForumPostQuery) GroupBy(field string, fields ...string) *ForumPostGroupBy {
	fpq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ForumPostGroupBy{build: fpq}
	grbuild.flds = &fpq.ctx.Fields
	grbuild.label = forumpost.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.ForumPost.Query().
//		Select(forumpost.FieldTitle).
//		Scan(ctx, &v)
func (fpq *ForumPostQuery) Select(fields ...string) *ForumPostSelect {
	fpq.ctx.Fields = append(fpq.ctx.Fields, fields...)
	sbuild := &ForumPostSelect{ForumPostQuery: fpq}
	sbuild.label = forumpost.Label
	sbuild.flds, sbuild.scan = &fpq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ForumPostSelect configured with the given aggregations.
func (fpq *ForumPostQuery) Aggregate(fns ...AggregateFunc) *ForumPostSelect {
	return fpq.Select().Aggregate(fns...)
}

func (fpq *ForumPostQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fpq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fpq); err != nil {
				return err
			}
		}
	}
	for _, f := range fpq.ctx.Fields {
		if !forumpost.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fpq.path != nil {
		prev, err := fpq.path(ctx)
		if err != nil {
			return err
		}
		fpq.sql = prev
	}
	return nil
}

func (fpq *ForumPostQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ForumPost, error) {
	var (
		nodes       = []*ForumPost{}
		withFKs     = fpq.withFKs
		_spec       = fpq.querySpec()
		loadedTypes = [6]bool{
			fpq.withAuthor != nil,
			fpq.withForum != nil,
			fpq.withParent != nil,
			fpq.withChildren != nil,
			fpq.withReactedUsers != nil,
			fpq.withReactions != nil,
		}
	)
	if fpq.withAuthor != nil || fpq.withForum != nil || fpq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, forumpost.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ForumPost).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ForumPost{config: fpq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fpq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fpq.withAuthor; query != nil {
		if err := fpq.loadAuthor(ctx, query, nodes, nil,
			func(n *ForumPost, e *User) { n.Edges.Author = e }); err != nil {
			return nil, err
		}
	}
	if query := fpq.withForum; query != nil {
		if err := fpq.loadForum(ctx, query, nodes, nil,
			func(n *ForumPost, e *Forum) { n.Edges.Forum = e }); err != nil {
			return nil, err
		}
	}
	if query := fpq.withParent; query != nil {
		if err := fpq.loadParent(ctx, query, nodes, nil,
			func(n *ForumPost, e *ForumPost) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := fpq.withChildren; query != nil {
		if err := fpq.loadChildren(ctx, query, nodes,
			func(n *ForumPost) { n.Edges.Children = []*ForumPost{} },
			func(n *ForumPost, e *ForumPost) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	if query := fpq.withReactedUsers; query != nil {
		if err := fpq.loadReactedUsers(ctx, query, nodes,
			func(n *ForumPost) { n.Edges.ReactedUsers = []*User{} },
			func(n *ForumPost, e *User) { n.Edges.ReactedUsers = append(n.Edges.ReactedUsers, e) }); err != nil {
			return nil, err
		}
	}
	if query := fpq.withReactions; query != nil {
		if err := fpq.loadReactions(ctx, query, nodes,
			func(n *ForumPost) { n.Edges.Reactions = []*Reaction{} },
			func(n *ForumPost, e *Reaction) { n.Edges.Reactions = append(n.Edges.Reactions, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fpq *ForumPostQuery) loadAuthor(ctx context.Context, query *UserQuery, nodes []*ForumPost, init func(*ForumPost), assign func(*ForumPost, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ForumPost)
	for i := range nodes {
		if nodes[i].forum_post_author == nil {
			continue
		}
		fk := *nodes[i].forum_post_author
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "forum_post_author" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fpq *ForumPostQuery) loadForum(ctx context.Context, query *ForumQuery, nodes []*ForumPost, init func(*ForumPost), assign func(*ForumPost, *Forum)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ForumPost)
	for i := range nodes {
		if nodes[i].forum_posts == nil {
			continue
		}
		fk := *nodes[i].forum_posts
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(forum.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "forum_posts" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fpq *ForumPostQuery) loadParent(ctx context.Context, query *ForumPostQuery, nodes []*ForumPost, init func(*ForumPost), assign func(*ForumPost, *ForumPost)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ForumPost)
	for i := range nodes {
		if nodes[i].forum_post_children == nil {
			continue
		}
		fk := *nodes[i].forum_post_children
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(forumpost.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "forum_post_children" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fpq *ForumPostQuery) loadChildren(ctx context.Context, query *ForumPostQuery, nodes []*ForumPost, init func(*ForumPost), assign func(*ForumPost, *ForumPost)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ForumPost)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ForumPost(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(forumpost.ChildrenColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.forum_post_children
		if fk == nil {
			return fmt.Errorf(`foreign-key "forum_post_children" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "forum_post_children" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (fpq *ForumPostQuery) loadReactedUsers(ctx context.Context, query *UserQuery, nodes []*ForumPost, init func(*ForumPost), assign func(*ForumPost, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*ForumPost)
	nids := make(map[int]map[*ForumPost]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(forumpost.ReactedUsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(forumpost.ReactedUsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(forumpost.ReactedUsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(forumpost.ReactedUsersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*ForumPost]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*User](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "reacted_users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (fpq *ForumPostQuery) loadReactions(ctx context.Context, query *ReactionQuery, nodes []*ForumPost, init func(*ForumPost), assign func(*ForumPost, *Reaction)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ForumPost)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(reaction.FieldForumPostID)
	}
	query.Where(predicate.Reaction(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(forumpost.ReactionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ForumPostID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "forum_post_id" returned %v for node %v`, fk, n)
		}
		assign(node, n)
	}
	return nil
}

func (fpq *ForumPostQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fpq.querySpec()
	_spec.Node.Columns = fpq.ctx.Fields
	if len(fpq.ctx.Fields) > 0 {
		_spec.Unique = fpq.ctx.Unique != nil && *fpq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fpq.driver, _spec)
}

func (fpq *ForumPostQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(forumpost.Table, forumpost.Columns, sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt))
	_spec.From = fpq.sql
	if unique := fpq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fpq.path != nil {
		_spec.Unique = true
	}
	if fields := fpq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, forumpost.FieldID)
		for i := range fields {
			if fields[i] != forumpost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fpq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fpq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fpq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fpq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fpq *ForumPostQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fpq.driver.Dialect())
	t1 := builder.Table(forumpost.Table)
	columns := fpq.ctx.Fields
	if len(columns) == 0 {
		columns = forumpost.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fpq.sql != nil {
		selector = fpq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fpq.ctx.Unique != nil && *fpq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fpq.predicates {
		p(selector)
	}
	for _, p := range fpq.order {
		p(selector)
	}
	if offset := fpq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fpq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForumPostGroupBy is the group-by builder for ForumPost entities.
type ForumPostGroupBy struct {
	selector
	build *ForumPostQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fpgb *ForumPostGroupBy) Aggregate(fns ...AggregateFunc) *ForumPostGroupBy {
	fpgb.fns = append(fpgb.fns, fns...)
	return fpgb
}

// Scan applies the selector query and scans the result into the given value.
func (fpgb *ForumPostGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fpgb.build.ctx, "GroupBy")
	if err := fpgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ForumPostQuery, *ForumPostGroupBy](ctx, fpgb.build, fpgb, fpgb.build.inters, v)
}

func (fpgb *ForumPostGroupBy) sqlScan(ctx context.Context, root *ForumPostQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fpgb.fns))
	for _, fn := range fpgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fpgb.flds)+len(fpgb.fns))
		for _, f := range *fpgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fpgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fpgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ForumPostSelect is the builder for selecting fields of ForumPost entities.
type ForumPostSelect struct {
	*ForumPostQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fps *ForumPostSelect) Aggregate(fns ...AggregateFunc) *ForumPostSelect {
	fps.fns = append(fps.fns, fns...)
	return fps
}

// Scan applies the selector query and scans the result into the given value.
func (fps *ForumPostSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fps.ctx, "Select")
	if err := fps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ForumPostQuery, *ForumPostSelect](ctx, fps.ForumPostQuery, fps, fps.inters, v)
}

func (fps *ForumPostSelect) sqlScan(ctx context.Context, root *ForumPostQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fps.fns))
	for _, fn := range fps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
