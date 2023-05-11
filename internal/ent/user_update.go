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
	"github.com/np-inprove/server/internal/ent/course"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/pet"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/redemption"
	"github.com/np-inprove/server/internal/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetPasswordHash sets the "password_hash" field.
func (uu *UserUpdate) SetPasswordHash(s string) *UserUpdate {
	uu.mutation.SetPasswordHash(s)
	return uu
}

// SetPoints sets the "points" field.
func (uu *UserUpdate) SetPoints(i int) *UserUpdate {
	uu.mutation.ResetPoints()
	uu.mutation.SetPoints(i)
	return uu
}

// AddPoints adds i to the "points" field.
func (uu *UserUpdate) AddPoints(i int) *UserUpdate {
	uu.mutation.AddPoints(i)
	return uu
}

// SetPointsAwardedCount sets the "points_awarded_count" field.
func (uu *UserUpdate) SetPointsAwardedCount(i int) *UserUpdate {
	uu.mutation.ResetPointsAwardedCount()
	uu.mutation.SetPointsAwardedCount(i)
	return uu
}

// AddPointsAwardedCount adds i to the "points_awarded_count" field.
func (uu *UserUpdate) AddPointsAwardedCount(i int) *UserUpdate {
	uu.mutation.AddPointsAwardedCount(i)
	return uu
}

// SetPointsAwardedResetTime sets the "points_awarded_reset_time" field.
func (uu *UserUpdate) SetPointsAwardedResetTime(t time.Time) *UserUpdate {
	uu.mutation.SetPointsAwardedResetTime(t)
	return uu
}

// SetNillablePointsAwardedResetTime sets the "points_awarded_reset_time" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePointsAwardedResetTime(t *time.Time) *UserUpdate {
	if t != nil {
		uu.SetPointsAwardedResetTime(*t)
	}
	return uu
}

// ClearPointsAwardedResetTime clears the value of the "points_awarded_reset_time" field.
func (uu *UserUpdate) ClearPointsAwardedResetTime() *UserUpdate {
	uu.mutation.ClearPointsAwardedResetTime()
	return uu
}

// SetGodMode sets the "god_mode" field.
func (uu *UserUpdate) SetGodMode(b bool) *UserUpdate {
	uu.mutation.SetGodMode(b)
	return uu
}

// AddInstitutionIDs adds the "institution" edge to the Institution entity by IDs.
func (uu *UserUpdate) AddInstitutionIDs(ids ...int) *UserUpdate {
	uu.mutation.AddInstitutionIDs(ids...)
	return uu
}

// AddInstitution adds the "institution" edges to the Institution entity.
func (uu *UserUpdate) AddInstitution(i ...*Institution) *UserUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uu.AddInstitutionIDs(ids...)
}

// SetCourseID sets the "course" edge to the Course entity by ID.
func (uu *UserUpdate) SetCourseID(id int) *UserUpdate {
	uu.mutation.SetCourseID(id)
	return uu
}

// SetNillableCourseID sets the "course" edge to the Course entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableCourseID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetCourseID(*id)
	}
	return uu
}

// SetCourse sets the "course" edge to the Course entity.
func (uu *UserUpdate) SetCourse(c *Course) *UserUpdate {
	return uu.SetCourseID(c.ID)
}

// AddRedemptionIDs adds the "redemptions" edge to the Redemption entity by IDs.
func (uu *UserUpdate) AddRedemptionIDs(ids ...int) *UserUpdate {
	uu.mutation.AddRedemptionIDs(ids...)
	return uu
}

// AddRedemptions adds the "redemptions" edges to the Redemption entity.
func (uu *UserUpdate) AddRedemptions(r ...*Redemption) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.AddRedemptionIDs(ids...)
}

// AddPetIDs adds the "pet" edge to the Pet entity by IDs.
func (uu *UserUpdate) AddPetIDs(ids ...int) *UserUpdate {
	uu.mutation.AddPetIDs(ids...)
	return uu
}

// AddPet adds the "pet" edges to the Pet entity.
func (uu *UserUpdate) AddPet(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPetIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uu *UserUpdate) AddGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.AddGroupIDs(ids...)
	return uu
}

// AddGroups adds the "groups" edges to the Group entity.
func (uu *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearInstitution clears all "institution" edges to the Institution entity.
func (uu *UserUpdate) ClearInstitution() *UserUpdate {
	uu.mutation.ClearInstitution()
	return uu
}

// RemoveInstitutionIDs removes the "institution" edge to Institution entities by IDs.
func (uu *UserUpdate) RemoveInstitutionIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveInstitutionIDs(ids...)
	return uu
}

