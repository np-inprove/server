// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/deadline"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/user"
)

// DeadlineUpdate is the builder for updating Deadline entities.
type DeadlineUpdate struct {
	config
	hooks    []Hook
	mutation *DeadlineMutation
}

// Where appends a list predicates to the DeadlineUpdate builder.
func (du *DeadlineUpdate) Where(ps ...predicate.Deadline) *DeadlineUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetName sets the "name" field.
func (du *DeadlineUpdate) SetName(s string) *DeadlineUpdate {
	du.mutation.SetName(s)
	return du
}

// SetDueTime sets the "due_time" field.
func (du *DeadlineUpdate) SetDueTime(t time.Time) *DeadlineUpdate {
	du.mutation.SetDueTime(t)
	return du
}

// SetNillableDueTime sets the "due_time" field if the given value is not nil.
func (du *DeadlineUpdate) SetNillableDueTime(t *time.Time) *DeadlineUpdate {
	if t != nil {
		du.SetDueTime(*t)
	}
	return du
}

// ClearDueTime clears the value of the "due_time" field.
func (du *DeadlineUpdate) ClearDueTime() *DeadlineUpdate {
	du.mutation.ClearDueTime()
	return du
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (du *DeadlineUpdate) SetAuthorID(id int) *DeadlineUpdate {
	du.mutation.SetAuthorID(id)
	return du
}

// SetAuthor sets the "author" edge to the User entity.
func (du *DeadlineUpdate) SetAuthor(u *User) *DeadlineUpdate {
	return du.SetAuthorID(u.ID)
}

// AddVotedUserIDs adds the "voted_users" edge to the User entity by IDs.
func (du *DeadlineUpdate) AddVotedUserIDs(ids ...int) *DeadlineUpdate {
	du.mutation.AddVotedUserIDs(ids...)
	return du
}

// AddVotedUsers adds the "voted_users" edges to the User entity.
func (du *DeadlineUpdate) AddVotedUsers(u ...*User) *DeadlineUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.AddVotedUserIDs(ids...)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (du *DeadlineUpdate) SetGroupID(id int) *DeadlineUpdate {
	du.mutation.SetGroupID(id)
	return du
}

// SetGroup sets the "group" edge to the Group entity.
func (du *DeadlineUpdate) SetGroup(g *Group) *DeadlineUpdate {
	return du.SetGroupID(g.ID)
}

// Mutation returns the DeadlineMutation object of the builder.
func (du *DeadlineUpdate) Mutation() *DeadlineMutation {
	return du.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (du *DeadlineUpdate) ClearAuthor() *DeadlineUpdate {
	du.mutation.ClearAuthor()
	return du
}

// ClearVotedUsers clears all "voted_users" edges to the User entity.
func (du *DeadlineUpdate) ClearVotedUsers() *DeadlineUpdate {
	du.mutation.ClearVotedUsers()
	return du
}

// RemoveVotedUserIDs removes the "voted_users" edge to User entities by IDs.
func (du *DeadlineUpdate) RemoveVotedUserIDs(ids ...int) *DeadlineUpdate {
	du.mutation.RemoveVotedUserIDs(ids...)
	return du
}

// RemoveVotedUsers removes "voted_users" edges to User entities.
func (du *DeadlineUpdate) RemoveVotedUsers(u ...*User) *DeadlineUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return du.RemoveVotedUserIDs(ids...)
}

// ClearGroup clears the "group" edge to the Group entity.
func (du *DeadlineUpdate) ClearGroup() *DeadlineUpdate {
	du.mutation.ClearGroup()
	return du
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeadlineUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeadlineUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeadlineUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeadlineUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (du *DeadlineUpdate) check() error {
	if v, ok := du.mutation.Name(); ok {
		if err := deadline.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Deadline.name": %w`, err)}
		}
	}
	if _, ok := du.mutation.AuthorID(); du.mutation.AuthorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Deadline.author"`)
	}
	if _, ok := du.mutation.GroupID(); du.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Deadline.group"`)
	}
	return nil
}

