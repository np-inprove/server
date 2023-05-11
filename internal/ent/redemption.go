// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/np-inprove/server/internal/ent/accessory"
	"github.com/np-inprove/server/internal/ent/redemption"
	"github.com/np-inprove/server/internal/ent/user"
	"github.com/np-inprove/server/internal/ent/voucher"
)

// Redemption is the model entity for the Redemption schema.
type Redemption struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Time when the prize was redeemed
	RedeemedAt time.Time `json:"redeemed_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RedemptionQuery when eager-loading is set.
	Edges                 RedemptionEdges `json:"edges"`
	accessory_redemptions *int
	redemption_user       *int
	voucher_redemptions   *int
	selectValues          sql.SelectValues
}

// RedemptionEdges holds the relations/edges for other nodes in the graph.
type RedemptionEdges struct {
	// Voucher holds the value of the voucher edge.
	Voucher *Voucher `json:"voucher,omitempty"`
	// Accessory holds the value of the accessory edge.
	Accessory *Accessory `json:"accessory,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// VoucherOrErr returns the Voucher value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RedemptionEdges) VoucherOrErr() (*Voucher, error) {
	if e.loadedTypes[0] {
		if e.Voucher == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: voucher.Label}
		}
		return e.Voucher, nil
	}
	return nil, &NotLoadedError{edge: "voucher"}
}

// AccessoryOrErr returns the Accessory value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RedemptionEdges) AccessoryOrErr() (*Accessory, error) {
	if e.loadedTypes[1] {
		if e.Accessory == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: accessory.Label}
		}
		return e.Accessory, nil
	}
	return nil, &NotLoadedError{edge: "accessory"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RedemptionEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[2] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Redemption) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case redemption.FieldID:
			values[i] = new(sql.NullInt64)
		case redemption.FieldRedeemedAt:
			values[i] = new(sql.NullTime)
		case redemption.ForeignKeys[0]: // accessory_redemptions
			values[i] = new(sql.NullInt64)
		case redemption.ForeignKeys[1]: // redemption_user
			values[i] = new(sql.NullInt64)
		case redemption.ForeignKeys[2]: // voucher_redemptions
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Redemption fields.
func (r *Redemption) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case redemption.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case redemption.FieldRedeemedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field redeemed_at", values[i])
			} else if value.Valid {
				r.RedeemedAt = value.Time
			}
		case redemption.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field accessory_redemptions", value)
			} else if value.Valid {
				r.accessory_redemptions = new(int)
				*r.accessory_redemptions = int(value.Int64)
			}
		case redemption.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field redemption_user", value)
			} else if value.Valid {
				r.redemption_user = new(int)
				*r.redemption_user = int(value.Int64)
			}
		case redemption.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field voucher_redemptions", value)
			} else if value.Valid {
				r.voucher_redemptions = new(int)
				*r.voucher_redemptions = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Redemption.
// This includes values selected through modifiers, order, etc.
func (r *Redemption) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryVoucher queries the "voucher" edge of the Redemption entity.
func (r *Redemption) QueryVoucher() *VoucherQuery {
	return NewRedemptionClient(r.config).QueryVoucher(r)
}

// QueryAccessory queries the "accessory" edge of the Redemption entity.
func (r *Redemption) QueryAccessory() *AccessoryQuery {
	return NewRedemptionClient(r.config).QueryAccessory(r)
}

// QueryUser queries the "user" edge of the Redemption entity.
func (r *Redemption) QueryUser() *UserQuery {
	return NewRedemptionClient(r.config).QueryUser(r)
}

// Update returns a builder for updating this Redemption.
// Note that you need to call Redemption.Unwrap() before calling this method if this Redemption
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Redemption) Update() *RedemptionUpdateOne {
	return NewRedemptionClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Redemption entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Redemption) Unwrap() *Redemption {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Redemption is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Redemption) String() string {
	var builder strings.Builder
	builder.WriteString("Redemption(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("redeemed_at=")
	builder.WriteString(r.RedeemedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Redemptions is a parsable slice of Redemption.
type Redemptions []*Redemption
