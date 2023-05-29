// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/np-inprove/server/internal/ent/accessory"
	"github.com/np-inprove/server/internal/ent/department"
	"github.com/np-inprove/server/internal/ent/group"
	"github.com/np-inprove/server/internal/ent/institution"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/ent/voucher"
)

// InstitutionUpdate is the builder for updating Institution entities.
type InstitutionUpdate struct {
	config
	hooks    []Hook
	mutation *InstitutionMutation
}

// Where appends a list predicates to the InstitutionUpdate builder.
func (iu *InstitutionUpdate) Where(ps ...predicate.Institution) *InstitutionUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetName sets the "name" field.
func (iu *InstitutionUpdate) SetName(s string) *InstitutionUpdate {
	iu.mutation.SetName(s)
	return iu
}

// SetShortName sets the "short_name" field.
func (iu *InstitutionUpdate) SetShortName(s string) *InstitutionUpdate {
	iu.mutation.SetShortName(s)
	return iu
}

// SetAdminDomain sets the "admin_domain" field.
func (iu *InstitutionUpdate) SetAdminDomain(s string) *InstitutionUpdate {
	iu.mutation.SetAdminDomain(s)
	return iu
}

// SetStudentDomain sets the "student_domain" field.
func (iu *InstitutionUpdate) SetStudentDomain(s string) *InstitutionUpdate {
	iu.mutation.SetStudentDomain(s)
	return iu
}

// AddVoucherIDs adds the "vouchers" edge to the Voucher entity by IDs.
func (iu *InstitutionUpdate) AddVoucherIDs(ids ...int) *InstitutionUpdate {
	iu.mutation.AddVoucherIDs(ids...)
	return iu
}

// AddVouchers adds the "vouchers" edges to the Voucher entity.
func (iu *InstitutionUpdate) AddVouchers(v ...*Voucher) *InstitutionUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return iu.AddVoucherIDs(ids...)
}

// AddAccessoryIDs adds the "accessories" edge to the Accessory entity by IDs.
func (iu *InstitutionUpdate) AddAccessoryIDs(ids ...int) *InstitutionUpdate {
	iu.mutation.AddAccessoryIDs(ids...)
	return iu
}

// AddAccessories adds the "accessories" edges to the Accessory entity.
func (iu *InstitutionUpdate) AddAccessories(a ...*Accessory) *InstitutionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iu.AddAccessoryIDs(ids...)
}

// AddDepartmentIDs adds the "departments" edge to the Department entity by IDs.
func (iu *InstitutionUpdate) AddDepartmentIDs(ids ...int) *InstitutionUpdate {
	iu.mutation.AddDepartmentIDs(ids...)
	return iu
}

// AddDepartments adds the "departments" edges to the Department entity.
func (iu *InstitutionUpdate) AddDepartments(d ...*Department) *InstitutionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iu.AddDepartmentIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (iu *InstitutionUpdate) AddGroupIDs(ids ...int) *InstitutionUpdate {
	iu.mutation.AddGroupIDs(ids...)
	return iu
}

// AddGroups adds the "groups" edges to the Group entity.
func (iu *InstitutionUpdate) AddGroups(g ...*Group) *InstitutionUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return iu.AddGroupIDs(ids...)
}

// Mutation returns the InstitutionMutation object of the builder.
func (iu *InstitutionUpdate) Mutation() *InstitutionMutation {
	return iu.mutation
}

// ClearVouchers clears all "vouchers" edges to the Voucher entity.
func (iu *InstitutionUpdate) ClearVouchers() *InstitutionUpdate {
	iu.mutation.ClearVouchers()
	return iu
}

// RemoveVoucherIDs removes the "vouchers" edge to Voucher entities by IDs.
func (iu *InstitutionUpdate) RemoveVoucherIDs(ids ...int) *InstitutionUpdate {
	iu.mutation.RemoveVoucherIDs(ids...)
	return iu
}

// RemoveVouchers removes "vouchers" edges to Voucher entities.
func (iu *InstitutionUpdate) RemoveVouchers(v ...*Voucher) *InstitutionUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return iu.RemoveVoucherIDs(ids...)
}

// ClearAccessories clears all "accessories" edges to the Accessory entity.
func (iu *InstitutionUpdate) ClearAccessories() *InstitutionUpdate {
	iu.mutation.ClearAccessories()
	return iu
}

// RemoveAccessoryIDs removes the "accessories" edge to Accessory entities by IDs.
func (iu *InstitutionUpdate) RemoveAccessoryIDs(ids ...int) *InstitutionUpdate {
	iu.mutation.RemoveAccessoryIDs(ids...)
	return iu
}

