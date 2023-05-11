// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/redemption"
	"github.com/np-inprove/server/internal/ent/voucher"
)

// VoucherCreate is the builder for creating a Voucher entity.
type VoucherCreate struct {
	config
	mutation *VoucherMutation
	hooks    []Hook
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
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt),
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

// VoucherCreateBulk is the builder for creating many Voucher entities in bulk.
type VoucherCreateBulk struct {
	config
	builders []*VoucherCreate
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
