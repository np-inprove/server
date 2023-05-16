// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/forumpost"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/user"
)

// ForumPostUpdate is the builder for updating ForumPost entities.
type ForumPostUpdate struct {
	config
	hooks    []Hook
	mutation *ForumPostMutation
}

// Where appends a list predicates to the ForumPostUpdate builder.
func (fpu *ForumPostUpdate) Where(ps ...predicate.ForumPost) *ForumPostUpdate {
	fpu.mutation.Where(ps...)
	return fpu
}

// SetTitle sets the "title" field.
func (fpu *ForumPostUpdate) SetTitle(s string) *ForumPostUpdate {
	fpu.mutation.SetTitle(s)
	return fpu
}

// SetContent sets the "content" field.
func (fpu *ForumPostUpdate) SetContent(s string) *ForumPostUpdate {
	fpu.mutation.SetContent(s)
	return fpu
}

// SetMentionedUsersJSON sets the "mentioned_users_json" field.
func (fpu *ForumPostUpdate) SetMentionedUsersJSON(i []int) *ForumPostUpdate {
	fpu.mutation.SetMentionedUsersJSON(i)
	return fpu
}

// AppendMentionedUsersJSON appends i to the "mentioned_users_json" field.
func (fpu *ForumPostUpdate) AppendMentionedUsersJSON(i []int) *ForumPostUpdate {
	fpu.mutation.AppendMentionedUsersJSON(i)
	return fpu
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (fpu *ForumPostUpdate) SetAuthorID(id int) *ForumPostUpdate {
	fpu.mutation.SetAuthorID(id)
	return fpu
}

// SetAuthor sets the "author" edge to the User entity.
func (fpu *ForumPostUpdate) SetAuthor(u *User) *ForumPostUpdate {
	return fpu.SetAuthorID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (fpu *ForumPostUpdate) SetGroupID(id int) *ForumPostUpdate {
	fpu.mutation.SetGroupID(id)
	return fpu
}

// SetGroup sets the "group" edge to the Group entity.
func (fpu *ForumPostUpdate) SetGroup(g *Group) *ForumPostUpdate {
	return fpu.SetGroupID(g.ID)
}

// SetParentID sets the "parent" edge to the ForumPost entity by ID.
func (fpu *ForumPostUpdate) SetParentID(id int) *ForumPostUpdate {
	fpu.mutation.SetParentID(id)
	return fpu
}

// SetNillableParentID sets the "parent" edge to the ForumPost entity by ID if the given value is not nil.
func (fpu *ForumPostUpdate) SetNillableParentID(id *int) *ForumPostUpdate {
	if id != nil {
		fpu = fpu.SetParentID(*id)
	}
	return fpu
}

// SetParent sets the "parent" edge to the ForumPost entity.
func (fpu *ForumPostUpdate) SetParent(f *ForumPost) *ForumPostUpdate {
	return fpu.SetParentID(f.ID)
}

// AddChildIDs adds the "children" edge to the ForumPost entity by IDs.
func (fpu *ForumPostUpdate) AddChildIDs(ids ...int) *ForumPostUpdate {
	fpu.mutation.AddChildIDs(ids...)
	return fpu
}

// AddChildren adds the "children" edges to the ForumPost entity.
func (fpu *ForumPostUpdate) AddChildren(f ...*ForumPost) *ForumPostUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpu.AddChildIDs(ids...)
}

// AddReactedUserIDs adds the "reacted_users" edge to the User entity by IDs.
func (fpu *ForumPostUpdate) AddReactedUserIDs(ids ...int) *ForumPostUpdate {
	fpu.mutation.AddReactedUserIDs(ids...)
	return fpu
}

// AddReactedUsers adds the "reacted_users" edges to the User entity.
func (fpu *ForumPostUpdate) AddReactedUsers(u ...*User) *ForumPostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return fpu.AddReactedUserIDs(ids...)
}

// Mutation returns the ForumPostMutation object of the builder.
func (fpu *ForumPostUpdate) Mutation() *ForumPostMutation {
	return fpu.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (fpu *ForumPostUpdate) ClearAuthor() *ForumPostUpdate {
	fpu.mutation.ClearAuthor()
	return fpu
}

// ClearGroup clears the "group" edge to the Group entity.
func (fpu *ForumPostUpdate) ClearGroup() *ForumPostUpdate {
	fpu.mutation.ClearGroup()
	return fpu
}

// ClearParent clears the "parent" edge to the ForumPost entity.
func (fpu *ForumPostUpdate) ClearParent() *ForumPostUpdate {
	fpu.mutation.ClearParent()
	return fpu
}

// ClearChildren clears all "children" edges to the ForumPost entity.
func (fpu *ForumPostUpdate) ClearChildren() *ForumPostUpdate {
	fpu.mutation.ClearChildren()
	return fpu
}

// RemoveChildIDs removes the "children" edge to ForumPost entities by IDs.
func (fpu *ForumPostUpdate) RemoveChildIDs(ids ...int) *ForumPostUpdate {
	fpu.mutation.RemoveChildIDs(ids...)
	return fpu
}

// RemoveChildren removes "children" edges to ForumPost entities.
func (fpu *ForumPostUpdate) RemoveChildren(f ...*ForumPost) *ForumPostUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpu.RemoveChildIDs(ids...)
}

