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
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/redemption"
	"github.com/np-inprove/server/internal/ent/voucher"
)

// VoucherQuery is the builder for querying Voucher entities.
type VoucherQuery struct {
	config
	ctx             *QueryContext
	order           []voucher.OrderOption
	inters          []Interceptor
	predicates      []predicate.Voucher
	withRedemptions *RedemptionQuery
	withInstitution *InstitutionQuery
	withFKs         bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VoucherQuery builder.
func (vq *VoucherQuery) Where(ps ...predicate.Voucher) *VoucherQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VoucherQuery) Limit(limit int) *VoucherQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VoucherQuery) Offset(offset int) *VoucherQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VoucherQuery) Unique(unique bool) *VoucherQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VoucherQuery) Order(o ...voucher.OrderOption) *VoucherQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryRedemptions chains the current query on the "redemptions" edge.
func (vq *VoucherQuery) QueryRedemptions() *RedemptionQuery {
	query := (&RedemptionClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voucher.Table, voucher.FieldID, selector),
			sqlgraph.To(redemption.Table, redemption.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, voucher.RedemptionsTable, voucher.RedemptionsColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryInstitution chains the current query on the "institution" edge.
func (vq *VoucherQuery) QueryInstitution() *InstitutionQuery {
	query := (&InstitutionClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(voucher.Table, voucher.FieldID, selector),
			sqlgraph.To(institution.Table, institution.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, voucher.InstitutionTable, voucher.InstitutionColumn),
		)
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Voucher entity from the query.
// Returns a *NotFoundError when no Voucher was found.
func (vq *VoucherQuery) First(ctx context.Context) (*Voucher, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{voucher.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VoucherQuery) FirstX(ctx context.Context) *Voucher {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Voucher ID from the query.
// Returns a *NotFoundError when no Voucher ID was found.
func (vq *VoucherQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{voucher.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VoucherQuery) FirstIDX(ctx context.Context) int {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Voucher entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Voucher entity is found.
// Returns a *NotFoundError when no Voucher entities are found.
func (vq *VoucherQuery) Only(ctx context.Context) (*Voucher, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{voucher.Label}
	default:
		return nil, &NotSingularError{voucher.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VoucherQuery) OnlyX(ctx context.Context) *Voucher {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Voucher ID in the query.
// Returns a *NotSingularError when more than one Voucher ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VoucherQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{voucher.Label}
	default:
		err = &NotSingularError{voucher.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VoucherQuery) OnlyIDX(ctx context.Context) int {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Vouchers.
func (vq *VoucherQuery) All(ctx context.Context) ([]*Voucher, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Voucher, *VoucherQuery]()
	return withInterceptors[[]*Voucher](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VoucherQuery) AllX(ctx context.Context) []*Voucher {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Voucher IDs.
func (vq *VoucherQuery) IDs(ctx context.Context) (ids []int, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err = vq.Select(voucher.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VoucherQuery) IDsX(ctx context.Context) []int {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VoucherQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VoucherQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VoucherQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VoucherQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VoucherQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VoucherQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VoucherQuery) Clone() *VoucherQuery {
	if vq == nil {
		return nil
	}
	return &VoucherQuery{
		config:          vq.config,
		ctx:             vq.ctx.Clone(),
		order:           append([]voucher.OrderOption{}, vq.order...),
		inters:          append([]Interceptor{}, vq.inters...),
		predicates:      append([]predicate.Voucher{}, vq.predicates...),
		withRedemptions: vq.withRedemptions.Clone(),
		withInstitution: vq.withInstitution.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithRedemptions tells the query-builder to eager-load the nodes that are connected to
// the "redemptions" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoucherQuery) WithRedemptions(opts ...func(*RedemptionQuery)) *VoucherQuery {
	query := (&RedemptionClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withRedemptions = query
	return vq
}

// WithInstitution tells the query-builder to eager-load the nodes that are connected to
// the "institution" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VoucherQuery) WithInstitution(opts ...func(*InstitutionQuery)) *VoucherQuery {
	query := (&InstitutionClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withInstitution = query
	return vq
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
//	client.Voucher.Query().
//		GroupBy(voucher.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (vq *VoucherQuery) GroupBy(field string, fields ...string) *VoucherGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VoucherGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = voucher.Label
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
//	client.Voucher.Query().
//		Select(voucher.FieldName).
//		Scan(ctx, &v)
func (vq *VoucherQuery) Select(fields ...string) *VoucherSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VoucherSelect{VoucherQuery: vq}
	sbuild.label = voucher.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VoucherSelect configured with the given aggregations.
func (vq *VoucherQuery) Aggregate(fns ...AggregateFunc) *VoucherSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VoucherQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !voucher.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VoucherQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Voucher, error) {
	var (
		nodes       = []*Voucher{}
		withFKs     = vq.withFKs
		_spec       = vq.querySpec()
		loadedTypes = [2]bool{
			vq.withRedemptions != nil,
			vq.withInstitution != nil,
		}
	)
	if vq.withInstitution != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, voucher.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Voucher).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Voucher{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withRedemptions; query != nil {
		if err := vq.loadRedemptions(ctx, query, nodes,
			func(n *Voucher) { n.Edges.Redemptions = []*Redemption{} },
			func(n *Voucher, e *Redemption) { n.Edges.Redemptions = append(n.Edges.Redemptions, e) }); err != nil {
			return nil, err
		}
	}
	if query := vq.withInstitution; query != nil {
		if err := vq.loadInstitution(ctx, query, nodes, nil,
			func(n *Voucher, e *Institution) { n.Edges.Institution = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VoucherQuery) loadRedemptions(ctx context.Context, query *RedemptionQuery, nodes []*Voucher, init func(*Voucher), assign func(*Voucher, *Redemption)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Voucher)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Redemption(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(voucher.RedemptionsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.voucher_redemptions
		if fk == nil {
			return fmt.Errorf(`foreign-key "voucher_redemptions" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "voucher_redemptions" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (vq *VoucherQuery) loadInstitution(ctx context.Context, query *InstitutionQuery, nodes []*Voucher, init func(*Voucher), assign func(*Voucher, *Institution)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Voucher)
	for i := range nodes {
		if nodes[i].institution_vouchers == nil {
			continue
		}
		fk := *nodes[i].institution_vouchers
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(institution.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "institution_vouchers" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (vq *VoucherQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VoucherQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(voucher.Table, voucher.Columns, sqlgraph.NewFieldSpec(voucher.FieldID, field.TypeInt))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, voucher.FieldID)
		for i := range fields {
			if fields[i] != voucher.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VoucherQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(voucher.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = voucher.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// VoucherGroupBy is the group-by builder for Voucher entities.
type VoucherGroupBy struct {
	selector
	build *VoucherQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VoucherGroupBy) Aggregate(fns ...AggregateFunc) *VoucherGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VoucherGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoucherQuery, *VoucherGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VoucherGroupBy) sqlScan(ctx context.Context, root *VoucherQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VoucherSelect is the builder for selecting fields of Voucher entities.
type VoucherSelect struct {
	*VoucherQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VoucherSelect) Aggregate(fns ...AggregateFunc) *VoucherSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VoucherSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VoucherQuery, *VoucherSelect](ctx, vs.VoucherQuery, vs, vs.inters, v)
}

func (vs *VoucherSelect) sqlScan(ctx context.Context, root *VoucherQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
