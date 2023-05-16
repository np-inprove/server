// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/groupuser"
	"github.com/np-inprove/server/internal/ent/user"
)

// GroupUserCreate is the builder for creating a GroupUser entity.
type GroupUserCreate struct {
	config
	mutation *GroupUserMutation
	hooks    []Hook
}

// SetGroupID sets the "group_id" field.
func (guc *GroupUserCreate) SetGroupID(i int) *GroupUserCreate {
	guc.mutation.SetGroupID(i)
	return guc
}

// SetUserID sets the "user_id" field.
func (guc *GroupUserCreate) SetUserID(i int) *GroupUserCreate {
	guc.mutation.SetUserID(i)
	return guc
}

// SetRole sets the "role" field.
func (guc *GroupUserCreate) SetRole(gr groupuser.Role) *GroupUserCreate {
	guc.mutation.SetRole(gr)
	return guc
}

// SetGroup sets the "group" edge to the Group entity.
func (guc *GroupUserCreate) SetGroup(g *Group) *GroupUserCreate {
	return guc.SetGroupID(g.ID)
}

// SetUser sets the "user" edge to the User entity.
func (guc *GroupUserCreate) SetUser(u *User) *GroupUserCreate {
	return guc.SetUserID(u.ID)
}

// Mutation returns the GroupUserMutation object of the builder.
func (guc *GroupUserCreate) Mutation() *GroupUserMutation {
	return guc.mutation
}

// Save creates the GroupUser in the database.
func (guc *GroupUserCreate) Save(ctx context.Context) (*GroupUser, error) {
	return withHooks(ctx, guc.sqlSave, guc.mutation, guc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (guc *GroupUserCreate) SaveX(ctx context.Context) *GroupUser {
	v, err := guc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (guc *GroupUserCreate) Exec(ctx context.Context) error {
	_, err := guc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guc *GroupUserCreate) ExecX(ctx context.Context) {
	if err := guc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guc *GroupUserCreate) check() error {
	if _, ok := guc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group_id", err: errors.New(`ent: missing required field "GroupUser.group_id"`)}
	}
	if _, ok := guc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "GroupUser.user_id"`)}
	}
	if _, ok := guc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "GroupUser.role"`)}
	}
	if v, ok := guc.mutation.Role(); ok {
		if err := groupuser.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "GroupUser.role": %w`, err)}
		}
	}
	if _, ok := guc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "GroupUser.group"`)}
	}
	if _, ok := guc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "GroupUser.user"`)}
	}
	return nil
}

func (guc *GroupUserCreate) sqlSave(ctx context.Context) (*GroupUser, error) {
	if err := guc.check(); err != nil {
		return nil, err
	}
	_node, _spec := guc.createSpec()
	if err := sqlgraph.CreateNode(ctx, guc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (guc *GroupUserCreate) createSpec() (*GroupUser, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupUser{config: guc.config}
		_spec = sqlgraph.NewCreateSpec(groupuser.Table, nil)
	)
	if value, ok := guc.mutation.Role(); ok {
		_spec.SetField(groupuser.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if nodes := guc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupuser.GroupTable,
			Columns: []string{groupuser.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.GroupID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := guc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   groupuser.UserTable,
			Columns: []string{groupuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupUserCreateBulk is the builder for creating many GroupUser entities in bulk.
type GroupUserCreateBulk struct {
	config
	builders []*GroupUserCreate
}

// Save creates the GroupUser entities in the database.
func (gucb *GroupUserCreateBulk) Save(ctx context.Context) ([]*GroupUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gucb.builders))
	nodes := make([]*GroupUser, len(gucb.builders))
	mutators := make([]Mutator, len(gucb.builders))
	for i := range gucb.builders {
		func(i int, root context.Context) {
			builder := gucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupUserMutation)
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
					_, err = mutators[i+1].Mutate(root, gucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
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
		if _, err := mutators[0].Mutate(ctx, gucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gucb *GroupUserCreateBulk) SaveX(ctx context.Context) []*GroupUser {
	v, err := gucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gucb *GroupUserCreateBulk) Exec(ctx context.Context) error {
	_, err := gucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gucb *GroupUserCreateBulk) ExecX(ctx context.Context) {
	if err := gucb.Exec(ctx); err != nil {
		panic(err)
	}
}