// ClearReactedUsers clears all "reacted_users" edges to the User entity.
func (fpu *ForumPostUpdate) ClearReactedUsers() *ForumPostUpdate {
	fpu.mutation.ClearReactedUsers()
	return fpu
}

// RemoveReactedUserIDs removes the "reacted_users" edge to User entities by IDs.
func (fpu *ForumPostUpdate) RemoveReactedUserIDs(ids ...int) *ForumPostUpdate {
	fpu.mutation.RemoveReactedUserIDs(ids...)
	return fpu
}

// RemoveReactedUsers removes "reacted_users" edges to User entities.
func (fpu *ForumPostUpdate) RemoveReactedUsers(u ...*User) *ForumPostUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return fpu.RemoveReactedUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fpu *ForumPostUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fpu.sqlSave, fpu.mutation, fpu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fpu *ForumPostUpdate) SaveX(ctx context.Context) int {
	affected, err := fpu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fpu *ForumPostUpdate) Exec(ctx context.Context) error {
	_, err := fpu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fpu *ForumPostUpdate) ExecX(ctx context.Context) {
	if err := fpu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fpu *ForumPostUpdate) check() error {
	if v, ok := fpu.mutation.Title(); ok {
		if err := forumpost.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "ForumPost.title": %w`, err)}
		}
	}
	if v, ok := fpu.mutation.Content(); ok {
		if err := forumpost.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "ForumPost.content": %w`, err)}
		}
	}
	if _, ok := fpu.mutation.AuthorID(); fpu.mutation.AuthorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ForumPost.author"`)
	}
	if _, ok := fpu.mutation.GroupID(); fpu.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ForumPost.group"`)
	}
	return nil
}

