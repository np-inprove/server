// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/milestone"
	"github.com/np-inprove/server/internal/ent/studyplan"
	"github.com/np-inprove/server/internal/ent/user"
)

// StudyPlanCreate is the builder for creating a StudyPlan entity.
type StudyPlanCreate struct {
	config
	mutation *StudyPlanMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (spc *StudyPlanCreate) SetName(s string) *StudyPlanCreate {
	spc.mutation.SetName(s)
	return spc
}

// SetShareCode sets the "share_code" field.
func (spc *StudyPlanCreate) SetShareCode(s string) *StudyPlanCreate {
	spc.mutation.SetShareCode(s)
	return spc
}

// SetNillableShareCode sets the "share_code" field if the given value is not nil.
func (spc *StudyPlanCreate) SetNillableShareCode(s *string) *StudyPlanCreate {
	if s != nil {
		spc.SetShareCode(*s)
	}
	return spc
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (spc *StudyPlanCreate) SetAuthorID(id int) *StudyPlanCreate {
	spc.mutation.SetAuthorID(id)
	return spc
}

// SetAuthor sets the "author" edge to the User entity.
func (spc *StudyPlanCreate) SetAuthor(u *User) *StudyPlanCreate {
	return spc.SetAuthorID(u.ID)
}

// AddMilestoneIDs adds the "milestones" edge to the Milestone entity by IDs.
func (spc *StudyPlanCreate) AddMilestoneIDs(ids ...int) *StudyPlanCreate {
	spc.mutation.AddMilestoneIDs(ids...)
	return spc
}

// AddMilestones adds the "milestones" edges to the Milestone entity.
func (spc *StudyPlanCreate) AddMilestones(m ...*Milestone) *StudyPlanCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return spc.AddMilestoneIDs(ids...)
}

// Mutation returns the StudyPlanMutation object of the builder.
func (spc *StudyPlanCreate) Mutation() *StudyPlanMutation {
	return spc.mutation
}

// Save creates the StudyPlan in the database.
func (spc *StudyPlanCreate) Save(ctx context.Context) (*StudyPlan, error) {
	return withHooks(ctx, spc.sqlSave, spc.mutation, spc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (spc *StudyPlanCreate) SaveX(ctx context.Context) *StudyPlan {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *StudyPlanCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *StudyPlanCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *StudyPlanCreate) check() error {
	if _, ok := spc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "StudyPlan.name"`)}
	}
	if v, ok := spc.mutation.Name(); ok {
		if err := studyplan.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "StudyPlan.name": %w`, err)}
		}
	}
	if _, ok := spc.mutation.AuthorID(); !ok {
		return &ValidationError{Name: "author", err: errors.New(`ent: missing required edge "StudyPlan.author"`)}
	}
	return nil
}

func (spc *StudyPlanCreate) sqlSave(ctx context.Context) (*StudyPlan, error) {
	if err := spc.check(); err != nil {
		return nil, err
	}
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	spc.mutation.id = &_node.ID
	spc.mutation.done = true
	return _node, nil
}

