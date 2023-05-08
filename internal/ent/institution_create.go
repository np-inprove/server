// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/academicschool"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/user"
)

// InstitutionCreate is the builder for creating a Institution entity.
type InstitutionCreate struct {
	config
	mutation *InstitutionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (ic *InstitutionCreate) SetName(s string) *InstitutionCreate {
	ic.mutation.SetName(s)
	return ic
}

// AddAdminIDs adds the "admins" edge to the User entity by IDs.
func (ic *InstitutionCreate) AddAdminIDs(ids ...int) *InstitutionCreate {
	ic.mutation.AddAdminIDs(ids...)
	return ic
}

// AddAdmins adds the "admins" edges to the User entity.
func (ic *InstitutionCreate) AddAdmins(u ...*User) *InstitutionCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ic.AddAdminIDs(ids...)
}

// AddAcademicSchoolIDs adds the "academic_schools" edge to the AcademicSchool entity by IDs.
func (ic *InstitutionCreate) AddAcademicSchoolIDs(ids ...int) *InstitutionCreate {
	ic.mutation.AddAcademicSchoolIDs(ids...)
	return ic
}

// AddAcademicSchools adds the "academic_schools" edges to the AcademicSchool entity.
func (ic *InstitutionCreate) AddAcademicSchools(a ...*AcademicSchool) *InstitutionCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ic.AddAcademicSchoolIDs(ids...)
}

// Mutation returns the InstitutionMutation object of the builder.
func (ic *InstitutionCreate) Mutation() *InstitutionMutation {
	return ic.mutation
}

// Save creates the Institution in the database.
func (ic *InstitutionCreate) Save(ctx context.Context) (*Institution, error) {
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InstitutionCreate) SaveX(ctx context.Context) *Institution {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InstitutionCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InstitutionCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InstitutionCreate) check() error {
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Institution.name"`)}
	}
	if v, ok := ic.mutation.Name(); ok {
		if err := institution.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Institution.name": %w`, err)}
		}
	}
	return nil
}

func (ic *InstitutionCreate) sqlSave(ctx context.Context) (*Institution, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InstitutionCreate) createSpec() (*Institution, *sqlgraph.CreateSpec) {
	var (
		_node = &Institution{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(institution.Table, sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(institution.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := ic.mutation.AdminsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   institution.AdminsTable,
			Columns: institution.AdminsPrimaryKey,
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
	if nodes := ic.mutation.AcademicSchoolsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.AcademicSchoolsTable,
			Columns: []string{institution.AcademicSchoolsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(academicschool.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InstitutionCreateBulk is the builder for creating many Institution entities in bulk.
type InstitutionCreateBulk struct {
	config
	builders []*InstitutionCreate
}

// Save creates the Institution entities in the database.
func (icb *InstitutionCreateBulk) Save(ctx context.Context) ([]*Institution, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Institution, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InstitutionMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InstitutionCreateBulk) SaveX(ctx context.Context) []*Institution {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InstitutionCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InstitutionCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