func (fpu *ForumPostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fpu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(forumpost.Table, forumpost.Columns, sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt))
	if ps := fpu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fpu.mutation.Title(); ok {
		_spec.SetField(forumpost.FieldTitle, field.TypeString, value)
	}
	if value, ok := fpu.mutation.Content(); ok {
		_spec.SetField(forumpost.FieldContent, field.TypeString, value)
	}
	if value, ok := fpu.mutation.MentionedUsersJSON(); ok {
		_spec.SetField(forumpost.FieldMentionedUsersJSON, field.TypeJSON, value)
	}
	if value, ok := fpu.mutation.AppendedMentionedUsersJSON(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, forumpost.FieldMentionedUsersJSON, value)
		})
	}
	if fpu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   forumpost.AuthorTable,
			Columns: []string{forumpost.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   forumpost.AuthorTable,
			Columns: []string{forumpost.AuthorColumn},
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
	if fpu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forumpost.GroupTable,
			Columns: []string{forumpost.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forumpost.GroupTable,
			Columns: []string{forumpost.GroupColumn},
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
	if fpu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forumpost.ParentTable,
			Columns: []string{forumpost.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forumpost.ParentTable,
			Columns: []string{forumpost.ParentColumn},
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
	if fpu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forumpost.ChildrenTable,
			Columns: []string{forumpost.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !fpu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forumpost.ChildrenTable,
			Columns: []string{forumpost.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forumpost.ChildrenTable,
			Columns: []string{forumpost.ChildrenColumn},
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
	if fpu.mutation.ReactedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   forumpost.ReactedUsersTable,
			Columns: forumpost.ReactedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpu.mutation.RemovedReactedUsersIDs(); len(nodes) > 0 && !fpu.mutation.ReactedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   forumpost.ReactedUsersTable,
			Columns: forumpost.ReactedUsersPrimaryKey,
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
	if nodes := fpu.mutation.ReactedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   forumpost.ReactedUsersTable,
			Columns: forumpost.ReactedUsersPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, fpu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{forumpost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fpu.mutation.done = true
	return n, nil
}

// ForumPostUpdateOne is the builder for updating a single ForumPost entity.
type ForumPostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ForumPostMutation
}

// SetTitle sets the "title" field.
func (fpuo *ForumPostUpdateOne) SetTitle(s string) *ForumPostUpdateOne {
	fpuo.mutation.SetTitle(s)
	return fpuo
}

// SetContent sets the "content" field.
func (fpuo *ForumPostUpdateOne) SetContent(s string) *ForumPostUpdateOne {
	fpuo.mutation.SetContent(s)
	return fpuo
}

// SetMentionedUsersJSON sets the "mentioned_users_json" field.
func (fpuo *ForumPostUpdateOne) SetMentionedUsersJSON(i []int) *ForumPostUpdateOne {
	fpuo.mutation.SetMentionedUsersJSON(i)
	return fpuo
}

// AppendMentionedUsersJSON appends i to the "mentioned_users_json" field.
func (fpuo *ForumPostUpdateOne) AppendMentionedUsersJSON(i []int) *ForumPostUpdateOne {
	fpuo.mutation.AppendMentionedUsersJSON(i)
	return fpuo
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (fpuo *ForumPostUpdateOne) SetAuthorID(id int) *ForumPostUpdateOne {
	fpuo.mutation.SetAuthorID(id)
	return fpuo
}

// SetAuthor sets the "author" edge to the User entity.
func (fpuo *ForumPostUpdateOne) SetAuthor(u *User) *ForumPostUpdateOne {
	return fpuo.SetAuthorID(u.ID)
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (fpuo *ForumPostUpdateOne) SetGroupID(id int) *ForumPostUpdateOne {
	fpuo.mutation.SetGroupID(id)
	return fpuo
}

// SetGroup sets the "group" edge to the Group entity.
func (fpuo *ForumPostUpdateOne) SetGroup(g *Group) *ForumPostUpdateOne {
	return fpuo.SetGroupID(g.ID)
}

// SetParentID sets the "parent" edge to the ForumPost entity by ID.
func (fpuo *ForumPostUpdateOne) SetParentID(id int) *ForumPostUpdateOne {
	fpuo.mutation.SetParentID(id)
	return fpuo
}

// SetNillableParentID sets the "parent" edge to the ForumPost entity by ID if the given value is not nil.
func (fpuo *ForumPostUpdateOne) SetNillableParentID(id *int) *ForumPostUpdateOne {
	if id != nil {
		fpuo = fpuo.SetParentID(*id)
	}
	return fpuo
}

// SetParent sets the "parent" edge to the ForumPost entity.
func (fpuo *ForumPostUpdateOne) SetParent(f *ForumPost) *ForumPostUpdateOne {
	return fpuo.SetParentID(f.ID)
}

// AddChildIDs adds the "children" edge to the ForumPost entity by IDs.
func (fpuo *ForumPostUpdateOne) AddChildIDs(ids ...int) *ForumPostUpdateOne {
	fpuo.mutation.AddChildIDs(ids...)
	return fpuo
}

// AddChildren adds the "children" edges to the ForumPost entity.
func (fpuo *ForumPostUpdateOne) AddChildren(f ...*ForumPost) *ForumPostUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpuo.AddChildIDs(ids...)
}

// AddReactedUserIDs adds the "reacted_users" edge to the User entity by IDs.
func (fpuo *ForumPostUpdateOne) AddReactedUserIDs(ids ...int) *ForumPostUpdateOne {
	fpuo.mutation.AddReactedUserIDs(ids...)
	return fpuo
}

// AddReactedUsers adds the "reacted_users" edges to the User entity.
func (fpuo *ForumPostUpdateOne) AddReactedUsers(u ...*User) *ForumPostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return fpuo.AddReactedUserIDs(ids...)
}

// Mutation returns the ForumPostMutation object of the builder.
func (fpuo *ForumPostUpdateOne) Mutation() *ForumPostMutation {
	return fpuo.mutation
}

// ClearAuthor clears the "author" edge to the User entity.
func (fpuo *ForumPostUpdateOne) ClearAuthor() *ForumPostUpdateOne {
	fpuo.mutation.ClearAuthor()
	return fpuo
}

// ClearGroup clears the "group" edge to the Group entity.
func (fpuo *ForumPostUpdateOne) ClearGroup() *ForumPostUpdateOne {
	fpuo.mutation.ClearGroup()
	return fpuo
}

// ClearParent clears the "parent" edge to the ForumPost entity.
func (fpuo *ForumPostUpdateOne) ClearParent() *ForumPostUpdateOne {
	fpuo.mutation.ClearParent()
	return fpuo
}

// ClearChildren clears all "children" edges to the ForumPost entity.
func (fpuo *ForumPostUpdateOne) ClearChildren() *ForumPostUpdateOne {
	fpuo.mutation.ClearChildren()
	return fpuo
}

// RemoveChildIDs removes the "children" edge to ForumPost entities by IDs.
func (fpuo *ForumPostUpdateOne) RemoveChildIDs(ids ...int) *ForumPostUpdateOne {
	fpuo.mutation.RemoveChildIDs(ids...)
	return fpuo
}

// RemoveChildren removes "children" edges to ForumPost entities.
func (fpuo *ForumPostUpdateOne) RemoveChildren(f ...*ForumPost) *ForumPostUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fpuo.RemoveChildIDs(ids...)
}

// ClearReactedUsers clears all "reacted_users" edges to the User entity.
func (fpuo *ForumPostUpdateOne) ClearReactedUsers() *ForumPostUpdateOne {
	fpuo.mutation.ClearReactedUsers()
	return fpuo
}

// RemoveReactedUserIDs removes the "reacted_users" edge to User entities by IDs.
func (fpuo *ForumPostUpdateOne) RemoveReactedUserIDs(ids ...int) *ForumPostUpdateOne {
	fpuo.mutation.RemoveReactedUserIDs(ids...)
	return fpuo
}

// RemoveReactedUsers removes "reacted_users" edges to User entities.
func (fpuo *ForumPostUpdateOne) RemoveReactedUsers(u ...*User) *ForumPostUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return fpuo.RemoveReactedUserIDs(ids...)
}

// Where appends a list predicates to the ForumPostUpdate builder.
func (fpuo *ForumPostUpdateOne) Where(ps ...predicate.ForumPost) *ForumPostUpdateOne {
	fpuo.mutation.Where(ps...)
	return fpuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fpuo *ForumPostUpdateOne) Select(field string, fields ...string) *ForumPostUpdateOne {
	fpuo.fields = append([]string{field}, fields...)
	return fpuo
}

// Save executes the query and returns the updated ForumPost entity.
func (fpuo *ForumPostUpdateOne) Save(ctx context.Context) (*ForumPost, error) {
	return withHooks(ctx, fpuo.sqlSave, fpuo.mutation, fpuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fpuo *ForumPostUpdateOne) SaveX(ctx context.Context) *ForumPost {
	node, err := fpuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fpuo *ForumPostUpdateOne) Exec(ctx context.Context) error {
	_, err := fpuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fpuo *ForumPostUpdateOne) ExecX(ctx context.Context) {
	if err := fpuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fpuo *ForumPostUpdateOne) check() error {
	if v, ok := fpuo.mutation.Title(); ok {
		if err := forumpost.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "ForumPost.title": %w`, err)}
		}
	}
	if v, ok := fpuo.mutation.Content(); ok {
		if err := forumpost.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "ForumPost.content": %w`, err)}
		}
	}
	if _, ok := fpuo.mutation.AuthorID(); fpuo.mutation.AuthorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ForumPost.author"`)
	}
	if _, ok := fpuo.mutation.GroupID(); fpuo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ForumPost.group"`)
	}
	return nil
}

