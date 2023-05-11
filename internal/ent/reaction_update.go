// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/forumpost"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/reaction"
	"github.com/np-inprove/server/internal/ent/user"
)

// ReactionUpdate is the builder for updating Reaction entities.
type ReactionUpdate struct {
	config
	hooks    []Hook
	mutation *ReactionMutation
}

// Where appends a list predicates to the ReactionUpdate builder.
func (ru *ReactionUpdate) Where(ps ...predicate.Reaction) *ReactionUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *ReactionUpdate) SetUserID(i int) *ReactionUpdate {
	ru.mutation.SetUserID(i)
	return ru
}

// SetForumPostID sets the "forum_post_id" field.
func (ru *ReactionUpdate) SetForumPostID(i int) *ReactionUpdate {
	ru.mutation.SetForumPostID(i)
	return ru
}

// SetEmoji sets the "emoji" field.
func (ru *ReactionUpdate) SetEmoji(s string) *ReactionUpdate {
	ru.mutation.SetEmoji(s)
	return ru
}

// SetUser sets the "user" edge to the User entity.
func (ru *ReactionUpdate) SetUser(u *User) *ReactionUpdate {
	return ru.SetUserID(u.ID)
}

// SetForumPost sets the "forum_post" edge to the ForumPost entity.
func (ru *ReactionUpdate) SetForumPost(f *ForumPost) *ReactionUpdate {
	return ru.SetForumPostID(f.ID)
}

// Mutation returns the ReactionMutation object of the builder.
func (ru *ReactionUpdate) Mutation() *ReactionMutation {
	return ru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ru *ReactionUpdate) ClearUser() *ReactionUpdate {
	ru.mutation.ClearUser()
	return ru
}

// ClearForumPost clears the "forum_post" edge to the ForumPost entity.
func (ru *ReactionUpdate) ClearForumPost() *ReactionUpdate {
	ru.mutation.ClearForumPost()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReactionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReactionUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReactionUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReactionUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReactionUpdate) check() error {
	if v, ok := ru.mutation.Emoji(); ok {
		if err := reaction.EmojiValidator(v); err != nil {
			return &ValidationError{Name: "emoji", err: fmt.Errorf(`ent: validator failed for field "Reaction.emoji": %w`, err)}
		}
	}
	if _, ok := ru.mutation.UserID(); ru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reaction.user"`)
	}
	if _, ok := ru.mutation.ForumPostID(); ru.mutation.ForumPostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reaction.forum_post"`)
	}
	return nil
}

func (ru *ReactionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(reaction.Table, reaction.Columns, sqlgraph.NewFieldSpec(reaction.FieldForumPostID, field.TypeInt), sqlgraph.NewFieldSpec(reaction.FieldUserID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Emoji(); ok {
		_spec.SetField(reaction.FieldEmoji, field.TypeString, value)
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ForumPostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.ForumPostTable,
			Columns: []string{reaction.ForumPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ForumPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.ForumPostTable,
			Columns: []string{reaction.ForumPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// ReactionUpdateOne is the builder for updating a single Reaction entity.
type ReactionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReactionMutation
}

// SetUserID sets the "user_id" field.
func (ruo *ReactionUpdateOne) SetUserID(i int) *ReactionUpdateOne {
	ruo.mutation.SetUserID(i)
	return ruo
}

// SetForumPostID sets the "forum_post_id" field.
func (ruo *ReactionUpdateOne) SetForumPostID(i int) *ReactionUpdateOne {
	ruo.mutation.SetForumPostID(i)
	return ruo
}

// SetEmoji sets the "emoji" field.
func (ruo *ReactionUpdateOne) SetEmoji(s string) *ReactionUpdateOne {
	ruo.mutation.SetEmoji(s)
	return ruo
}

// SetUser sets the "user" edge to the User entity.
func (ruo *ReactionUpdateOne) SetUser(u *User) *ReactionUpdateOne {
	return ruo.SetUserID(u.ID)
}

// SetForumPost sets the "forum_post" edge to the ForumPost entity.
func (ruo *ReactionUpdateOne) SetForumPost(f *ForumPost) *ReactionUpdateOne {
	return ruo.SetForumPostID(f.ID)
}

// Mutation returns the ReactionMutation object of the builder.
func (ruo *ReactionUpdateOne) Mutation() *ReactionMutation {
	return ruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ruo *ReactionUpdateOne) ClearUser() *ReactionUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// ClearForumPost clears the "forum_post" edge to the ForumPost entity.
func (ruo *ReactionUpdateOne) ClearForumPost() *ReactionUpdateOne {
	ruo.mutation.ClearForumPost()
	return ruo
}

// Where appends a list predicates to the ReactionUpdate builder.
func (ruo *ReactionUpdateOne) Where(ps ...predicate.Reaction) *ReactionUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReactionUpdateOne) Select(field string, fields ...string) *ReactionUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Reaction entity.
func (ruo *ReactionUpdateOne) Save(ctx context.Context) (*Reaction, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReactionUpdateOne) SaveX(ctx context.Context) *Reaction {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReactionUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReactionUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReactionUpdateOne) check() error {
	if v, ok := ruo.mutation.Emoji(); ok {
		if err := reaction.EmojiValidator(v); err != nil {
			return &ValidationError{Name: "emoji", err: fmt.Errorf(`ent: validator failed for field "Reaction.emoji": %w`, err)}
		}
	}
	if _, ok := ruo.mutation.UserID(); ruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reaction.user"`)
	}
	if _, ok := ruo.mutation.ForumPostID(); ruo.mutation.ForumPostCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Reaction.forum_post"`)
	}
	return nil
}

func (ruo *ReactionUpdateOne) sqlSave(ctx context.Context) (_node *Reaction, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(reaction.Table, reaction.Columns, sqlgraph.NewFieldSpec(reaction.FieldForumPostID, field.TypeInt), sqlgraph.NewFieldSpec(reaction.FieldUserID, field.TypeInt))
	if id, ok := ruo.mutation.ForumPostID(); !ok {
		return nil, &ValidationError{Name: "forum_post_id", err: errors.New(`ent: missing "Reaction.forum_post_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ruo.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "user_id", err: errors.New(`ent: missing "Reaction.user_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !reaction.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Emoji(); ok {
		_spec.SetField(reaction.FieldEmoji, field.TypeString, value)
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.UserTable,
			Columns: []string{reaction.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ForumPostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.ForumPostTable,
			Columns: []string{reaction.ForumPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ForumPostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   reaction.ForumPostTable,
			Columns: []string{reaction.ForumPostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Reaction{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reaction.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
