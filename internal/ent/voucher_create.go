// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	entinstitution "github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/redemption"
	"github.com/np-inprove/server/internal/ent/voucher"
)

// VoucherCreate is the builder for creating a Voucher entity.
type VoucherCreate struct {
	config
	mutation *VoucherMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (vc *VoucherCreate) SetName(s string) *VoucherCreate {
	vc.mutation.SetName(s)
	return vc
}

// SetDescription sets the "description" field.
func (vc *VoucherCreate) SetDescription(s string) *VoucherCreate {
	vc.mutation.SetDescription(s)
	return vc
}

// SetPointsRequired sets the "points_required" field.
func (vc *VoucherCreate) SetPointsRequired(i int) *VoucherCreate {
	vc.mutation.SetPointsRequired(i)
	return vc
}

// AddRedemptionIDs adds the "redemptions" edge to the Redemption entity by IDs.
func (vc *VoucherCreate) AddRedemptionIDs(ids ...int) *VoucherCreate {
	vc.mutation.AddRedemptionIDs(ids...)
	return vc
}

// AddRedemptions adds the "redemptions" edges to the Redemption entity.
func (vc *VoucherCreate) AddRedemptions(r ...*Redemption) *VoucherCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return vc.AddRedemptionIDs(ids...)
}

// SetInstitutionID sets the "institution" edge to the Institution entity by ID.
func (vc *VoucherCreate) SetInstitutionID(id int) *VoucherCreate {
	vc.mutation.SetInstitutionID(id)
	return vc
}

// SetInstitution sets the "institution" edge to the Institution entity.
func (vc *VoucherCreate) SetInstitution(i *Institution) *VoucherCreate {
	return vc.SetInstitutionID(i.ID)
}

// Mutation returns the VoucherMutation object of the builder.
func (vc *VoucherCreate) Mutation() *VoucherMutation {
	return vc.mutation
}

// Save creates the Voucher in the database.
func (vc *VoucherCreate) Save(ctx context.Context) (*Voucher, error) {
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VoucherCreate) SaveX(ctx context.Context) *Voucher {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VoucherCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VoucherCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VoucherCreate) check() error {
	if _, ok := vc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Voucher.name"`)}
	}
	if v, ok := vc.mutation.Name(); ok {
		if err := voucher.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Voucher.name": %w`, err)}
		}
	}
	if _, ok := vc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Voucher.description"`)}
	}
	if _, ok := vc.mutation.PointsRequired(); !ok {
		return &ValidationError{Name: "points_required", err: errors.New(`ent: missing required field "Voucher.points_required"`)}
	}
	if v, ok := vc.mutation.PointsRequired(); ok {
		if err := voucher.PointsRequiredValidator(v); err != nil {
			return &ValidationError{Name: "points_required", err: fmt.Errorf(`ent: validator failed for field "Voucher.points_required": %w`, err)}
		}
	}
	if _, ok := vc.mutation.InstitutionID(); !ok {
		return &ValidationError{Name: "institution", err: errors.New(`ent: missing required edge "Voucher.institution"`)}
	}
	return nil
}