func (du *DeadlineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := du.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(deadline.Table, deadline.Columns, sqlgraph.NewFieldSpec(deadline.FieldID, field.TypeInt))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(deadline.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.DueTime(); ok {
		_spec.SetField(deadline.FieldDueTime, field.TypeTime, value)
	}
	if du.mutation.DueTimeCleared() {
		_spec.ClearField(deadline.FieldDueTime, field.TypeTime)
	}
	if du.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deadline.AuthorTable,
			Columns: []string{deadline.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deadline.AuthorTable,
			Columns: []string{deadline.AuthorColumn},
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
	if du.mutation.VotedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   deadline.VotedUsersTable,
			Columns: deadline.VotedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedVotedUsersIDs(); len(nodes) > 0 && !du.mutation.VotedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   deadline.VotedUsersTable,
			Columns: deadline.VotedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.VotedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   deadline.VotedUsersTable,
			Columns: deadline.VotedUsersPrimaryKey,
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
	if du.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deadline.GroupTable,
			Columns: []string{deadline.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deadline.GroupTable,
			Columns: []string{deadline.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deadline.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DeadlineUpdateOne is the builder for updating a single Deadline entity.
type DeadlineUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeadlineMutation
}

// SetName sets the "name" field.
func (duo *DeadlineUpdateOne) SetName(s string) *DeadlineUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetDueTime sets the "due_time" field.
func (duo *DeadlineUpdateOne) SetDueTime(t time.Time) *DeadlineUpdateOne {
	duo.mutation.SetDueTime(t)
	return duo
}

// SetNillableDueTime sets the "due_time" field if the given value is not nil.
func (duo *DeadlineUpdateOne) SetNillableDueTime(t *time.Time) *DeadlineUpdateOne {
	if t != nil {
		duo.SetDueTime(*t)
	}
	return duo
}

// ClearDueTime clears the value of the "due_time" field.
func (duo *DeadlineUpdateOne) ClearDueTime() *DeadlineUpdateOne {
	duo.mutation.ClearDueTime()
	return duo
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (duo *DeadlineUpdateOne) SetAuthorID(id int) *DeadlineUpdateOne {
	duo.mutation.SetAuthorID(id)
	return duo
}

// SetAuthor sets the "author" edge to the User entity.
func (duo *DeadlineUpdateOne) SetAuthor(u *User) *DeadlineUpdateOne {
	return duo.SetAuthorID(u.ID)
}

// AddVotedUserIDs adds the "voted_users" edge to the User entity by IDs.
func (duo *DeadlineUpdateOne) AddVotedUserIDs(ids ...int) *DeadlineUpdateOne {
	duo.mutation.AddVotedUserIDs(ids...)
	return duo
}

// AddVotedUsers adds the "voted_users" edges to the User entity.
func (duo *DeadlineUpdateOne) AddVotedUsers(u ...*User) *DeadlineUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.AddVotedUserIDs(ids...)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (duo *DeadlineUpdateOne) SetGroupID(id int) *DeadlineUpdateOne {
	duo.mutation.SetGroupID(id)
	return duo
}

// SetGroup sets the "group" edge to the Group entity.
func (duo *DeadlineUpdateOne) SetGroup(g *Group) *DeadlineUpdateOne {
	return duo.SetGroupID(g.ID)
}

// Mutation returns the DeadlineMutation object of the builder.
func (duo *DeadlineUpdateOne) Mutation() *DeadlineMutation {
	return duo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (duo *DeadlineUpdateOne) ClearAuthor() *DeadlineUpdateOne {
	duo.mutation.ClearAuthor()
	return duo
}

// ClearVotedUsers clears all "voted_users" edges to the User entity.
func (duo *DeadlineUpdateOne) ClearVotedUsers() *DeadlineUpdateOne {
	duo.mutation.ClearVotedUsers()
	return duo
}

// RemoveVotedUserIDs removes the "voted_users" edge to User entities by IDs.
func (duo *DeadlineUpdateOne) RemoveVotedUserIDs(ids ...int) *DeadlineUpdateOne {
	duo.mutation.RemoveVotedUserIDs(ids...)
	return duo
}

// RemoveVotedUsers removes "voted_users" edges to User entities.
func (duo *DeadlineUpdateOne) RemoveVotedUsers(u ...*User) *DeadlineUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return duo.RemoveVotedUserIDs(ids...)
}

// ClearGroup clears the "group" edge to the Group entity.
func (duo *DeadlineUpdateOne) ClearGroup() *DeadlineUpdateOne {
	duo.mutation.ClearGroup()
	return duo
}

// Where appends a list predicates to the DeadlineUpdate builder.
func (duo *DeadlineUpdateOne) Where(ps ...predicate.Deadline) *DeadlineUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeadlineUpdateOne) Select(field string, fields ...string) *DeadlineUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Deadline entity.
func (duo *DeadlineUpdateOne) Save(ctx context.Context) (*Deadline, error) {
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeadlineUpdateOne) SaveX(ctx context.Context) *Deadline {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeadlineUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeadlineUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (duo *DeadlineUpdateOne) check() error {
	if v, ok := duo.mutation.Name(); ok {
		if err := deadline.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Deadline.name": %w`, err)}
		}
	}
	if _, ok := duo.mutation.AuthorID(); duo.mutation.AuthorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Deadline.author"`)
	}
	if _, ok := duo.mutation.GroupID(); duo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Deadline.group"`)
	}
	return nil
}

func (duo *DeadlineUpdateOne) sqlSave(ctx context.Context) (_node *Deadline, err error) {
	if err := duo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(deadline.Table, deadline.Columns, sqlgraph.NewFieldSpec(deadline.FieldID, field.TypeInt))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Deadline.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, deadline.FieldID)
		for _, f := range fields {
			if !deadline.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != deadline.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(deadline.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.DueTime(); ok {
		_spec.SetField(deadline.FieldDueTime, field.TypeTime, value)
	}
	if duo.mutation.DueTimeCleared() {
		_spec.ClearField(deadline.FieldDueTime, field.TypeTime)
	}
	if duo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deadline.AuthorTable,
			Columns: []string{deadline.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   deadline.AuthorTable,
			Columns: []string{deadline.AuthorColumn},
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
	if duo.mutation.VotedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   deadline.VotedUsersTable,
			Columns: deadline.VotedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedVotedUsersIDs(); len(nodes) > 0 && !duo.mutation.VotedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   deadline.VotedUsersTable,
			Columns: deadline.VotedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.VotedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   deadline.VotedUsersTable,
			Columns: deadline.VotedUsersPrimaryKey,
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
	if duo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deadline.GroupTable,
			Columns: []string{deadline.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   deadline.GroupTable,
			Columns: []string{deadline.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Deadline{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{deadline.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