func (spc *StudyPlanCreate) createSpec() (*StudyPlan, *sqlgraph.CreateSpec) {
	var (
		_node = &StudyPlan{config: spc.config}
		_spec = sqlgraph.NewCreateSpec(studyplan.Table, sqlgraph.NewFieldSpec(studyplan.FieldID, field.TypeInt))
	)
	_spec.OnConflict = spc.conflict
	if value, ok := spc.mutation.Name(); ok {
		_spec.SetField(studyplan.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := spc.mutation.ShareCode(); ok {
		_spec.SetField(studyplan.FieldShareCode, field.TypeString, value)
		_node.ShareCode = value
	}
	if nodes := spc.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   studyplan.AuthorTable,
			Columns: []string{studyplan.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.study_plan_author = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := spc.mutation.MilestonesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   studyplan.MilestonesTable,
			Columns: []string{studyplan.MilestonesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(milestone.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.StudyPlan.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StudyPlanUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (spc *StudyPlanCreate) OnConflict(opts ...sql.ConflictOption) *StudyPlanUpsertOne {
	spc.conflict = opts
	return &StudyPlanUpsertOne{
		create: spc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.StudyPlan.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (spc *StudyPlanCreate) OnConflictColumns(columns ...string) *StudyPlanUpsertOne {
	spc.conflict = append(spc.conflict, sql.ConflictColumns(columns...))
	return &StudyPlanUpsertOne{
		create: spc,
	}
}

type (
	// StudyPlanUpsertOne is the builder for "upsert"-ing
	//  one StudyPlan node.
	StudyPlanUpsertOne struct {
		create *StudyPlanCreate
	}

	// StudyPlanUpsert is the "OnConflict" setter.
	StudyPlanUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *StudyPlanUpsert) SetName(v string) *StudyPlanUpsert {
	u.Set(studyplan.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StudyPlanUpsert) UpdateName() *StudyPlanUpsert {
	u.SetExcluded(studyplan.FieldName)
	return u
}

// SetShareCode sets the "share_code" field.
func (u *StudyPlanUpsert) SetShareCode(v string) *StudyPlanUpsert {
	u.Set(studyplan.FieldShareCode, v)
	return u
}

// UpdateShareCode sets the "share_code" field to the value that was provided on create.
func (u *StudyPlanUpsert) UpdateShareCode() *StudyPlanUpsert {
	u.SetExcluded(studyplan.FieldShareCode)
	return u
}

// ClearShareCode clears the value of the "share_code" field.
func (u *StudyPlanUpsert) ClearShareCode() *StudyPlanUpsert {
	u.SetNull(studyplan.FieldShareCode)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.StudyPlan.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *StudyPlanUpsertOne) UpdateNewValues() *StudyPlanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.StudyPlan.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *StudyPlanUpsertOne) Ignore() *StudyPlanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StudyPlanUpsertOne) DoNothing() *StudyPlanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StudyPlanCreate.OnConflict
// documentation for more info.
func (u *StudyPlanUpsertOne) Update(set func(*StudyPlanUpsert)) *StudyPlanUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StudyPlanUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StudyPlanUpsertOne) SetName(v string) *StudyPlanUpsertOne {
	return u.Update(func(s *StudyPlanUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StudyPlanUpsertOne) UpdateName() *StudyPlanUpsertOne {
	return u.Update(func(s *StudyPlanUpsert) {
		s.UpdateName()
	})
}

// SetShareCode sets the "share_code" field.
func (u *StudyPlanUpsertOne) SetShareCode(v string) *StudyPlanUpsertOne {
	return u.Update(func(s *StudyPlanUpsert) {
		s.SetShareCode(v)
	})
}

// UpdateShareCode sets the "share_code" field to the value that was provided on create.
func (u *StudyPlanUpsertOne) UpdateShareCode() *StudyPlanUpsertOne {
	return u.Update(func(s *StudyPlanUpsert) {
		s.UpdateShareCode()
	})
}

// ClearShareCode clears the value of the "share_code" field.
func (u *StudyPlanUpsertOne) ClearShareCode() *StudyPlanUpsertOne {
	return u.Update(func(s *StudyPlanUpsert) {
		s.ClearShareCode()
	})
}

// Exec executes the query.
func (u *StudyPlanUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StudyPlanCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StudyPlanUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StudyPlanUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StudyPlanUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StudyPlanCreateBulk is the builder for creating many StudyPlan entities in bulk.
type StudyPlanCreateBulk struct {
	config
	builders []*StudyPlanCreate
	conflict []sql.ConflictOption
}

// Save creates the StudyPlan entities in the database.
func (spcb *StudyPlanCreateBulk) Save(ctx context.Context) ([]*StudyPlan, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*StudyPlan, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StudyPlanMutation)
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
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = spcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *StudyPlanCreateBulk) SaveX(ctx context.Context) []*StudyPlan {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *StudyPlanCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *StudyPlanCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.StudyPlan.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StudyPlanUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (spcb *StudyPlanCreateBulk) OnConflict(opts ...sql.ConflictOption) *StudyPlanUpsertBulk {
	spcb.conflict = opts
	return &StudyPlanUpsertBulk{
		create: spcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.StudyPlan.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (spcb *StudyPlanCreateBulk) OnConflictColumns(columns ...string) *StudyPlanUpsertBulk {
	spcb.conflict = append(spcb.conflict, sql.ConflictColumns(columns...))
	return &StudyPlanUpsertBulk{
		create: spcb,
	}
}

// StudyPlanUpsertBulk is the builder for "upsert"-ing
// a bulk of StudyPlan nodes.
type StudyPlanUpsertBulk struct {
	create *StudyPlanCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.StudyPlan.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *StudyPlanUpsertBulk) UpdateNewValues() *StudyPlanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.StudyPlan.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *StudyPlanUpsertBulk) Ignore() *StudyPlanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StudyPlanUpsertBulk) DoNothing() *StudyPlanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StudyPlanCreateBulk.OnConflict
// documentation for more info.
func (u *StudyPlanUpsertBulk) Update(set func(*StudyPlanUpsert)) *StudyPlanUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StudyPlanUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *StudyPlanUpsertBulk) SetName(v string) *StudyPlanUpsertBulk {
	return u.Update(func(s *StudyPlanUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *StudyPlanUpsertBulk) UpdateName() *StudyPlanUpsertBulk {
	return u.Update(func(s *StudyPlanUpsert) {
		s.UpdateName()
	})
}

// SetShareCode sets the "share_code" field.
func (u *StudyPlanUpsertBulk) SetShareCode(v string) *StudyPlanUpsertBulk {
	return u.Update(func(s *StudyPlanUpsert) {
		s.SetShareCode(v)
	})
}

// UpdateShareCode sets the "share_code" field to the value that was provided on create.
func (u *StudyPlanUpsertBulk) UpdateShareCode() *StudyPlanUpsertBulk {
	return u.Update(func(s *StudyPlanUpsert) {
		s.UpdateShareCode()
	})
}

// ClearShareCode clears the value of the "share_code" field.
func (u *StudyPlanUpsertBulk) ClearShareCode() *StudyPlanUpsertBulk {
	return u.Update(func(s *StudyPlanUpsert) {
		s.ClearShareCode()
	})
}

// Exec executes the query.
func (u *StudyPlanUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StudyPlanCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StudyPlanCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StudyPlanUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