// RemoveAccessories removes "accessories" edges to Accessory entities.
func (iu *InstitutionUpdate) RemoveAccessories(a ...*Accessory) *InstitutionUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iu.RemoveAccessoryIDs(ids...)
}

// ClearDepartments clears all "departments" edges to the Department entity.
func (iu *InstitutionUpdate) ClearDepartments() *InstitutionUpdate {
	iu.mutation.ClearDepartments()
	return iu
}

// RemoveDepartmentIDs removes the "departments" edge to Department entities by IDs.
func (iu *InstitutionUpdate) RemoveDepartmentIDs(ids ...int) *InstitutionUpdate {
	iu.mutation.RemoveDepartmentIDs(ids...)
	return iu
}

// RemoveDepartments removes "departments" edges to Department entities.
func (iu *InstitutionUpdate) RemoveDepartments(d ...*Department) *InstitutionUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iu.RemoveDepartmentIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (iu *InstitutionUpdate) ClearGroups() *InstitutionUpdate {
	iu.mutation.ClearGroups()
	return iu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (iu *InstitutionUpdate) RemoveGroupIDs(ids ...int) *InstitutionUpdate {
	iu.mutation.RemoveGroupIDs(ids...)
	return iu
}

// RemoveGroups removes "groups" edges to Group entities.
func (iu *InstitutionUpdate) RemoveGroups(g ...*Group) *InstitutionUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return iu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InstitutionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InstitutionUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InstitutionUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InstitutionUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *InstitutionUpdate) check() error {
	if v, ok := iu.mutation.Name(); ok {
		if err := institution.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Institution.name": %w`, err)}
		}
	}
	if v, ok := iu.mutation.ShortName(); ok {
		if err := institution.ShortNameValidator(v); err != nil {
			return &ValidationError{Name: "short_name", err: fmt.Errorf(`ent: validator failed for field "Institution.short_name": %w`, err)}
		}
	}
	if v, ok := iu.mutation.AdminDomain(); ok {
		if err := institution.AdminDomainValidator(v); err != nil {
			return &ValidationError{Name: "admin_domain", err: fmt.Errorf(`ent: validator failed for field "Institution.admin_domain": %w`, err)}
		}
	}
	if v, ok := iu.mutation.StudentDomain(); ok {
		if err := institution.StudentDomainValidator(v); err != nil {
			return &ValidationError{Name: "student_domain", err: fmt.Errorf(`ent: validator failed for field "Institution.student_domain": %w`, err)}
		}
	}
	return nil
}