// RemoveInstitution removes "institution" edges to Institution entities.
func (uu *UserUpdate) RemoveInstitution(i ...*Institution) *UserUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uu.RemoveInstitutionIDs(ids...)
}

// ClearCourse clears the "course" edge to the Course entity.
func (uu *UserUpdate) ClearCourse() *UserUpdate {
	uu.mutation.ClearCourse()
	return uu
}

// ClearRedemptions clears all "redemptions" edges to the Redemption entity.
func (uu *UserUpdate) ClearRedemptions() *UserUpdate {
	uu.mutation.ClearRedemptions()
	return uu
}

// RemoveRedemptionIDs removes the "redemptions" edge to Redemption entities by IDs.
func (uu *UserUpdate) RemoveRedemptionIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveRedemptionIDs(ids...)
	return uu
}

// RemoveRedemptions removes "redemptions" edges to Redemption entities.
func (uu *UserUpdate) RemoveRedemptions(r ...*Redemption) *UserUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uu.RemoveRedemptionIDs(ids...)
}

// ClearPet clears all "pet" edges to the Pet entity.
func (uu *UserUpdate) ClearPet() *UserUpdate {
	uu.mutation.ClearPet()
	return uu
}

// RemovePetIDs removes the "pet" edge to Pet entities by IDs.
func (uu *UserUpdate) RemovePetIDs(ids ...int) *UserUpdate {
	uu.mutation.RemovePetIDs(ids...)
	return uu
}

// RemovePet removes "pet" edges to Pet entities.
func (uu *UserUpdate) RemovePet(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePetIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uu *UserUpdate) ClearGroups() *UserUpdate {
	uu.mutation.ClearGroups()
	return uu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uu *UserUpdate) RemoveGroupIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveGroupIDs(ids...)
	return uu
}

