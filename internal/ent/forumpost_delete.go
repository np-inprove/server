// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/forumpost"
	"github.com/np-inprove/server/internal/ent/predicate"
)

// ForumPostDelete is the builder for deleting a ForumPost entity.
type ForumPostDelete struct {
	config
	hooks    []Hook
	mutation *ForumPostMutation
}

// Where appends a list predicates to the ForumPostDelete builder.
func (fpd *ForumPostDelete) Where(ps ...predicate.ForumPost) *ForumPostDelete {
	fpd.mutation.Where(ps...)
	return fpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fpd *ForumPostDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fpd.sqlExec, fpd.mutation, fpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fpd *ForumPostDelete) ExecX(ctx context.Context) int {
	n, err := fpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fpd *ForumPostDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(forumpost.Table, sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt))
	if ps := fpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fpd.mutation.done = true
	return affected, err
}

// ForumPostDeleteOne is the builder for deleting a single ForumPost entity.
type ForumPostDeleteOne struct {
	fpd *ForumPostDelete
}

// Where appends a list predicates to the ForumPostDelete builder.
func (fpdo *ForumPostDeleteOne) Where(ps ...predicate.ForumPost) *ForumPostDeleteOne {
	fpdo.fpd.mutation.Where(ps...)
	return fpdo
}

// Exec executes the deletion query.
func (fpdo *ForumPostDeleteOne) Exec(ctx context.Context) error {
	n, err := fpdo.fpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{forumpost.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fpdo *ForumPostDeleteOne) ExecX(ctx context.Context) {
	if err := fpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