func (iu *InstitutionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(institution.Table, institution.Columns, sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Name(); ok {
		_spec.SetField(institution.FieldName, field.TypeString, value)
	}
	if value, ok := iu.mutation.ShortName(); ok {
		_spec.SetField(institution.FieldShortName, field.TypeString, value)
	}
	if value, ok := iu.mutation.AdminDomain(); ok {
		_spec.SetField(institution.FieldAdminDomain, field.TypeString, value)
	}
	if value, ok := iu.mutation.StudentDomain(); ok {
		_spec.SetField(institution.FieldStudentDomain, field.TypeString, value)
	}
	if iu.mutation.VouchersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.VouchersTable,
			Columns: []string{institution.VouchersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(voucher.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedVouchersIDs(); len(nodes) > 0 && !iu.mutation.VouchersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.VouchersTable,
			Columns: []string{institution.VouchersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(voucher.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.VouchersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.VouchersTable,
			Columns: []string{institution.VouchersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(voucher.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.AccessoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.AccessoriesTable,
			Columns: []string{institution.AccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accessory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedAccessoriesIDs(); len(nodes) > 0 && !iu.mutation.AccessoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.AccessoriesTable,
			Columns: []string{institution.AccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accessory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.AccessoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.AccessoriesTable,
			Columns: []string{institution.AccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accessory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.DepartmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.DepartmentsTable,
			Columns: []string{institution.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedDepartmentsIDs(); len(nodes) > 0 && !iu.mutation.DepartmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.DepartmentsTable,
			Columns: []string{institution.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.DepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.DepartmentsTable,
			Columns: []string{institution.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.GroupsTable,
			Columns: []string{institution.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !iu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.GroupsTable,
			Columns: []string{institution.GroupsColumn},
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
	if nodes := iu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.GroupsTable,
			Columns: []string{institution.GroupsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{institution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// InstitutionUpdateOne is the builder for updating a single Institution entity.
type InstitutionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InstitutionMutation
}

// SetName sets the "name" field.
func (iuo *InstitutionUpdateOne) SetName(s string) *InstitutionUpdateOne {
	iuo.mutation.SetName(s)
	return iuo
}

// SetShortName sets the "short_name" field.
func (iuo *InstitutionUpdateOne) SetShortName(s string) *InstitutionUpdateOne {
	iuo.mutation.SetShortName(s)
	return iuo
}

// SetAdminDomain sets the "admin_domain" field.
func (iuo *InstitutionUpdateOne) SetAdminDomain(s string) *InstitutionUpdateOne {
	iuo.mutation.SetAdminDomain(s)
	return iuo
}

// SetStudentDomain sets the "student_domain" field.
func (iuo *InstitutionUpdateOne) SetStudentDomain(s string) *InstitutionUpdateOne {
	iuo.mutation.SetStudentDomain(s)
	return iuo
}

// AddVoucherIDs adds the "vouchers" edge to the Voucher entity by IDs.
func (iuo *InstitutionUpdateOne) AddVoucherIDs(ids ...int) *InstitutionUpdateOne {
	iuo.mutation.AddVoucherIDs(ids...)
	return iuo
}

// AddVouchers adds the "vouchers" edges to the Voucher entity.
func (iuo *InstitutionUpdateOne) AddVouchers(v ...*Voucher) *InstitutionUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return iuo.AddVoucherIDs(ids...)
}

// AddAccessoryIDs adds the "accessories" edge to the Accessory entity by IDs.
func (iuo *InstitutionUpdateOne) AddAccessoryIDs(ids ...int) *InstitutionUpdateOne {
	iuo.mutation.AddAccessoryIDs(ids...)
	return iuo
}

// AddAccessories adds the "accessories" edges to the Accessory entity.
func (iuo *InstitutionUpdateOne) AddAccessories(a ...*Accessory) *InstitutionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iuo.AddAccessoryIDs(ids...)
}

// AddDepartmentIDs adds the "departments" edge to the Department entity by IDs.
func (iuo *InstitutionUpdateOne) AddDepartmentIDs(ids ...int) *InstitutionUpdateOne {
	iuo.mutation.AddDepartmentIDs(ids...)
	return iuo
}

// AddDepartments adds the "departments" edges to the Department entity.
func (iuo *InstitutionUpdateOne) AddDepartments(d ...*Department) *InstitutionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iuo.AddDepartmentIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (iuo *InstitutionUpdateOne) AddGroupIDs(ids ...int) *InstitutionUpdateOne {
	iuo.mutation.AddGroupIDs(ids...)
	return iuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (iuo *InstitutionUpdateOne) AddGroups(g ...*Group) *InstitutionUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return iuo.AddGroupIDs(ids...)
}

// Mutation returns the InstitutionMutation object of the builder.
func (iuo *InstitutionUpdateOne) Mutation() *InstitutionMutation {
	return iuo.mutation
}

// ClearVouchers clears all "vouchers" edges to the Voucher entity.
func (iuo *InstitutionUpdateOne) ClearVouchers() *InstitutionUpdateOne {
	iuo.mutation.ClearVouchers()
	return iuo
}

// RemoveVoucherIDs removes the "vouchers" edge to Voucher entities by IDs.
func (iuo *InstitutionUpdateOne) RemoveVoucherIDs(ids ...int) *InstitutionUpdateOne {
	iuo.mutation.RemoveVoucherIDs(ids...)
	return iuo
}

// RemoveVouchers removes "vouchers" edges to Voucher entities.
func (iuo *InstitutionUpdateOne) RemoveVouchers(v ...*Voucher) *InstitutionUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return iuo.RemoveVoucherIDs(ids...)
}

// ClearAccessories clears all "accessories" edges to the Accessory entity.
func (iuo *InstitutionUpdateOne) ClearAccessories() *InstitutionUpdateOne {
	iuo.mutation.ClearAccessories()
	return iuo
}

// RemoveAccessoryIDs removes the "accessories" edge to Accessory entities by IDs.
func (iuo *InstitutionUpdateOne) RemoveAccessoryIDs(ids ...int) *InstitutionUpdateOne {
	iuo.mutation.RemoveAccessoryIDs(ids...)
	return iuo
}

// RemoveAccessories removes "accessories" edges to Accessory entities.
func (iuo *InstitutionUpdateOne) RemoveAccessories(a ...*Accessory) *InstitutionUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iuo.RemoveAccessoryIDs(ids...)
}

// ClearDepartments clears all "departments" edges to the Department entity.
func (iuo *InstitutionUpdateOne) ClearDepartments() *InstitutionUpdateOne {
	iuo.mutation.ClearDepartments()
	return iuo
}

// RemoveDepartmentIDs removes the "departments" edge to Department entities by IDs.
func (iuo *InstitutionUpdateOne) RemoveDepartmentIDs(ids ...int) *InstitutionUpdateOne {
	iuo.mutation.RemoveDepartmentIDs(ids...)
	return iuo
}

// RemoveDepartments removes "departments" edges to Department entities.
func (iuo *InstitutionUpdateOne) RemoveDepartments(d ...*Department) *InstitutionUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return iuo.RemoveDepartmentIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (iuo *InstitutionUpdateOne) ClearGroups() *InstitutionUpdateOne {
	iuo.mutation.ClearGroups()
	return iuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (iuo *InstitutionUpdateOne) RemoveGroupIDs(ids ...int) *InstitutionUpdateOne {
	iuo.mutation.RemoveGroupIDs(ids...)
	return iuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (iuo *InstitutionUpdateOne) RemoveGroups(g ...*Group) *InstitutionUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return iuo.RemoveGroupIDs(ids...)
}

// Where appends a list predicates to the InstitutionUpdate builder.
func (iuo *InstitutionUpdateOne) Where(ps ...predicate.Institution) *InstitutionUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InstitutionUpdateOne) Select(field string, fields ...string) *InstitutionUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Institution entity.
func (iuo *InstitutionUpdateOne) Save(ctx context.Context) (*Institution, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InstitutionUpdateOne) SaveX(ctx context.Context) *Institution {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InstitutionUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InstitutionUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *InstitutionUpdateOne) check() error {
	if v, ok := iuo.mutation.Name(); ok {
		if err := institution.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Institution.name": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.ShortName(); ok {
		if err := institution.ShortNameValidator(v); err != nil {
			return &ValidationError{Name: "short_name", err: fmt.Errorf(`ent: validator failed for field "Institution.short_name": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.AdminDomain(); ok {
		if err := institution.AdminDomainValidator(v); err != nil {
			return &ValidationError{Name: "admin_domain", err: fmt.Errorf(`ent: validator failed for field "Institution.admin_domain": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.StudentDomain(); ok {
		if err := institution.StudentDomainValidator(v); err != nil {
			return &ValidationError{Name: "student_domain", err: fmt.Errorf(`ent: validator failed for field "Institution.student_domain": %w`, err)}
		}
	}
	return nil
}

func (iuo *InstitutionUpdateOne) sqlSave(ctx context.Context) (_node *Institution, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(institution.Table, institution.Columns, sqlgraph.NewFieldSpec(institution.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Institution.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, institution.FieldID)
		for _, f := range fields {
			if !institution.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != institution.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.Name(); ok {
		_spec.SetField(institution.FieldName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.ShortName(); ok {
		_spec.SetField(institution.FieldShortName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.AdminDomain(); ok {
		_spec.SetField(institution.FieldAdminDomain, field.TypeString, value)
	}
	if value, ok := iuo.mutation.StudentDomain(); ok {
		_spec.SetField(institution.FieldStudentDomain, field.TypeString, value)
	}
	if iuo.mutation.VouchersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.VouchersTable,
			Columns: []string{institution.VouchersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(voucher.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedVouchersIDs(); len(nodes) > 0 && !iuo.mutation.VouchersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.VouchersTable,
			Columns: []string{institution.VouchersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(voucher.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.VouchersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.VouchersTable,
			Columns: []string{institution.VouchersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(voucher.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.AccessoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.AccessoriesTable,
			Columns: []string{institution.AccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accessory.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedAccessoriesIDs(); len(nodes) > 0 && !iuo.mutation.AccessoriesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.AccessoriesTable,
			Columns: []string{institution.AccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accessory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.AccessoriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.AccessoriesTable,
			Columns: []string{institution.AccessoriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(accessory.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.DepartmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.DepartmentsTable,
			Columns: []string{institution.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedDepartmentsIDs(); len(nodes) > 0 && !iuo.mutation.DepartmentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.DepartmentsTable,
			Columns: []string{institution.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.DepartmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.DepartmentsTable,
			Columns: []string{institution.DepartmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(department.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.GroupsTable,
			Columns: []string{institution.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !iuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.GroupsTable,
			Columns: []string{institution.GroupsColumn},
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
	if nodes := iuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   institution.GroupsTable,
			Columns: []string{institution.GroupsColumn},
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
	_node = &Institution{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{institution.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
