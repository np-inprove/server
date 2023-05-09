// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/pet"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/ent/userpet"
)

// UserPetCreate is the builder for creating a UserPet entity.
type UserPetCreate struct {
	config
	mutation *UserPetMutation
	hooks    []Hook
}

// SetHungerPercentage sets the "hunger_percentage" field.
func (upc *UserPetCreate) SetHungerPercentage(f float64) *UserPetCreate {
	upc.mutation.SetHungerPercentage(f)
	return upc
}

// SetEnabledSvgGroupElementIds sets the "enabled_svg_group_element_ids" field.
func (upc *UserPetCreate) SetEnabledSvgGroupElementIds(m map[string]bool) *UserPetCreate {
	upc.mutation.SetEnabledSvgGroupElementIds(m)
	return upc
}

// SetPetID sets the "pet_id" field.
func (upc *UserPetCreate) SetPetID(i int) *UserPetCreate {
	upc.mutation.SetPetID(i)
	return upc
}

// SetUserID sets the "user_id" field.
func (upc *UserPetCreate) SetUserID(i int) *UserPetCreate {
	upc.mutation.SetUserID(i)
	return upc
}

// SetPet sets the "pet" edge to the Pet entity.
func (upc *UserPetCreate) SetPet(p *Pet) *UserPetCreate {
	return upc.SetPetID(p.ID)
}

// SetUser sets the "user" edge to the User entity.
func (upc *UserPetCreate) SetUser(u *User) *UserPetCreate {
	return upc.SetUserID(u.ID)
}

// Mutation returns the UserPetMutation object of the builder.
func (upc *UserPetCreate) Mutation() *UserPetMutation {
	return upc.mutation
}

// Save creates the UserPet in the database.
func (upc *UserPetCreate) Save(ctx context.Context) (*UserPet, error) {
	upc.defaults()
	return withHooks(ctx, upc.sqlSave, upc.mutation, upc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (upc *UserPetCreate) SaveX(ctx context.Context) *UserPet {
	v, err := upc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upc *UserPetCreate) Exec(ctx context.Context) error {
	_, err := upc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upc *UserPetCreate) ExecX(ctx context.Context) {
	if err := upc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (upc *UserPetCreate) defaults() {
	if _, ok := upc.mutation.EnabledSvgGroupElementIds(); !ok {
		v := userpet.DefaultEnabledSvgGroupElementIds
		upc.mutation.SetEnabledSvgGroupElementIds(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (upc *UserPetCreate) check() error {
	if _, ok := upc.mutation.HungerPercentage(); !ok {
		return &ValidationError{Name: "hunger_percentage", err: errors.New(`ent: missing required field "UserPet.hunger_percentage"`)}
	}
	if v, ok := upc.mutation.HungerPercentage(); ok {
		if err := userpet.HungerPercentageValidator(v); err != nil {
			return &ValidationError{Name: "hunger_percentage", err: fmt.Errorf(`ent: validator failed for field "UserPet.hunger_percentage": %w`, err)}
		}
	}
	if _, ok := upc.mutation.EnabledSvgGroupElementIds(); !ok {
		return &ValidationError{Name: "enabled_svg_group_element_ids", err: errors.New(`ent: missing required field "UserPet.enabled_svg_group_element_ids"`)}
	}
	if _, ok := upc.mutation.PetID(); !ok {
		return &ValidationError{Name: "pet_id", err: errors.New(`ent: missing required field "UserPet.pet_id"`)}
	}
	if _, ok := upc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserPet.user_id"`)}
	}
	if _, ok := upc.mutation.PetID(); !ok {
		return &ValidationError{Name: "pet", err: errors.New(`ent: missing required edge "UserPet.pet"`)}
	}
	if _, ok := upc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserPet.user"`)}
	}
	return nil
}

func (upc *UserPetCreate) sqlSave(ctx context.Context) (*UserPet, error) {
	if err := upc.check(); err != nil {
		return nil, err
	}
	_node, _spec := upc.createSpec()
	if err := sqlgraph.CreateNode(ctx, upc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

func (upc *UserPetCreate) createSpec() (*UserPet, *sqlgraph.CreateSpec) {
	var (
		_node = &UserPet{config: upc.config}
		_spec = sqlgraph.NewCreateSpec(userpet.Table, nil)
	)
	if value, ok := upc.mutation.HungerPercentage(); ok {
		_spec.SetField(userpet.FieldHungerPercentage, field.TypeFloat64, value)
		_node.HungerPercentage = value
	}
	if value, ok := upc.mutation.EnabledSvgGroupElementIds(); ok {
		_spec.SetField(userpet.FieldEnabledSvgGroupElementIds, field.TypeJSON, value)
		_node.EnabledSvgGroupElementIds = value
	}
	if nodes := upc.mutation.PetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userpet.PetTable,
			Columns: []string{userpet.PetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.PetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := upc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   userpet.UserTable,
			Columns: []string{userpet.UserColumn},
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

// UserPetCreateBulk is the builder for creating many UserPet entities in bulk.
type UserPetCreateBulk struct {
	config
	builders []*UserPetCreate
}

// Save creates the UserPet entities in the database.
func (upcb *UserPetCreateBulk) Save(ctx context.Context) ([]*UserPet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(upcb.builders))
	nodes := make([]*UserPet, len(upcb.builders))
	mutators := make([]Mutator, len(upcb.builders))
	for i := range upcb.builders {
		func(i int, root context.Context) {
			builder := upcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserPetMutation)
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
					_, err = mutators[i+1].Mutate(root, upcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, upcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, upcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (upcb *UserPetCreateBulk) SaveX(ctx context.Context) []*UserPet {
	v, err := upcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (upcb *UserPetCreateBulk) Exec(ctx context.Context) error {
	_, err := upcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (upcb *UserPetCreateBulk) ExecX(ctx context.Context) {
	if err := upcb.Exec(ctx); err != nil {
		panic(err)
	}
}
