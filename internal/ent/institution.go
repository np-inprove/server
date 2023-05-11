// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/np-inprove/server/internal/ent/institution"
)

// Institution is the model entity for the Institution schema.
type Institution struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name of the institution
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InstitutionQuery when eager-loading is set.
	Edges        InstitutionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// InstitutionEdges holds the relations/edges for other nodes in the graph.
type InstitutionEdges struct {
	// Admins of the institution
	Admins []*User `json:"admins,omitempty"`
	// Prizes (vouchers) available to be redeemed by users of the institution
	Vouchers []*Voucher `json:"vouchers,omitempty"`
	// Prizes (accessories) available to be redeemed by users of the institution
	Accessories []*Accessory `json:"accessories,omitempty"`
	// Academic schools of the institution
	AcademicSchools []*AcademicSchool `json:"academic_schools,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// AdminsOrErr returns the Admins value or an error if the edge
// was not loaded in eager-loading.
func (e InstitutionEdges) AdminsOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Admins, nil
	}
	return nil, &NotLoadedError{edge: "admins"}
}

// VouchersOrErr returns the Vouchers value or an error if the edge
// was not loaded in eager-loading.
func (e InstitutionEdges) VouchersOrErr() ([]*Voucher, error) {
	if e.loadedTypes[1] {
		return e.Vouchers, nil
	}
	return nil, &NotLoadedError{edge: "vouchers"}
}

// AccessoriesOrErr returns the Accessories value or an error if the edge
// was not loaded in eager-loading.
func (e InstitutionEdges) AccessoriesOrErr() ([]*Accessory, error) {
	if e.loadedTypes[2] {
		return e.Accessories, nil
	}
	return nil, &NotLoadedError{edge: "accessories"}
}

// AcademicSchoolsOrErr returns the AcademicSchools value or an error if the edge
// was not loaded in eager-loading.
func (e InstitutionEdges) AcademicSchoolsOrErr() ([]*AcademicSchool, error) {
	if e.loadedTypes[3] {
		return e.AcademicSchools, nil
	}
	return nil, &NotLoadedError{edge: "academic_schools"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Institution) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case institution.FieldID:
			values[i] = new(sql.NullInt64)
		case institution.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Institution fields.
func (i *Institution) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case institution.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case institution.FieldName:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[j])
			} else if value.Valid {
				i.Name = value.String
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Institution.
// This includes values selected through modifiers, order, etc.
func (i *Institution) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryAdmins queries the "admins" edge of the Institution entity.
func (i *Institution) QueryAdmins() *UserQuery {
	return NewInstitutionClient(i.config).QueryAdmins(i)
}

// QueryVouchers queries the "vouchers" edge of the Institution entity.
func (i *Institution) QueryVouchers() *VoucherQuery {
	return NewInstitutionClient(i.config).QueryVouchers(i)
}

// QueryAccessories queries the "accessories" edge of the Institution entity.
func (i *Institution) QueryAccessories() *AccessoryQuery {
	return NewInstitutionClient(i.config).QueryAccessories(i)
}

// QueryAcademicSchools queries the "academic_schools" edge of the Institution entity.
func (i *Institution) QueryAcademicSchools() *AcademicSchoolQuery {
	return NewInstitutionClient(i.config).QueryAcademicSchools(i)
}

// Update returns a builder for updating this Institution.
// Note that you need to call Institution.Unwrap() before calling this method if this Institution
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Institution) Update() *InstitutionUpdateOne {
	return NewInstitutionClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Institution entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Institution) Unwrap() *Institution {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Institution is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Institution) String() string {
	var builder strings.Builder
	builder.WriteString("Institution(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("name=")
	builder.WriteString(i.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Institutions is a parsable slice of Institution.
type Institutions []*Institution
