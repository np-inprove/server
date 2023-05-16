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
	"github.com/np-inprove/server/internal/ent/milestone"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/studyplan"
	"github.com/np-inprove/server/internal/ent/user"
)

// StudyPlanQuery is the builder for querying StudyPlan entities.
type StudyPlanQuery struct {
	config
	ctx            *QueryContext
	order          []studyplan.OrderOption
	inters         []Interceptor
	predicates     []predicate.StudyPlan
	withAuthor     *UserQuery
	withMilestones *MilestoneQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the StudyPlanQuery builder.
func (spq *StudyPlanQuery) Where(ps ...predicate.StudyPlan) *StudyPlanQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit the number of records to be returned by this query.
func (spq *StudyPlanQuery) Limit(limit int) *StudyPlanQuery {
	spq.ctx.Limit = &limit
	return spq
}

// Offset to start from.
func (spq *StudyPlanQuery) Offset(offset int) *StudyPlanQuery {
	spq.ctx.Offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *StudyPlanQuery) Unique(unique bool) *StudyPlanQuery {
	spq.ctx.Unique = &unique
	return spq
}

// Order specifies how the records should be ordered.
func (spq *StudyPlanQuery) Order(o ...studyplan.OrderOption) *StudyPlanQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// QueryAuthor chains the current query on the "author" edge.
func (spq *StudyPlanQuery) QueryAuthor() *UserQuery {
	query := (&UserClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(studyplan.Table, studyplan.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, studyplan.AuthorTable, studyplan.AuthorColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMilestones chains the current query on the "milestones" edge.
func (spq *StudyPlanQuery) QueryMilestones() *MilestoneQuery {
	query := (&MilestoneClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := spq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(studyplan.Table, studyplan.FieldID, selector),
			sqlgraph.To(milestone.Table, milestone.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, studyplan.MilestonesTable, studyplan.MilestonesColumn),
		)
		fromU = sqlgraph.SetNeighbors(spq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first StudyPlan entity from the query.
// Returns a *NotFoundError when no StudyPlan was found.
func (spq *StudyPlanQuery) First(ctx context.Context) (*StudyPlan, error) {
	nodes, err := spq.Limit(1).All(setContextOp(ctx, spq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{studyplan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *StudyPlanQuery) FirstX(ctx context.Context) *StudyPlan {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first StudyPlan ID from the query.
// Returns a *NotFoundError when no StudyPlan ID was found.
func (spq *StudyPlanQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = spq.Limit(1).IDs(setContextOp(ctx, spq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{studyplan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *StudyPlanQuery) FirstIDX(ctx context.Context) int {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single StudyPlan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one StudyPlan entity is found.
// Returns a *NotFoundError when no StudyPlan entities are found.
func (spq *StudyPlanQuery) Only(ctx context.Context) (*StudyPlan, error) {
	nodes, err := spq.Limit(2).All(setContextOp(ctx, spq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{studyplan.Label}
	default:
		return nil, &NotSingularError{studyplan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *StudyPlanQuery) OnlyX(ctx context.Context) *StudyPlan {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only StudyPlan ID in the query.
// Returns a *NotSingularError when more than one StudyPlan ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *StudyPlanQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = spq.Limit(2).IDs(setContextOp(ctx, spq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{studyplan.Label}
	default:
		err = &NotSingularError{studyplan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *StudyPlanQuery) OnlyIDX(ctx context.Context) int {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of StudyPlans.
func (spq *StudyPlanQuery) All(ctx context.Context) ([]*StudyPlan, error) {
	ctx = setContextOp(ctx, spq.ctx, "All")
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*StudyPlan, *StudyPlanQuery]()
	return withInterceptors[[]*StudyPlan](ctx, spq, qr, spq.inters)
}

// AllX is like All, but panics if an error occurs.
func (spq *StudyPlanQuery) AllX(ctx context.Context) []*StudyPlan {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of StudyPlan IDs.
func (spq *StudyPlanQuery) IDs(ctx context.Context) (ids []int, err error) {
	if spq.ctx.Unique == nil && spq.path != nil {
		spq.Unique(true)
	}
	ctx = setContextOp(ctx, spq.ctx, "IDs")
	if err = spq.Select(studyplan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *StudyPlanQuery) IDsX(ctx context.Context) []int {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *StudyPlanQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, spq.ctx, "Count")
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, spq, querierCount[*StudyPlanQuery](), spq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (spq *StudyPlanQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *StudyPlanQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, spq.ctx, "Exist")
	switch _, err := spq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *StudyPlanQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the StudyPlanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *StudyPlanQuery) Clone() *StudyPlanQuery {
	if spq == nil {
		return nil
	}
	return &StudyPlanQuery{
		config:         spq.config,
		ctx:            spq.ctx.Clone(),
		order:          append([]studyplan.OrderOption{}, spq.order...),
		inters:         append([]Interceptor{}, spq.inters...),
		predicates:     append([]predicate.StudyPlan{}, spq.predicates...),
		withAuthor:     spq.withAuthor.Clone(),
		withMilestones: spq.withMilestones.Clone(),
		// clone intermediate query.
		sql:  spq.sql.Clone(),
		path: spq.path,
	}
}

// WithAuthor tells the query-builder to eager-load the nodes that are connected to
// the "author" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StudyPlanQuery) WithAuthor(opts ...func(*UserQuery)) *StudyPlanQuery {
	query := (&UserClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withAuthor = query
	return spq
}

// WithMilestones tells the query-builder to eager-load the nodes that are connected to
// the "milestones" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *StudyPlanQuery) WithMilestones(opts ...func(*MilestoneQuery)) *StudyPlanQuery {
	query := (&MilestoneClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withMilestones = query
	return spq
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
//	client.StudyPlan.Query().
//		GroupBy(studyplan.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (spq *StudyPlanQuery) GroupBy(field string, fields ...string) *StudyPlanGroupBy {
	spq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &StudyPlanGroupBy{build: spq}
	grbuild.flds = &spq.ctx.Fields
	grbuild.label = studyplan.Label
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
//	client.StudyPlan.Query().
//		Select(studyplan.FieldName).
//		Scan(ctx, &v)
func (spq *StudyPlanQuery) Select(fields ...string) *StudyPlanSelect {
	spq.ctx.Fields = append(spq.ctx.Fields, fields...)
	sbuild := &StudyPlanSelect{StudyPlanQuery: spq}
	sbuild.label = studyplan.Label
	sbuild.flds, sbuild.scan = &spq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a StudyPlanSelect configured with the given aggregations.
func (spq *StudyPlanQuery) Aggregate(fns ...AggregateFunc) *StudyPlanSelect {
	return spq.Select().Aggregate(fns...)
}

func (spq *StudyPlanQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range spq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, spq); err != nil {
				return err
			}
		}
	}
	for _, f := range spq.ctx.Fields {
		if !studyplan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.sql = prev
	}
	return nil
}

func (spq *StudyPlanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*StudyPlan, error) {
	var (
		nodes       = []*StudyPlan{}
		withFKs     = spq.withFKs
		_spec       = spq.querySpec()
		loadedTypes = [2]bool{
			spq.withAuthor != nil,
			spq.withMilestones != nil,
		}
	)
	if spq.withAuthor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, studyplan.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*StudyPlan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &StudyPlan{config: spq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, spq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := spq.withAuthor; query != nil {
		if err := spq.loadAuthor(ctx, query, nodes, nil,
			func(n *StudyPlan, e *User) { n.Edges.Author = e }); err != nil {
			return nil, err
		}
	}
	if query := spq.withMilestones; query != nil {
		if err := spq.loadMilestones(ctx, query, nodes,
			func(n *StudyPlan) { n.Edges.Milestones = []*Milestone{} },
			func(n *StudyPlan, e *Milestone) { n.Edges.Milestones = append(n.Edges.Milestones, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (spq *StudyPlanQuery) loadAuthor(ctx context.Context, query *UserQuery, nodes []*StudyPlan, init func(*StudyPlan), assign func(*StudyPlan, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*StudyPlan)
	for i := range nodes {
		if nodes[i].study_plan_author == nil {
			continue
		}
		fk := *nodes[i].study_plan_author
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
			return fmt.Errorf(`unexpected foreign-key "study_plan_author" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (spq *StudyPlanQuery) loadMilestones(ctx context.Context, query *MilestoneQuery, nodes []*StudyPlan, init func(*StudyPlan), assign func(*StudyPlan, *Milestone)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*StudyPlan)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Milestone(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(studyplan.MilestonesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.study_plan_milestones
		if fk == nil {
			return fmt.Errorf(`foreign-key "study_plan_milestones" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "study_plan_milestones" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (spq *StudyPlanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := spq.querySpec()
	_spec.Node.Columns = spq.ctx.Fields
	if len(spq.ctx.Fields) > 0 {
		_spec.Unique = spq.ctx.Unique != nil && *spq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, spq.driver, _spec)
}

func (spq *StudyPlanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(studyplan.Table, studyplan.Columns, sqlgraph.NewFieldSpec(studyplan.FieldID, field.TypeInt))
	_spec.From = spq.sql
	if unique := spq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if spq.path != nil {
		_spec.Unique = true
	}
	if fields := spq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, studyplan.FieldID)
		for i := range fields {
			if fields[i] != studyplan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := spq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := spq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := spq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := spq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (spq *StudyPlanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(spq.driver.Dialect())
	t1 := builder.Table(studyplan.Table)
	columns := spq.ctx.Fields
	if len(columns) == 0 {
		columns = studyplan.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if spq.sql != nil {
		selector = spq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if spq.ctx.Unique != nil && *spq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range spq.predicates {
		p(selector)
	}
	for _, p := range spq.order {
		p(selector)
	}
	if offset := spq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := spq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// StudyPlanGroupBy is the group-by builder for StudyPlan entities.
type StudyPlanGroupBy struct {
	selector
	build *StudyPlanQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *StudyPlanGroupBy) Aggregate(fns ...AggregateFunc) *StudyPlanGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the selector query and scans the result into the given value.
func (spgb *StudyPlanGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, spgb.build.ctx, "GroupBy")
	if err := spgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StudyPlanQuery, *StudyPlanGroupBy](ctx, spgb.build, spgb, spgb.build.inters, v)
}

func (spgb *StudyPlanGroupBy) sqlScan(ctx context.Context, root *StudyPlanQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(spgb.fns))
	for _, fn := range spgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*spgb.flds)+len(spgb.fns))
		for _, f := range *spgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*spgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := spgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// StudyPlanSelect is the builder for selecting fields of StudyPlan entities.
type StudyPlanSelect struct {
	*StudyPlanQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sps *StudyPlanSelect) Aggregate(fns ...AggregateFunc) *StudyPlanSelect {
	sps.fns = append(sps.fns, fns...)
	return sps
}

// Scan applies the selector query and scans the result into the given value.
func (sps *StudyPlanSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sps.ctx, "Select")
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*StudyPlanQuery, *StudyPlanSelect](ctx, sps.StudyPlanQuery, sps, sps.inters, v)
}

func (sps *StudyPlanSelect) sqlScan(ctx context.Context, root *StudyPlanQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sps.fns))
	for _, fn := range sps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}