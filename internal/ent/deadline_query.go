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
	"github.com/np-inprove/server/internal/ent/deadline"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/user"
)

// DeadlineQuery is the builder for querying Deadline entities.
type DeadlineQuery struct {
	config
	ctx            *QueryContext
	order          []deadline.OrderOption
	inters         []Interceptor
	predicates     []predicate.Deadline
	withAuthor     *UserQuery
	withVotedUsers *UserQuery
	withGroup      *GroupQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeadlineQuery builder.
func (dq *DeadlineQuery) Where(ps ...predicate.Deadline) *DeadlineQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DeadlineQuery) Limit(limit int) *DeadlineQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DeadlineQuery) Offset(offset int) *DeadlineQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DeadlineQuery) Unique(unique bool) *DeadlineQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DeadlineQuery) Order(o ...deadline.OrderOption) *DeadlineQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryAuthor chains the current query on the "author" edge.
func (dq *DeadlineQuery) QueryAuthor() *UserQuery {
	query := (&UserClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deadline.Table, deadline.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, deadline.AuthorTable, deadline.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryVotedUsers chains the current query on the "voted_users" edge.
func (dq *DeadlineQuery) QueryVotedUsers() *UserQuery {
	query := (&UserClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deadline.Table, deadline.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, deadline.VotedUsersTable, deadline.VotedUsersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryGroup chains the current query on the "group" edge.
func (dq *DeadlineQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(deadline.Table, deadline.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, deadline.GroupTable, deadline.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Deadline entity from the query.
// Returns a *NotFoundError when no Deadline was found.
func (dq *DeadlineQuery) First(ctx context.Context) (*Deadline, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{deadline.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DeadlineQuery) FirstX(ctx context.Context) *Deadline {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Deadline ID from the query.
// Returns a *NotFoundError when no Deadline ID was found.
func (dq *DeadlineQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{deadline.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DeadlineQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Deadline entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Deadline entity is found.
// Returns a *NotFoundError when no Deadline entities are found.
func (dq *DeadlineQuery) Only(ctx context.Context) (*Deadline, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{deadline.Label}
	default:
		return nil, &NotSingularError{deadline.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DeadlineQuery) OnlyX(ctx context.Context) *Deadline {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Deadline ID in the query.
// Returns a *NotSingularError when more than one Deadline ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DeadlineQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{deadline.Label}
	default:
		err = &NotSingularError{deadline.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DeadlineQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Deadlines.
func (dq *DeadlineQuery) All(ctx context.Context) ([]*Deadline, error) {
	ctx = setContextOp(ctx, dq.ctx, "All")
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Deadline, *DeadlineQuery]()
	return withInterceptors[[]*Deadline](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DeadlineQuery) AllX(ctx context.Context) []*Deadline {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Deadline IDs.
func (dq *DeadlineQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, "IDs")
	if err = dq.Select(deadline.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DeadlineQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DeadlineQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, "Count")
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DeadlineQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DeadlineQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DeadlineQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, "Exist")
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DeadlineQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeadlineQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DeadlineQuery) Clone() *DeadlineQuery {
	if dq == nil {
		return nil
	}
	return &DeadlineQuery{
		config:         dq.config,
		ctx:            dq.ctx.Clone(),
		order:          append([]deadline.OrderOption{}, dq.order...),
		inters:         append([]Interceptor{}, dq.inters...),
		predicates:     append([]predicate.Deadline{}, dq.predicates...),
		withAuthor:     dq.withAuthor.Clone(),
		withVotedUsers: dq.withVotedUsers.Clone(),
		withGroup:      dq.withGroup.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeadlineQuery) WithAuthor(opts ...func(*UserQuery)) *DeadlineQuery {
	query := (&UserClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withAuthor = query
	return dq
}

// WithVotedUsers tells the query-builder to eager-load the nodes that are connected to
// the "voted_users" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeadlineQuery) WithVotedUsers(opts ...func(*UserQuery)) *DeadlineQuery {
	query := (&UserClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withVotedUsers = query
	return dq
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeadlineQuery) WithGroup(opts ...func(*GroupQuery)) *DeadlineQuery {
	query := (&GroupClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withGroup = query
	return dq
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
//	client.Deadline.Query().
//		GroupBy(deadline.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DeadlineQuery) GroupBy(field string, fields ...string) *DeadlineGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DeadlineGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = deadline.Label
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
//	client.Deadline.Query().
//		Select(deadline.FieldName).
//		Scan(ctx, &v)
func (dq *DeadlineQuery) Select(fields ...string) *DeadlineSelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DeadlineSelect{DeadlineQuery: dq}
	sbuild.label = deadline.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DeadlineSelect configured with the given aggregations.
func (dq *DeadlineQuery) Aggregate(fns ...AggregateFunc) *DeadlineSelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DeadlineQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !deadline.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DeadlineQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Deadline, error) {
	var (
		nodes       = []*Deadline{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [3]bool{
			dq.withAuthor != nil,
			dq.withVotedUsers != nil,
			dq.withGroup != nil,
		}
	)
	if dq.withAuthor != nil || dq.withGroup != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, deadline.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Deadline).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Deadline{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withAuthor; query != nil {
		if err := dq.loadAuthor(ctx, query, nodes, nil,
			func(n *Deadline, e *User) { n.Edges.Author = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withVotedUsers; query != nil {
		if err := dq.loadVotedUsers(ctx, query, nodes,
			func(n *Deadline) { n.Edges.VotedUsers = []*User{} },
			func(n *Deadline, e *User) { n.Edges.VotedUsers = append(n.Edges.VotedUsers, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withGroup; query != nil {
		if err := dq.loadGroup(ctx, query, nodes, nil,
			func(n *Deadline, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DeadlineQuery) loadAuthor(ctx context.Context, query *UserQuery, nodes []*Deadline, init func(*Deadline), assign func(*Deadline, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Deadline)
	for i := range nodes {
		if nodes[i].deadline_author == nil {
			continue
		}
		fk := *nodes[i].deadline_author
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
			return fmt.Errorf(`unexpected foreign-key "deadline_author" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DeadlineQuery) loadVotedUsers(ctx context.Context, query *UserQuery, nodes []*Deadline, init func(*Deadline), assign func(*Deadline, *User)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Deadline)
	nids := make(map[int]map[*Deadline]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(deadline.VotedUsersTable)
		s.Join(joinT).On(s.C(user.FieldID), joinT.C(deadline.VotedUsersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(deadline.VotedUsersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(deadline.VotedUsersPrimaryKey[0]))
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
					nids[inValue] = map[*Deadline]struct{}{byID[outValue]: {}}
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
			return fmt.Errorf(`unexpected "voted_users" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (dq *DeadlineQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*Deadline, init func(*Deadline), assign func(*Deadline, *Group)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Deadline)
	for i := range nodes {
		if nodes[i].group_deadlines == nil {
			continue
		}
		fk := *nodes[i].group_deadlines
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
			return fmt.Errorf(`unexpected foreign-key "group_deadlines" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (dq *DeadlineQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DeadlineQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(deadline.Table, deadline.Columns, sqlgraph.NewFieldSpec(deadline.FieldID, field.TypeInt))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deadline.FieldID)
		for i := range fields {
			if fields[i] != deadline.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DeadlineQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(deadline.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = deadline.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DeadlineGroupBy is the group-by builder for Deadline entities.
type DeadlineGroupBy struct {
	selector
	build *DeadlineQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DeadlineGroupBy) Aggregate(fns ...AggregateFunc) *DeadlineGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DeadlineGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, "GroupBy")
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeadlineQuery, *DeadlineGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DeadlineGroupBy) sqlScan(ctx context.Context, root *DeadlineQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DeadlineSelect is the builder for selecting fields of Deadline entities.
type DeadlineSelect struct {
	*DeadlineQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DeadlineSelect) Aggregate(fns ...AggregateFunc) *DeadlineSelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DeadlineSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, "Select")
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeadlineQuery, *DeadlineSelect](ctx, ds.DeadlineQuery, ds, ds.inters, v)
}

func (ds *DeadlineSelect) sqlScan(ctx context.Context, root *DeadlineQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
