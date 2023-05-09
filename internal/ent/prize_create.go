// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/prize"
	"github.com/np-inprove/server/internal/ent/user"
)

// PrizeCreate is the builder for creating a Prize entity.
type PrizeCreate struct {
	config
	mutation *PrizeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *PrizeCreate) SetName(s string) *PrizeCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PrizeCreate) SetDescription(s string) *PrizeCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetPointsRequired sets the "points_required" field.
func (pc *PrizeCreate) SetPointsRequired(i int) *PrizeCreate {
	pc.mutation.SetPointsRequired(i)
	return pc
}

// SetDiscriminator sets the "discriminator" field.
func (pc *PrizeCreate) SetDiscriminator(pr prize.Discriminator) *PrizeCreate {
	pc.mutation.SetDiscriminator(pr)
	return pc
}

// SetInstitutionID sets the "institution" edge to the Institution entity by ID.
func (pc *PrizeCreate) SetInstitutionID(id int) *PrizeCreate {
	pc.mutation.SetInstitutionID(id)
	return pc
}

// SetNillableInstitutionID sets the "institution" edge to the Institution entity by ID if the given value is not nil.
func (pc *PrizeCreate) SetNillableInstitutionID(id *int) *PrizeCreate {
	if id != nil {
		pc = pc.SetInstitutionID(*id)
	}
	return pc
}

// SetInstitution sets the "institution" edge to the Institution entity.
func (pc *PrizeCreate) SetInstitution(i *Institution) *PrizeCreate {
	return pc.SetInstitutionID(i.ID)
}

// AddRedemptionUserIDs adds the "redemption_users" edge to the User entity by IDs.
func (pc *PrizeCreate) AddRedemptionUserIDs(ids ...int) *PrizeCreate {
	pc.mutation.AddRedemptionUserIDs(ids...)
	return pc
}

// AddRedemptionUsers adds the "redemption_users" edges to the User entity.
func (pc *PrizeCreate) AddRedemptionUsers(u ...*User) *PrizeCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddRedemptionUserIDs(ids...)
}

// Mutation returns the PrizeMutation object of the builder.
func (pc *PrizeCreate) Mutation() *PrizeMutation {
	return pc.mutation
}

// Save creates the Prize in the database.
func (pc *PrizeCreate) Save(ctx context.Context) (*Prize, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PrizeCreate) SaveX(ctx context.Context) *Prize {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PrizeCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PrizeCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PrizeCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Prize.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := prize.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Prize.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Prize.description"`)}
	}
	if _, ok := pc.mutation.PointsRequired(); !ok {
		return &ValidationError{Name: "points_required", err: errors.New(`ent: missing required field "Prize.points_required"`)}
	}
	if v, ok := pc.mutation.PointsRequired(); ok {
		if err := prize.PointsRequiredValidator(v); err != nil {
			return &ValidationError{Name: "points_required", err: fmt.Errorf(`ent: validator failed for field "Prize.points_required": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Discriminator(); !ok {
		return &ValidationError{Name: "discriminator", err: errors.New(`ent: missing required field "Prize.discriminator"`)}
	}
	if v, ok := pc.mutation.Discriminator(); ok {
		if err := prize.DiscriminatorValidator(v); err != nil {
			return &ValidationError{Name: "discriminator", err: fmt.Errorf(`ent: validator failed for field "Prize.discriminator": %w`, err)}
		}
	}
	return nil
}

func (pc *PrizeCreate) sqlSave(ctx context.Context) (*Prize, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PrizeCreate) createSpec() (*Prize, *sqlgraph.CreateSpec) {
	var (
		_node = &Prize{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(prize.Table, sqlgraph.NewFieldSpec(prize.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(prize.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(prize.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.PointsRequired(); ok {
		_spec.SetField(prize.FieldPointsRequired, field.TypeInt, value)
		_node.PointsRequired = value
	}
	if value, ok := pc.mutation.Discriminator(); ok {
		_spec.SetField(prize.FieldDiscriminator, field.TypeEnum, value)
		_node.Discriminator = value
	}
	if nodes := pc.mutation.InstitutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prize.InstitutionTable,
			Columns: []string{prize.InstitutionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.institution_prizes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.RedemptionUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   prize.RedemptionUsersTable,
			Columns: prize.RedemptionUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PrizeCreateBulk is the builder for creating many Prize entities in bulk.
type PrizeCreateBulk struct {
	config
	builders []*PrizeCreate
}

// Save creates the Prize entities in the database.
func (pcb *PrizeCreateBulk) Save(ctx context.Context) ([]*Prize, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Prize, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PrizeMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PrizeCreateBulk) SaveX(ctx context.Context) []*Prize {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PrizeCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PrizeCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