func (fpuo *ForumPostUpdateOne) sqlSave(ctx context.Context) (_node *ForumPost, err error) {
	if err := fpuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(forumpost.Table, forumpost.Columns, sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt))
	id, ok := fpuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ForumPost.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fpuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, forumpost.FieldID)
		for _, f := range fields {
			if !forumpost.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != forumpost.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fpuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fpuo.mutation.Title(); ok {
		_spec.SetField(forumpost.FieldTitle, field.TypeString, value)
	}
	if value, ok := fpuo.mutation.Content(); ok {
		_spec.SetField(forumpost.FieldContent, field.TypeString, value)
	}
	if value, ok := fpuo.mutation.MentionedUsersJSON(); ok {
		_spec.SetField(forumpost.FieldMentionedUsersJSON, field.TypeJSON, value)
	}
	if value, ok := fpuo.mutation.AppendedMentionedUsersJSON(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, forumpost.FieldMentionedUsersJSON, value)
		})
	}
	if fpuo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   forumpost.AuthorTable,
			Columns: []string{forumpost.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   forumpost.AuthorTable,
			Columns: []string{forumpost.AuthorColumn},
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
	if fpuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forumpost.GroupTable,
			Columns: []string{forumpost.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forumpost.GroupTable,
			Columns: []string{forumpost.GroupColumn},
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
	if fpuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forumpost.ParentTable,
			Columns: []string{forumpost.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   forumpost.ParentTable,
			Columns: []string{forumpost.ParentColumn},
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
	if fpuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forumpost.ChildrenTable,
			Columns: []string{forumpost.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !fpuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forumpost.ChildrenTable,
			Columns: []string{forumpost.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(forumpost.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   forumpost.ChildrenTable,
			Columns: []string{forumpost.ChildrenColumn},
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
	if fpuo.mutation.ReactedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   forumpost.ReactedUsersTable,
			Columns: forumpost.ReactedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fpuo.mutation.RemovedReactedUsersIDs(); len(nodes) > 0 && !fpuo.mutation.ReactedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   forumpost.ReactedUsersTable,
			Columns: forumpost.ReactedUsersPrimaryKey,
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
	if nodes := fpuo.mutation.ReactedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   forumpost.ReactedUsersTable,
			Columns: forumpost.ReactedUsersPrimaryKey,
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
	_node = &ForumPost{config: fpuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fpuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{forumpost.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fpuo.mutation.done = true
	return _node, nil
}