// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/academicschool"
	"github.com/np-inprove/server/internal/ent/predicate"
)

// AcademicSchoolDelete is the builder for deleting a AcademicSchool entity.
type AcademicSchoolDelete struct {
	config
	hooks    []Hook
	mutation *AcademicSchoolMutation
}

// Where appends a list predicates to the AcademicSchoolDelete builder.
func (asd *AcademicSchoolDelete) Where(ps ...predicate.AcademicSchool) *AcademicSchoolDelete {
	asd.mutation.Where(ps...)
	return asd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (asd *AcademicSchoolDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, asd.sqlExec, asd.mutation, asd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (asd *AcademicSchoolDelete) ExecX(ctx context.Context) int {
	n, err := asd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (asd *AcademicSchoolDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(academicschool.Table, sqlgraph.NewFieldSpec(academicschool.FieldID, field.TypeInt))
	if ps := asd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, asd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	asd.mutation.done = true
	return affected, err
}

// AcademicSchoolDeleteOne is the builder for deleting a single AcademicSchool entity.
type AcademicSchoolDeleteOne struct {
	asd *AcademicSchoolDelete
}

// Where appends a list predicates to the AcademicSchoolDelete builder.
func (asdo *AcademicSchoolDeleteOne) Where(ps ...predicate.AcademicSchool) *AcademicSchoolDeleteOne {
	asdo.asd.mutation.Where(ps...)
	return asdo
}

// Exec executes the deletion query.
func (asdo *AcademicSchoolDeleteOne) Exec(ctx context.Context) error {
	n, err := asdo.asd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{academicschool.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (asdo *AcademicSchoolDeleteOne) ExecX(ctx context.Context) {
	if err := asdo.Exec(ctx); err != nil {
		panic(err)
	}
}