func (vc *VoucherCreate) sqlSave(ctx context.Context) (*Voucher, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VoucherCreate) createSpec() (*Voucher, *sqlgraph.CreateSpec) {
	var (
		_node = &Voucher{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(voucher.Table, sqlgraph.NewFieldSpec(voucher.FieldID, field.TypeInt))
	)
	_spec.OnConflict = vc.conflict
	if value, ok := vc.mutation.Name(); ok {
		_spec.SetField(voucher.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := vc.mutation.Description(); ok {
		_spec.SetField(voucher.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := vc.mutation.PointsRequired(); ok {
		_spec.SetField(voucher.FieldPointsRequired, field.TypeInt, value)
		_node.PointsRequired = value
	}
	if nodes := vc.mutation.RedemptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   voucher.RedemptionsTable,
			Columns: []string{voucher.RedemptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redemption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vc.mutation.InstitutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   voucher.InstitutionTable,
			Columns: []string{voucher.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entinstitution.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.institution_vouchers = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Voucher.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VoucherUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (vc *VoucherCreate) OnConflict(opts ...sql.ConflictOption) *VoucherUpsertOne {
	vc.conflict = opts
	return &VoucherUpsertOne{
		create: vc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Voucher.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vc *VoucherCreate) OnConflictColumns(columns ...string) *VoucherUpsertOne {
	vc.conflict = append(vc.conflict, sql.ConflictColumns(columns...))
	return &VoucherUpsertOne{
		create: vc,
	}
}

type (
	// VoucherUpsertOne is the builder for "upsert"-ing
	//  one Voucher node.
	VoucherUpsertOne struct {
		create *VoucherCreate
	}

	// VoucherUpsert is the "OnConflict" setter.
	VoucherUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *VoucherUpsert) SetName(v string) *VoucherUpsert {
	u.Set(voucher.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *VoucherUpsert) UpdateName() *VoucherUpsert {
	u.SetExcluded(voucher.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *VoucherUpsert) SetDescription(v string) *VoucherUpsert {
	u.Set(voucher.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *VoucherUpsert) UpdateDescription() *VoucherUpsert {
	u.SetExcluded(voucher.FieldDescription)
	return u
}

// SetPointsRequired sets the "points_required" field.
func (u *VoucherUpsert) SetPointsRequired(v int) *VoucherUpsert {
	u.Set(voucher.FieldPointsRequired, v)
	return u
}

// UpdatePointsRequired sets the "points_required" field to the value that was provided on create.
func (u *VoucherUpsert) UpdatePointsRequired() *VoucherUpsert {
	u.SetExcluded(voucher.FieldPointsRequired)
	return u
}

// AddPointsRequired adds v to the "points_required" field.
func (u *VoucherUpsert) AddPointsRequired(v int) *VoucherUpsert {
	u.Add(voucher.FieldPointsRequired, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Voucher.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *VoucherUpsertOne) UpdateNewValues() *VoucherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Voucher.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *VoucherUpsertOne) Ignore() *VoucherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VoucherUpsertOne) DoNothing() *VoucherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VoucherCreate.OnConflict
// documentation for more info.
func (u *VoucherUpsertOne) Update(set func(*VoucherUpsert)) *VoucherUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VoucherUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *VoucherUpsertOne) SetName(v string) *VoucherUpsertOne {
	return u.Update(func(s *VoucherUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *VoucherUpsertOne) UpdateName() *VoucherUpsertOne {
	return u.Update(func(s *VoucherUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *VoucherUpsertOne) SetDescription(v string) *VoucherUpsertOne {
	return u.Update(func(s *VoucherUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *VoucherUpsertOne) UpdateDescription() *VoucherUpsertOne {
	return u.Update(func(s *VoucherUpsert) {
		s.UpdateDescription()
	})
}

// SetPointsRequired sets the "points_required" field.
func (u *VoucherUpsertOne) SetPointsRequired(v int) *VoucherUpsertOne {
	return u.Update(func(s *VoucherUpsert) {
		s.SetPointsRequired(v)
	})
}

// AddPointsRequired adds v to the "points_required" field.
func (u *VoucherUpsertOne) AddPointsRequired(v int) *VoucherUpsertOne {
	return u.Update(func(s *VoucherUpsert) {
		s.AddPointsRequired(v)
	})
}

// UpdatePointsRequired sets the "points_required" field to the value that was provided on create.
func (u *VoucherUpsertOne) UpdatePointsRequired() *VoucherUpsertOne {
	return u.Update(func(s *VoucherUpsert) {
		s.UpdatePointsRequired()
	})
}

// Exec executes the query.
func (u *VoucherUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for VoucherCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VoucherUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *VoucherUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *VoucherUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// VoucherCreateBulk is the builder for creating many Voucher entities in bulk.
type VoucherCreateBulk struct {
	config
	builders []*VoucherCreate
	conflict []sql.ConflictOption
}

// Save creates the Voucher entities in the database.
func (vcb *VoucherCreateBulk) Save(ctx context.Context) ([]*Voucher, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Voucher, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VoucherMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = vcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VoucherCreateBulk) SaveX(ctx context.Context) []*Voucher {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VoucherCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VoucherCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Voucher.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.VoucherUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (vcb *VoucherCreateBulk) OnConflict(opts ...sql.ConflictOption) *VoucherUpsertBulk {
	vcb.conflict = opts
	return &VoucherUpsertBulk{
		create: vcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Voucher.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (vcb *VoucherCreateBulk) OnConflictColumns(columns ...string) *VoucherUpsertBulk {
	vcb.conflict = append(vcb.conflict, sql.ConflictColumns(columns...))
	return &VoucherUpsertBulk{
		create: vcb,
	}
}

// VoucherUpsertBulk is the builder for "upsert"-ing
// a bulk of Voucher nodes.
type VoucherUpsertBulk struct {
	create *VoucherCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Voucher.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *VoucherUpsertBulk) UpdateNewValues() *VoucherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Voucher.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *VoucherUpsertBulk) Ignore() *VoucherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *VoucherUpsertBulk) DoNothing() *VoucherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the VoucherCreateBulk.OnConflict
// documentation for more info.
func (u *VoucherUpsertBulk) Update(set func(*VoucherUpsert)) *VoucherUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&VoucherUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *VoucherUpsertBulk) SetName(v string) *VoucherUpsertBulk {
	return u.Update(func(s *VoucherUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *VoucherUpsertBulk) UpdateName() *VoucherUpsertBulk {
	return u.Update(func(s *VoucherUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *VoucherUpsertBulk) SetDescription(v string) *VoucherUpsertBulk {
	return u.Update(func(s *VoucherUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *VoucherUpsertBulk) UpdateDescription() *VoucherUpsertBulk {
	return u.Update(func(s *VoucherUpsert) {
		s.UpdateDescription()
	})
}

// SetPointsRequired sets the "points_required" field.
func (u *VoucherUpsertBulk) SetPointsRequired(v int) *VoucherUpsertBulk {
	return u.Update(func(s *VoucherUpsert) {
		s.SetPointsRequired(v)
	})
}

// AddPointsRequired adds v to the "points_required" field.
func (u *VoucherUpsertBulk) AddPointsRequired(v int) *VoucherUpsertBulk {
	return u.Update(func(s *VoucherUpsert) {
		s.AddPointsRequired(v)
	})
}

// UpdatePointsRequired sets the "points_required" field to the value that was provided on create.
func (u *VoucherUpsertBulk) UpdatePointsRequired() *VoucherUpsertBulk {
	return u.Update(func(s *VoucherUpsert) {
		s.UpdatePointsRequired()
	})
}

// Exec executes the query.
func (u *VoucherUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the VoucherCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for VoucherCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *VoucherUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
