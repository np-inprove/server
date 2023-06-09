// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/accessory"
	entinstitution "github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/redemption"
)

// AccessoryCreate is the builder for creating a Accessory entity.
type AccessoryCreate struct {
	config
	mutation *AccessoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (ac *AccessoryCreate) SetName(s string) *AccessoryCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetDescription sets the "description" field.
func (ac *AccessoryCreate) SetDescription(s string) *AccessoryCreate {
	ac.mutation.SetDescription(s)
	return ac
}

// SetPointsRequired sets the "points_required" field.
func (ac *AccessoryCreate) SetPointsRequired(i int) *AccessoryCreate {
	ac.mutation.SetPointsRequired(i)
	return ac
}

// AddRedemptionIDs adds the "redemptions" edge to the Redemption entity by IDs.
func (ac *AccessoryCreate) AddRedemptionIDs(ids ...int) *AccessoryCreate {
	ac.mutation.AddRedemptionIDs(ids...)
	return ac
}

// AddRedemptions adds the "redemptions" edges to the Redemption entity.
func (ac *AccessoryCreate) AddRedemptions(r ...*Redemption) *AccessoryCreate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ac.AddRedemptionIDs(ids...)
}

// SetInstitutionID sets the "institution" edge to the Institution entity by ID.
func (ac *AccessoryCreate) SetInstitutionID(id int) *AccessoryCreate {
	ac.mutation.SetInstitutionID(id)
	return ac
}

// SetInstitution sets the "institution" edge to the Institution entity.
func (ac *AccessoryCreate) SetInstitution(i *Institution) *AccessoryCreate {
	return ac.SetInstitutionID(i.ID)
}

// Mutation returns the AccessoryMutation object of the builder.
func (ac *AccessoryCreate) Mutation() *AccessoryMutation {
	return ac.mutation
}

// Save creates the Accessory in the database.
func (ac *AccessoryCreate) Save(ctx context.Context) (*Accessory, error) {
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccessoryCreate) SaveX(ctx context.Context) *Accessory {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccessoryCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccessoryCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccessoryCreate) check() error {
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Accessory.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := accessory.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Accessory.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Accessory.description"`)}
	}
	if _, ok := ac.mutation.PointsRequired(); !ok {
		return &ValidationError{Name: "points_required", err: errors.New(`ent: missing required field "Accessory.points_required"`)}
	}
	if v, ok := ac.mutation.PointsRequired(); ok {
		if err := accessory.PointsRequiredValidator(v); err != nil {
			return &ValidationError{Name: "points_required", err: fmt.Errorf(`ent: validator failed for field "Accessory.points_required": %w`, err)}
		}
	}
	if _, ok := ac.mutation.InstitutionID(); !ok {
		return &ValidationError{Name: "institution", err: errors.New(`ent: missing required edge "Accessory.institution"`)}
	}
	return nil
}

