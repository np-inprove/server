// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/forum"
	"github.com/np-inprove/server/internal/ent/forumpost"
	entgroup "github.com/np-inprove/server/internal/ent/group"
)

// ForumCreate is the builder for creating a Forum entity.
type ForumCreate struct {
	config
	mutation *ForumMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (fc *ForumCreate) SetName(s string) *ForumCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetShortName sets the "short_name" field.
func (fc *ForumCreate) SetShortName(s string) *ForumCreate {
	fc.mutation.SetShortName(s)
	return fc
}

// SetDescription sets the "description" field.
func (fc *ForumCreate) SetDescription(s string) *ForumCreate {
	fc.mutation.SetDescription(s)
	return fc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (fc *ForumCreate) SetGroupID(id int) *ForumCreate {
	fc.mutation.SetGroupID(id)
	return fc
}

// SetGroup sets the "group" edge to the Group entity.
func (fc *ForumCreate) SetGroup(g *Group) *ForumCreate {
	return fc.SetGroupID(g.ID)
}

// AddPostIDs adds the "posts" edge to the ForumPost entity by IDs.
func (fc *ForumCreate) AddPostIDs(ids ...int) *ForumCreate {
	fc.mutation.AddPostIDs(ids...)
	return fc
}

// AddPosts adds the "posts" edges to the ForumPost entity.
func (fc *ForumCreate) AddPosts(f ...*ForumPost) *ForumCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddPostIDs(ids...)
}

// Mutation returns the ForumMutation object of the builder.
func (fc *ForumCreate) Mutation() *ForumMutation {
	return fc.mutation
}

// Save creates the Forum in the database.
func (fc *ForumCreate) Save(ctx context.Context) (*Forum, error) {
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *ForumCreate) SaveX(ctx context.Context) *Forum {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *ForumCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *ForumCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *ForumCreate) check() error {
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Forum.name"`)}
	}
	if v, ok := fc.mutation.Name(); ok {
		if err := forum.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Forum.name": %w`, err)}
		}
	}
	if _, ok := fc.mutation.ShortName(); !ok {
		return &ValidationError{Name: "short_name", err: errors.New(`ent: missing required field "Forum.short_name"`)}
	}
	if v, ok := fc.mutation.ShortName(); ok {
		if err := forum.ShortNameValidator(v); err != nil {
			return &ValidationError{Name: "short_name", err: fmt.Errorf(`ent: validator failed for field "Forum.short_name": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Forum.description"`)}
	}
	if _, ok := fc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "Forum.group"`)}
	}
	return nil
}

func (fc *ForumCreate) sqlSave(ctx context.Context) (*Forum, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *ForumCreate) createSpec() (*Forum, *sqlgraph.CreateSpec) {
	var (
		_node = &Forum{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(forum.Table, sqlgraph.NewFieldSpec(forum.FieldID, field.TypeInt))
	)
	_spec.OnConflict = fc.conflict
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(forum.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.ShortName(); ok {
		_spec.SetField(forum.FieldShortName, field.TypeString, value)
		_node.ShortName = value
	}
	if value, ok := fc.mutation.Description(); ok {
		_spec.SetField(forum.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := fc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forum.GroupTable,
			Columns: []string{forum.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(entgroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_forums = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.PostsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forum.PostsTable,
			Columns: []string{forum.PostsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
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
//	client.Forum.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ForumUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (fc *ForumCreate) OnConflict(opts ...sql.ConflictOption) *ForumUpsertOne {
	fc.conflict = opts
	return &ForumUpsertOne{
		create: fc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Forum.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fc *ForumCreate) OnConflictColumns(columns ...string) *ForumUpsertOne {
	fc.conflict = append(fc.conflict, sql.ConflictColumns(columns...))
	return &ForumUpsertOne{
		create: fc,
	}
}

type (
	// ForumUpsertOne is the builder for "upsert"-ing
	//  one Forum node.
	ForumUpsertOne struct {
		create *ForumCreate
	}

	// ForumUpsert is the "OnConflict" setter.
	ForumUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *ForumUpsert) SetName(v string) *ForumUpsert {
	u.Set(forum.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ForumUpsert) UpdateName() *ForumUpsert {
	u.SetExcluded(forum.FieldName)
	return u
}

// SetShortName sets the "short_name" field.
func (u *ForumUpsert) SetShortName(v string) *ForumUpsert {
	u.Set(forum.FieldShortName, v)
	return u
}

// UpdateShortName sets the "short_name" field to the value that was provided on create.
func (u *ForumUpsert) UpdateShortName() *ForumUpsert {
	u.SetExcluded(forum.FieldShortName)
	return u
}

// SetDescription sets the "description" field.
func (u *ForumUpsert) SetDescription(v string) *ForumUpsert {
	u.Set(forum.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ForumUpsert) UpdateDescription() *ForumUpsert {
	u.SetExcluded(forum.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Forum.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ForumUpsertOne) UpdateNewValues() *ForumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Forum.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ForumUpsertOne) Ignore() *ForumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ForumUpsertOne) DoNothing() *ForumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ForumCreate.OnConflict
// documentation for more info.
func (u *ForumUpsertOne) Update(set func(*ForumUpsert)) *ForumUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ForumUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ForumUpsertOne) SetName(v string) *ForumUpsertOne {
	return u.Update(func(s *ForumUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ForumUpsertOne) UpdateName() *ForumUpsertOne {
	return u.Update(func(s *ForumUpsert) {
		s.UpdateName()
	})
}

// SetShortName sets the "short_name" field.
func (u *ForumUpsertOne) SetShortName(v string) *ForumUpsertOne {
	return u.Update(func(s *ForumUpsert) {
		s.SetShortName(v)
	})
}

// UpdateShortName sets the "short_name" field to the value that was provided on create.
func (u *ForumUpsertOne) UpdateShortName() *ForumUpsertOne {
	return u.Update(func(s *ForumUpsert) {
		s.UpdateShortName()
	})
}

// SetDescription sets the "description" field.
func (u *ForumUpsertOne) SetDescription(v string) *ForumUpsertOne {
	return u.Update(func(s *ForumUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ForumUpsertOne) UpdateDescription() *ForumUpsertOne {
	return u.Update(func(s *ForumUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *ForumUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ForumCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ForumUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ForumUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ForumUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ForumCreateBulk is the builder for creating many Forum entities in bulk.
type ForumCreateBulk struct {
	config
	builders []*ForumCreate
	conflict []sql.ConflictOption
}

// Save creates the Forum entities in the database.
func (fcb *ForumCreateBulk) Save(ctx context.Context) ([]*Forum, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Forum, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ForumMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = fcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *ForumCreateBulk) SaveX(ctx context.Context) []*Forum {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *ForumCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *ForumCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Forum.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ForumUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (fcb *ForumCreateBulk) OnConflict(opts ...sql.ConflictOption) *ForumUpsertBulk {
	fcb.conflict = opts
	return &ForumUpsertBulk{
		create: fcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Forum.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (fcb *ForumCreateBulk) OnConflictColumns(columns ...string) *ForumUpsertBulk {
	fcb.conflict = append(fcb.conflict, sql.ConflictColumns(columns...))
	return &ForumUpsertBulk{
		create: fcb,
	}
}

// ForumUpsertBulk is the builder for "upsert"-ing
// a bulk of Forum nodes.
type ForumUpsertBulk struct {
	create *ForumCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Forum.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ForumUpsertBulk) UpdateNewValues() *ForumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Forum.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ForumUpsertBulk) Ignore() *ForumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ForumUpsertBulk) DoNothing() *ForumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ForumCreateBulk.OnConflict
// documentation for more info.
func (u *ForumUpsertBulk) Update(set func(*ForumUpsert)) *ForumUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ForumUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *ForumUpsertBulk) SetName(v string) *ForumUpsertBulk {
	return u.Update(func(s *ForumUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ForumUpsertBulk) UpdateName() *ForumUpsertBulk {
	return u.Update(func(s *ForumUpsert) {
		s.UpdateName()
	})
}

// SetShortName sets the "short_name" field.
func (u *ForumUpsertBulk) SetShortName(v string) *ForumUpsertBulk {
	return u.Update(func(s *ForumUpsert) {
		s.SetShortName(v)
	})
}

// UpdateShortName sets the "short_name" field to the value that was provided on create.
func (u *ForumUpsertBulk) UpdateShortName() *ForumUpsertBulk {
	return u.Update(func(s *ForumUpsert) {
		s.UpdateShortName()
	})
}

// SetDescription sets the "description" field.
func (u *ForumUpsertBulk) SetDescription(v string) *ForumUpsertBulk {
	return u.Update(func(s *ForumUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ForumUpsertBulk) UpdateDescription() *ForumUpsertBulk {
	return u.Update(func(s *ForumUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *ForumUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ForumCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ForumCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ForumUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}