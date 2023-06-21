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
	entgroup "github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/predicate"
)

// ForumQuery is the builder for querying Forum entities.
type ForumQuery struct {
	config
	ctx        *QueryContext
	order      []forum.OrderOption
	inters     []Interceptor
	predicates []predicate.Forum
	withGroup  *GroupQuery
	withPosts  *ForumPostQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ForumQuery builder.
func (fq *ForumQuery) Where(ps ...predicate.Forum) *ForumQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *ForumQuery) Limit(limit int) *ForumQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *ForumQuery) Offset(offset int) *ForumQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *ForumQuery) Unique(unique bool) *ForumQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *ForumQuery) Order(o ...forum.OrderOption) *ForumQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryGroup chains the current query on the "group" edge.
func (fq *ForumQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(forum.Table, forum.FieldID, selector),
			sqlgraph.To(entgroup.Table, entgroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, forum.GroupTable, forum.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPosts chains the current query on the "posts" edge.
func (fq *ForumQuery) QueryPosts() *ForumPostQuery {
	query := (&ForumPostClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(forum.Table, forum.FieldID, selector),
			sqlgraph.To(forumpost.Table, forumpost.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, forum.PostsTable, forum.PostsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Forum entity from the query.
// Returns a *NotFoundError when no Forum was found.
func (fq *ForumQuery) First(ctx context.Context) (*Forum, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{forum.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *ForumQuery) FirstX(ctx context.Context) *Forum {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Forum ID from the query.
// Returns a *NotFoundError when no Forum ID was found.
func (fq *ForumQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{forum.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *ForumQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Forum entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Forum entity is found.
// Returns a *NotFoundError when no Forum entities are found.
func (fq *ForumQuery) Only(ctx context.Context) (*Forum, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{forum.Label}
	default:
		return nil, &NotSingularError{forum.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *ForumQuery) OnlyX(ctx context.Context) *Forum {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Forum ID in the query.
// Returns a *NotSingularError when more than one Forum ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *ForumQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{forum.Label}
	default:
		err = &NotSingularError{forum.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *ForumQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Forums.
func (fq *ForumQuery) All(ctx context.Context) ([]*Forum, error) {
	ctx = setContextOp(ctx, fq.ctx, "All")
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Forum, *ForumQuery]()
	return withInterceptors[[]*Forum](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *ForumQuery) AllX(ctx context.Context) []*Forum {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Forum IDs.
func (fq *ForumQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, "IDs")
	if err = fq.Select(forum.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *ForumQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *ForumQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, "Count")
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*ForumQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *ForumQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *ForumQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, "Exist")
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *ForumQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ForumQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *ForumQuery) Clone() *ForumQuery {
	if fq == nil {
		return nil
	}
	return &ForumQuery{
		config:     fq.config,
		ctx:        fq.ctx.Clone(),
		order:      append([]forum.OrderOption{}, fq.order...),
		inters:     append([]Interceptor{}, fq.inters...),
		predicates: append([]predicate.Forum{}, fq.predicates...),
		withGroup:  fq.withGroup.Clone(),
		withPosts:  fq.withPosts.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *ForumQuery) WithGroup(opts ...func(*GroupQuery)) *ForumQuery {
	query := (&GroupClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withGroup = query
	return fq
}

// WithPosts tells the query-builder to eager-load the nodes that are connected to
// the "posts" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *ForumQuery) WithPosts(opts ...func(*ForumPostQuery)) *ForumQuery {
	query := (&ForumPostClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withPosts = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Forum.Query().
//		GroupBy(forum.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *ForumQuery) GroupBy(field string, fields ...string) *ForumGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ForumGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = forum.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Forum.Query().
//		Select(forum.FieldName).
//		Scan(ctx, &v)
func (fq *ForumQuery) Select(fields ...string) *ForumSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &ForumSelect{ForumQuery: fq}
	sbuild.label = forum.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ForumSelect configured with the given aggregations.
func (fq *ForumQuery) Aggregate(fns ...AggregateFunc) *ForumSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *ForumQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !forum.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *ForumQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Forum, error) {
	var (
		nodes       = []*Forum{}
		withFKs     = fq.withFKs
		_spec       = fq.querySpec()
		loadedTypes = [2]bool{
			fq.withGroup != nil,
			fq.withPosts != nil,
		}
	)
	if fq.withGroup != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, forum.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Forum).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Forum{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withGroup; query != nil {
		if err := fq.loadGroup(ctx, query, nodes, nil,
			func(n *Forum, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	if query := fq.withPosts; query != nil {
		if err := fq.loadPosts(ctx, query, nodes,
			func(n *Forum) { n.Edges.Posts = []*ForumPost{} },
			func(n *Forum, e *ForumPost) { n.Edges.Posts = append(n.Edges.Posts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *ForumQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*Forum, init func(*Forum), assign func(*Forum, *Group)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Forum)
	for i := range nodes {
		if nodes[i].group_forums == nil {
			continue
		}
		fk := *nodes[i].group_forums
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(entgroup.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_forums" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (fq *ForumQuery) loadPosts(ctx context.Context, query *ForumPostQuery, nodes []*Forum, init func(*Forum), assign func(*Forum, *ForumPost)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Forum)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.ForumPost(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(forum.PostsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.forum_posts
		if fk == nil {
			return fmt.Errorf(`foreign-key "forum_posts" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "forum_posts" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fq *ForumQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *ForumQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(forum.Table, forum.Columns, sqlgraph.NewFieldSpec(forum.FieldID, field.TypeInt))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, forum.FieldID)
		for i := range fields {
			if fields[i] != forum.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *ForumQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(forum.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = forum.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForumGroupBy is the group-by builder for Forum entities.
type ForumGroupBy struct {
	selector
	build *ForumQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *ForumGroupBy) Aggregate(fns ...AggregateFunc) *ForumGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *ForumGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, "GroupBy")
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ForumQuery, *ForumGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *ForumGroupBy) sqlScan(ctx context.Context, root *ForumQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ForumSelect is the builder for selecting fields of Forum entities.
type ForumSelect struct {
	*ForumQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *ForumSelect) Aggregate(fns ...AggregateFunc) *ForumSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *ForumSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, "Select")
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ForumQuery, *ForumSelect](ctx, fs.ForumQuery, fs, fs.inters, v)
}

func (fs *ForumSelect) sqlScan(ctx context.Context, root *ForumQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}