func (ac *AccessoryCreate) sqlSave(ctx context.Context) (*Accessory, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccessoryCreate) createSpec() (*Accessory, *sqlgraph.CreateSpec) {
	var (
		_node = &Accessory{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(accessory.Table, sqlgraph.NewFieldSpec(accessory.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ac.conflict
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(accessory.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Description(); ok {
		_spec.SetField(accessory.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ac.mutation.PointsRequired(); ok {
		_spec.SetField(accessory.FieldPointsRequired, field.TypeInt, value)
		_node.PointsRequired = value
	}
	if nodes := ac.mutation.RedemptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   accessory.RedemptionsTable,
			Columns: []string{accessory.RedemptionsColumn},
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
	if nodes := ac.mutation.InstitutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   accessory.InstitutionTable,
			Columns: []string{accessory.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entinstitution.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.institution_accessories = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Accessory.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccessoryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (ac *AccessoryCreate) OnConflict(opts ...sql.ConflictOption) *AccessoryUpsertOne {
	ac.conflict = opts
	return &AccessoryUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Accessory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AccessoryCreate) OnConflictColumns(columns ...string) *AccessoryUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AccessoryUpsertOne{
		create: ac,
	}
}

type (
	// AccessoryUpsertOne is the builder for "upsert"-ing
	//  one Accessory node.
	AccessoryUpsertOne struct {
		create *AccessoryCreate
	}

	// AccessoryUpsert is the "OnConflict" setter.
	AccessoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *AccessoryUpsert) SetName(v string) *AccessoryUpsert {
	u.Set(accessory.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccessoryUpsert) UpdateName() *AccessoryUpsert {
	u.SetExcluded(accessory.FieldName)
	return u
}

// SetDescription sets the "description" field.
func (u *AccessoryUpsert) SetDescription(v string) *AccessoryUpsert {
	u.Set(accessory.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AccessoryUpsert) UpdateDescription() *AccessoryUpsert {
	u.SetExcluded(accessory.FieldDescription)
	return u
}

// SetPointsRequired sets the "points_required" field.
func (u *AccessoryUpsert) SetPointsRequired(v int) *AccessoryUpsert {
	u.Set(accessory.FieldPointsRequired, v)
	return u
}

// UpdatePointsRequired sets the "points_required" field to the value that was provided on create.
func (u *AccessoryUpsert) UpdatePointsRequired() *AccessoryUpsert {
	u.SetExcluded(accessory.FieldPointsRequired)
	return u
}

// AddPointsRequired adds v to the "points_required" field.
func (u *AccessoryUpsert) AddPointsRequired(v int) *AccessoryUpsert {
	u.Add(accessory.FieldPointsRequired, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Accessory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AccessoryUpsertOne) UpdateNewValues() *AccessoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Accessory.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AccessoryUpsertOne) Ignore() *AccessoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccessoryUpsertOne) DoNothing() *AccessoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccessoryCreate.OnConflict
// documentation for more info.
func (u *AccessoryUpsertOne) Update(set func(*AccessoryUpsert)) *AccessoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccessoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AccessoryUpsertOne) SetName(v string) *AccessoryUpsertOne {
	return u.Update(func(s *AccessoryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccessoryUpsertOne) UpdateName() *AccessoryUpsertOne {
	return u.Update(func(s *AccessoryUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *AccessoryUpsertOne) SetDescription(v string) *AccessoryUpsertOne {
	return u.Update(func(s *AccessoryUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AccessoryUpsertOne) UpdateDescription() *AccessoryUpsertOne {
	return u.Update(func(s *AccessoryUpsert) {
		s.UpdateDescription()
	})
}

// SetPointsRequired sets the "points_required" field.
func (u *AccessoryUpsertOne) SetPointsRequired(v int) *AccessoryUpsertOne {
	return u.Update(func(s *AccessoryUpsert) {
		s.SetPointsRequired(v)
	})
}

// AddPointsRequired adds v to the "points_required" field.
func (u *AccessoryUpsertOne) AddPointsRequired(v int) *AccessoryUpsertOne {
	return u.Update(func(s *AccessoryUpsert) {
		s.AddPointsRequired(v)
	})
}

// UpdatePointsRequired sets the "points_required" field to the value that was provided on create.
func (u *AccessoryUpsertOne) UpdatePointsRequired() *AccessoryUpsertOne {
	return u.Update(func(s *AccessoryUpsert) {
		s.UpdatePointsRequired()
	})
}

// Exec executes the query.
func (u *AccessoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccessoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccessoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AccessoryUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AccessoryUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AccessoryCreateBulk is the builder for creating many Accessory entities in bulk.
type AccessoryCreateBulk struct {
	config
	builders []*AccessoryCreate
	conflict []sql.ConflictOption
}

// Save creates the Accessory entities in the database.
func (acb *AccessoryCreateBulk) Save(ctx context.Context) ([]*Accessory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Accessory, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccessoryMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccessoryCreateBulk) SaveX(ctx context.Context) []*Accessory {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccessoryCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccessoryCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Accessory.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccessoryUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (acb *AccessoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *AccessoryUpsertBulk {
	acb.conflict = opts
	return &AccessoryUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Accessory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AccessoryCreateBulk) OnConflictColumns(columns ...string) *AccessoryUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AccessoryUpsertBulk{
		create: acb,
	}
}

// AccessoryUpsertBulk is the builder for "upsert"-ing
// a bulk of Accessory nodes.
type AccessoryUpsertBulk struct {
	create *AccessoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Accessory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *AccessoryUpsertBulk) UpdateNewValues() *AccessoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Accessory.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AccessoryUpsertBulk) Ignore() *AccessoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccessoryUpsertBulk) DoNothing() *AccessoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccessoryCreateBulk.OnConflict
// documentation for more info.
func (u *AccessoryUpsertBulk) Update(set func(*AccessoryUpsert)) *AccessoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccessoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *AccessoryUpsertBulk) SetName(v string) *AccessoryUpsertBulk {
	return u.Update(func(s *AccessoryUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *AccessoryUpsertBulk) UpdateName() *AccessoryUpsertBulk {
	return u.Update(func(s *AccessoryUpsert) {
		s.UpdateName()
	})
}

// SetDescription sets the "description" field.
func (u *AccessoryUpsertBulk) SetDescription(v string) *AccessoryUpsertBulk {
	return u.Update(func(s *AccessoryUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *AccessoryUpsertBulk) UpdateDescription() *AccessoryUpsertBulk {
	return u.Update(func(s *AccessoryUpsert) {
		s.UpdateDescription()
	})
}

// SetPointsRequired sets the "points_required" field.
func (u *AccessoryUpsertBulk) SetPointsRequired(v int) *AccessoryUpsertBulk {
	return u.Update(func(s *AccessoryUpsert) {
		s.SetPointsRequired(v)
	})
}

// AddPointsRequired adds v to the "points_required" field.
func (u *AccessoryUpsertBulk) AddPointsRequired(v int) *AccessoryUpsertBulk {
	return u.Update(func(s *AccessoryUpsert) {
		s.AddPointsRequired(v)
	})
}

// UpdatePointsRequired sets the "points_required" field to the value that was provided on create.
func (u *AccessoryUpsertBulk) UpdatePointsRequired() *AccessoryUpsertBulk {
	return u.Update(func(s *AccessoryUpsert) {
		s.UpdatePointsRequired()
	})
}

// Exec executes the query.
func (u *AccessoryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AccessoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccessoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccessoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