// RemoveGroups removes "groups" edges to Group entities.
func (uu *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uu.mutation.PasswordHash(); ok {
		if err := user.PasswordHashValidator(v); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf(`ent: validator failed for field "User.password_hash": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Points(); ok {
		if err := user.PointsValidator(v); err != nil {
			return &ValidationError{Name: "points", err: fmt.Errorf(`ent: validator failed for field "User.points": %w`, err)}
		}
	}
	if v, ok := uu.mutation.PointsAwardedCount(); ok {
		if err := user.PointsAwardedCountValidator(v); err != nil {
			return &ValidationError{Name: "points_awarded_count", err: fmt.Errorf(`ent: validator failed for field "User.points_awarded_count": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
	}
	if value, ok := uu.mutation.Points(); ok {
		_spec.SetField(user.FieldPoints, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedPoints(); ok {
		_spec.AddField(user.FieldPoints, field.TypeInt, value)
	}
	if value, ok := uu.mutation.PointsAwardedCount(); ok {
		_spec.SetField(user.FieldPointsAwardedCount, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedPointsAwardedCount(); ok {
		_spec.AddField(user.FieldPointsAwardedCount, field.TypeInt, value)
	}
	if value, ok := uu.mutation.PointsAwardedResetTime(); ok {
		_spec.SetField(user.FieldPointsAwardedResetTime, field.TypeTime, value)
	}
	if uu.mutation.PointsAwardedResetTimeCleared() {
		_spec.ClearField(user.FieldPointsAwardedResetTime, field.TypeTime)
	}
	if value, ok := uu.mutation.GodMode(); ok {
		_spec.SetField(user.FieldGodMode, field.TypeBool, value)
	}
	if uu.mutation.InstitutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.InstitutionTable,
			Columns: user.InstitutionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedInstitutionIDs(); len(nodes) > 0 && !uu.mutation.InstitutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.InstitutionTable,
			Columns: user.InstitutionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.InstitutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.InstitutionTable,
			Columns: user.InstitutionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CourseTable,
			Columns: []string{user.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CourseTable,
			Columns: []string{user.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.RedemptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.RedemptionsTable,
			Columns: []string{user.RedemptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redemption.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedRedemptionsIDs(); len(nodes) > 0 && !uu.mutation.RedemptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.RedemptionsTable,
			Columns: []string{user.RedemptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redemption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RedemptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.RedemptionsTable,
			Columns: []string{user.RedemptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redemption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.PetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.PetTable,
			Columns: user.PetPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		createE := &UserPetCreate{config: uu.config, mutation: newUserPetMutation(uu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedPetIDs(); len(nodes) > 0 && !uu.mutation.PetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.PetTable,
			Columns: user.PetPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserPetCreate{config: uu.config, mutation: newUserPetMutation(uu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.PetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.PetTable,
			Columns: user.PetPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserPetCreate{config: uu.config, mutation: newUserPetMutation(uu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
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
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetPasswordHash sets the "password_hash" field.
func (uuo *UserUpdateOne) SetPasswordHash(s string) *UserUpdateOne {
	uuo.mutation.SetPasswordHash(s)
	return uuo
}

// SetPoints sets the "points" field.
func (uuo *UserUpdateOne) SetPoints(i int) *UserUpdateOne {
	uuo.mutation.ResetPoints()
	uuo.mutation.SetPoints(i)
	return uuo
}

// AddPoints adds i to the "points" field.
func (uuo *UserUpdateOne) AddPoints(i int) *UserUpdateOne {
	uuo.mutation.AddPoints(i)
	return uuo
}

// SetPointsAwardedCount sets the "points_awarded_count" field.
func (uuo *UserUpdateOne) SetPointsAwardedCount(i int) *UserUpdateOne {
	uuo.mutation.ResetPointsAwardedCount()
	uuo.mutation.SetPointsAwardedCount(i)
	return uuo
}

// AddPointsAwardedCount adds i to the "points_awarded_count" field.
func (uuo *UserUpdateOne) AddPointsAwardedCount(i int) *UserUpdateOne {
	uuo.mutation.AddPointsAwardedCount(i)
	return uuo
}

// SetPointsAwardedResetTime sets the "points_awarded_reset_time" field.
func (uuo *UserUpdateOne) SetPointsAwardedResetTime(t time.Time) *UserUpdateOne {
	uuo.mutation.SetPointsAwardedResetTime(t)
	return uuo
}

// SetNillablePointsAwardedResetTime sets the "points_awarded_reset_time" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePointsAwardedResetTime(t *time.Time) *UserUpdateOne {
	if t != nil {
		uuo.SetPointsAwardedResetTime(*t)
	}
	return uuo
}

// ClearPointsAwardedResetTime clears the value of the "points_awarded_reset_time" field.
func (uuo *UserUpdateOne) ClearPointsAwardedResetTime() *UserUpdateOne {
	uuo.mutation.ClearPointsAwardedResetTime()
	return uuo
}

// SetGodMode sets the "god_mode" field.
func (uuo *UserUpdateOne) SetGodMode(b bool) *UserUpdateOne {
	uuo.mutation.SetGodMode(b)
	return uuo
}

// AddInstitutionIDs adds the "institution" edge to the Institution entity by IDs.
func (uuo *UserUpdateOne) AddInstitutionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddInstitutionIDs(ids...)
	return uuo
}

// AddInstitution adds the "institution" edges to the Institution entity.
func (uuo *UserUpdateOne) AddInstitution(i ...*Institution) *UserUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uuo.AddInstitutionIDs(ids...)
}

// SetCourseID sets the "course" edge to the Course entity by ID.
func (uuo *UserUpdateOne) SetCourseID(id int) *UserUpdateOne {
	uuo.mutation.SetCourseID(id)
	return uuo
}

// SetNillableCourseID sets the "course" edge to the Course entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCourseID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetCourseID(*id)
	}
	return uuo
}

// SetCourse sets the "course" edge to the Course entity.
func (uuo *UserUpdateOne) SetCourse(c *Course) *UserUpdateOne {
	return uuo.SetCourseID(c.ID)
}

// AddRedemptionIDs adds the "redemptions" edge to the Redemption entity by IDs.
func (uuo *UserUpdateOne) AddRedemptionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddRedemptionIDs(ids...)
	return uuo
}

// AddRedemptions adds the "redemptions" edges to the Redemption entity.
func (uuo *UserUpdateOne) AddRedemptions(r ...*Redemption) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.AddRedemptionIDs(ids...)
}

// AddPetIDs adds the "pet" edge to the Pet entity by IDs.
func (uuo *UserUpdateOne) AddPetIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddPetIDs(ids...)
	return uuo
}

// AddPet adds the "pet" edges to the Pet entity.
func (uuo *UserUpdateOne) AddPet(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPetIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddGroupIDs(ids...)
	return uuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (uuo *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearInstitution clears all "institution" edges to the Institution entity.
func (uuo *UserUpdateOne) ClearInstitution() *UserUpdateOne {
	uuo.mutation.ClearInstitution()
	return uuo
}

// RemoveInstitutionIDs removes the "institution" edge to Institution entities by IDs.
func (uuo *UserUpdateOne) RemoveInstitutionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveInstitutionIDs(ids...)
	return uuo
}

// RemoveInstitution removes "institution" edges to Institution entities.
func (uuo *UserUpdateOne) RemoveInstitution(i ...*Institution) *UserUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return uuo.RemoveInstitutionIDs(ids...)
}

// ClearCourse clears the "course" edge to the Course entity.
func (uuo *UserUpdateOne) ClearCourse() *UserUpdateOne {
	uuo.mutation.ClearCourse()
	return uuo
}

// ClearRedemptions clears all "redemptions" edges to the Redemption entity.
func (uuo *UserUpdateOne) ClearRedemptions() *UserUpdateOne {
	uuo.mutation.ClearRedemptions()
	return uuo
}

// RemoveRedemptionIDs removes the "redemptions" edge to Redemption entities by IDs.
func (uuo *UserUpdateOne) RemoveRedemptionIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveRedemptionIDs(ids...)
	return uuo
}

// RemoveRedemptions removes "redemptions" edges to Redemption entities.
func (uuo *UserUpdateOne) RemoveRedemptions(r ...*Redemption) *UserUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return uuo.RemoveRedemptionIDs(ids...)
}

// ClearPet clears all "pet" edges to the Pet entity.
func (uuo *UserUpdateOne) ClearPet() *UserUpdateOne {
	uuo.mutation.ClearPet()
	return uuo
}

// RemovePetIDs removes the "pet" edge to Pet entities by IDs.
func (uuo *UserUpdateOne) RemovePetIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemovePetIDs(ids...)
	return uuo
}

// RemovePet removes "pet" edges to Pet entities.
func (uuo *UserUpdateOne) RemovePet(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePetIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (uuo *UserUpdateOne) ClearGroups() *UserUpdateOne {
	uuo.mutation.ClearGroups()
	return uuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveGroupIDs(ids...)
	return uuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (uuo *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := user.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "User.email": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.PasswordHash(); ok {
		if err := user.PasswordHashValidator(v); err != nil {
			return &ValidationError{Name: "password_hash", err: fmt.Errorf(`ent: validator failed for field "User.password_hash": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Points(); ok {
		if err := user.PointsValidator(v); err != nil {
			return &ValidationError{Name: "points", err: fmt.Errorf(`ent: validator failed for field "User.points": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.PointsAwardedCount(); ok {
		if err := user.PointsAwardedCountValidator(v); err != nil {
			return &ValidationError{Name: "points_awarded_count", err: fmt.Errorf(`ent: validator failed for field "User.points_awarded_count": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.PasswordHash(); ok {
		_spec.SetField(user.FieldPasswordHash, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Points(); ok {
		_spec.SetField(user.FieldPoints, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedPoints(); ok {
		_spec.AddField(user.FieldPoints, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.PointsAwardedCount(); ok {
		_spec.SetField(user.FieldPointsAwardedCount, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedPointsAwardedCount(); ok {
		_spec.AddField(user.FieldPointsAwardedCount, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.PointsAwardedResetTime(); ok {
		_spec.SetField(user.FieldPointsAwardedResetTime, field.TypeTime, value)
	}
	if uuo.mutation.PointsAwardedResetTimeCleared() {
		_spec.ClearField(user.FieldPointsAwardedResetTime, field.TypeTime)
	}
	if value, ok := uuo.mutation.GodMode(); ok {
		_spec.SetField(user.FieldGodMode, field.TypeBool, value)
	}
	if uuo.mutation.InstitutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.InstitutionTable,
			Columns: user.InstitutionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedInstitutionIDs(); len(nodes) > 0 && !uuo.mutation.InstitutionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.InstitutionTable,
			Columns: user.InstitutionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.InstitutionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.InstitutionTable,
			Columns: user.InstitutionPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.CourseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CourseTable,
			Columns: []string{user.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CourseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CourseTable,
			Columns: []string{user.CourseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.RedemptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.RedemptionsTable,
			Columns: []string{user.RedemptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redemption.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedRedemptionsIDs(); len(nodes) > 0 && !uuo.mutation.RedemptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.RedemptionsTable,
			Columns: []string{user.RedemptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redemption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RedemptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.RedemptionsTable,
			Columns: []string{user.RedemptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(redemption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.PetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.PetTable,
			Columns: user.PetPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		createE := &UserPetCreate{config: uuo.config, mutation: newUserPetMutation(uuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedPetIDs(); len(nodes) > 0 && !uuo.mutation.PetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.PetTable,
			Columns: user.PetPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserPetCreate{config: uuo.config, mutation: newUserPetMutation(uuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.PetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.PetTable,
			Columns: user.PetPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserPetCreate{config: uuo.config, mutation: newUserPetMutation(uuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !uuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
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
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
