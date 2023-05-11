// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/groupuser"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/user"
)

// GroupUserQuery is the builder for querying GroupUser entities.
type GroupUserQuery struct {
	config
	ctx        *QueryContext
	order      []groupuser.OrderOption
	inters     []Interceptor
	predicates []predicate.GroupUser
	withGroup  *GroupQuery
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GroupUserQuery builder.
func (guq *GroupUserQuery) Where(ps ...predicate.GroupUser) *GroupUserQuery {
	guq.predicates = append(guq.predicates, ps...)
	return guq
}

// Limit the number of records to be returned by this query.
func (guq *GroupUserQuery) Limit(limit int) *GroupUserQuery {
	guq.ctx.Limit = &limit
	return guq
}

// Offset to start from.
func (guq *GroupUserQuery) Offset(offset int) *GroupUserQuery {
	guq.ctx.Offset = &offset
	return guq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (guq *GroupUserQuery) Unique(unique bool) *GroupUserQuery {
	guq.ctx.Unique = &unique
	return guq
}

// Order specifies how the records should be ordered.
func (guq *GroupUserQuery) Order(o ...groupuser.OrderOption) *GroupUserQuery {
	guq.order = append(guq.order, o...)
	return guq
}

// QueryGroup chains the current query on the "group" edge.
func (guq *GroupUserQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: guq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := guq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := guq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupuser.Table, groupuser.GroupColumn, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, groupuser.GroupTable, groupuser.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(guq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (guq *GroupUserQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: guq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := guq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := guq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(groupuser.Table, groupuser.UserColumn, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, groupuser.UserTable, groupuser.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(guq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first GroupUser entity from the query.
// Returns a *NotFoundError when no GroupUser was found.
func (guq *GroupUserQuery) First(ctx context.Context) (*GroupUser, error) {
	nodes, err := guq.Limit(1).All(setContextOp(ctx, guq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{groupuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (guq *GroupUserQuery) FirstX(ctx context.Context) *GroupUser {
	node, err := guq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// Only returns a single GroupUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GroupUser entity is found.
// Returns a *NotFoundError when no GroupUser entities are found.
func (guq *GroupUserQuery) Only(ctx context.Context) (*GroupUser, error) {
	nodes, err := guq.Limit(2).All(setContextOp(ctx, guq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{groupuser.Label}
	default:
		return nil, &NotSingularError{groupuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (guq *GroupUserQuery) OnlyX(ctx context.Context) *GroupUser {
	node, err := guq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// All executes the query and returns a list of GroupUsers.
func (guq *GroupUserQuery) All(ctx context.Context) ([]*GroupUser, error) {
	ctx = setContextOp(ctx, guq.ctx, "All")
	if err := guq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*GroupUser, *GroupUserQuery]()
	return withInterceptors[[]*GroupUser](ctx, guq, qr, guq.inters)
}

// AllX is like All, but panics if an error occurs.
func (guq *GroupUserQuery) AllX(ctx context.Context) []*GroupUser {
	nodes, err := guq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// Count returns the count of the given query.
func (guq *GroupUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, guq.ctx, "Count")
	if err := guq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, guq, querierCount[*GroupUserQuery](), guq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (guq *GroupUserQuery) CountX(ctx context.Context) int {
	count, err := guq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (guq *GroupUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, guq.ctx, "Exist")
	switch _, err := guq.First(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (guq *GroupUserQuery) ExistX(ctx context.Context) bool {
	exist, err := guq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GroupUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (guq *GroupUserQuery) Clone() *GroupUserQuery {
	if guq == nil {
		return nil
	}
	return &GroupUserQuery{
		config:     guq.config,
		ctx:        guq.ctx.Clone(),
		order:      append([]groupuser.OrderOption{}, guq.order...),
		inters:     append([]Interceptor{}, guq.inters...),
		predicates: append([]predicate.GroupUser{}, guq.predicates...),
		withGroup:  guq.withGroup.Clone(),
		withUser:   guq.withUser.Clone(),
		// clone intermediate query.
		sql:  guq.sql.Clone(),
		path: guq.path,
	}
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (guq *GroupUserQuery) WithGroup(opts ...func(*GroupQuery)) *GroupUserQuery {
	query := (&GroupClient{config: guq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	guq.withGroup = query
	return guq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (guq *GroupUserQuery) WithUser(opts ...func(*UserQuery)) *GroupUserQuery {
	query := (&UserClient{config: guq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	guq.withUser = query
	return guq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GroupID int `json:"group_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupUser.Query().
//		GroupBy(groupuser.FieldGroupID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (guq *GroupUserQuery) GroupBy(field string, fields ...string) *GroupUserGroupBy {
	guq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GroupUserGroupBy{build: guq}
	grbuild.flds = &guq.ctx.Fields
	grbuild.label = groupuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GroupID int `json:"group_id,omitempty"`
//	}
//
//	client.GroupUser.Query().
//		Select(groupuser.FieldGroupID).
//		Scan(ctx, &v)
func (guq *GroupUserQuery) Select(fields ...string) *GroupUserSelect {
	guq.ctx.Fields = append(guq.ctx.Fields, fields...)
	sbuild := &GroupUserSelect{GroupUserQuery: guq}
	sbuild.label = groupuser.Label
	sbuild.flds, sbuild.scan = &guq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GroupUserSelect configured with the given aggregations.
func (guq *GroupUserQuery) Aggregate(fns ...AggregateFunc) *GroupUserSelect {
	return guq.Select().Aggregate(fns...)
}

func (guq *GroupUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range guq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, guq); err != nil {
				return err
			}
		}
	}
	for _, f := range guq.ctx.Fields {
		if !groupuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if guq.path != nil {
		prev, err := guq.path(ctx)
		if err != nil {
			return err
		}
		guq.sql = prev
	}
	return nil
}

func (guq *GroupUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*GroupUser, error) {
	var (
		nodes       = []*GroupUser{}
		_spec       = guq.querySpec()
		loadedTypes = [2]bool{
			guq.withGroup != nil,
			guq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*GroupUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &GroupUser{config: guq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, guq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := guq.withGroup; query != nil {
		if err := guq.loadGroup(ctx, query, nodes, nil,
			func(n *GroupUser, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	if query := guq.withUser; query != nil {
		if err := guq.loadUser(ctx, query, nodes, nil,
			func(n *GroupUser, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (guq *GroupUserQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*GroupUser, init func(*GroupUser), assign func(*GroupUser, *Group)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GroupUser)
	for i := range nodes {
		fk := nodes[i].GroupID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (guq *GroupUserQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*GroupUser, init func(*GroupUser), assign func(*GroupUser, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*GroupUser)
	for i := range nodes {
		fk := nodes[i].UserID
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
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (guq *GroupUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := guq.querySpec()
	_spec.Unique = false
	_spec.Node.Columns = nil
	return sqlgraph.CountNodes(ctx, guq.driver, _spec)
}

func (guq *GroupUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(groupuser.Table, groupuser.Columns, nil)
	_spec.From = guq.sql
	if unique := guq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if guq.path != nil {
		_spec.Unique = true
	}
	if fields := guq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		for i := range fields {
			_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
		}
		if guq.withGroup != nil {
			_spec.Node.AddColumnOnce(groupuser.FieldGroupID)
		}
		if guq.withUser != nil {
			_spec.Node.AddColumnOnce(groupuser.FieldUserID)
		}
	}
	if ps := guq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := guq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := guq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := guq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (guq *GroupUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(guq.driver.Dialect())
	t1 := builder.Table(groupuser.Table)
	columns := guq.ctx.Fields
	if len(columns) == 0 {
		columns = groupuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if guq.sql != nil {
		selector = guq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if guq.ctx.Unique != nil && *guq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range guq.predicates {
		p(selector)
	}
	for _, p := range guq.order {
		p(selector)
	}
	if offset := guq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := guq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GroupUserGroupBy is the group-by builder for GroupUser entities.
type GroupUserGroupBy struct {
	selector
	build *GroupUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gugb *GroupUserGroupBy) Aggregate(fns ...AggregateFunc) *GroupUserGroupBy {
	gugb.fns = append(gugb.fns, fns...)
	return gugb
}

// Scan applies the selector query and scans the result into the given value.
func (gugb *GroupUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gugb.build.ctx, "GroupBy")
	if err := gugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupUserQuery, *GroupUserGroupBy](ctx, gugb.build, gugb, gugb.build.inters, v)
}

func (gugb *GroupUserGroupBy) sqlScan(ctx context.Context, root *GroupUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(gugb.fns))
	for _, fn := range gugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*gugb.flds)+len(gugb.fns))
		for _, f := range *gugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*gugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GroupUserSelect is the builder for selecting fields of GroupUser entities.
type GroupUserSelect struct {
	*GroupUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gus *GroupUserSelect) Aggregate(fns ...AggregateFunc) *GroupUserSelect {
	gus.fns = append(gus.fns, fns...)
	return gus
}

// Scan applies the selector query and scans the result into the given value.
func (gus *GroupUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gus.ctx, "Select")
	if err := gus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GroupUserQuery, *GroupUserSelect](ctx, gus.GroupUserQuery, gus, gus.inters, v)
}

func (gus *GroupUserSelect) sqlScan(ctx context.Context, root *GroupUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gus.fns))
	for _, fn := range gus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
