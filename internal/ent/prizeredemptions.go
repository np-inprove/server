// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/np-inprove/server/internal/ent/prize"
	"github.com/np-inprove/server/internal/ent/prizeredemptions"
	"github.com/np-inprove/server/internal/ent/user"
)

// PrizeRedemptions is the model entity for the PrizeRedemptions schema.
type PrizeRedemptions struct {
	config `json:"-"`
	// Time when the prize was redeemed
	RedeemedAt time.Time `json:"redeemed_at,omitempty"`
	// ID of the prize
	PrizeID int `json:"prize_id,omitempty"`
	// ID of the user
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PrizeRedemptionsQuery when eager-loading is set.
	Edges        PrizeRedemptionsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PrizeRedemptionsEdges holds the relations/edges for other nodes in the graph.
type PrizeRedemptionsEdges struct {
	// Prize holds the value of the prize edge.
	Prize *Prize `json:"prize,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PrizeOrErr returns the Prize value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrizeRedemptionsEdges) PrizeOrErr() (*Prize, error) {
	if e.loadedTypes[0] {
		if e.Prize == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: prize.Label}
		}
		return e.Prize, nil
	}
	return nil, &NotLoadedError{edge: "prize"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrizeRedemptionsEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PrizeRedemptions) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case prizeredemptions.FieldPrizeID, prizeredemptions.FieldUserID:
			values[i] = new(sql.NullInt64)
		case prizeredemptions.FieldRedeemedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PrizeRedemptions fields.
func (pr *PrizeRedemptions) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case prizeredemptions.FieldRedeemedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field redeemed_at", values[i])
			} else if value.Valid {
				pr.RedeemedAt = value.Time
			}
		case prizeredemptions.FieldPrizeID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field prize_id", values[i])
			} else if value.Valid {
				pr.PrizeID = int(value.Int64)
			}
		case prizeredemptions.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				pr.UserID = int(value.Int64)
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PrizeRedemptions.
// This includes values selected through modifiers, order, etc.
func (pr *PrizeRedemptions) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryPrize queries the "prize" edge of the PrizeRedemptions entity.
func (pr *PrizeRedemptions) QueryPrize() *PrizeQuery {
	return NewPrizeRedemptionsClient(pr.config).QueryPrize(pr)
}

// QueryUser queries the "user" edge of the PrizeRedemptions entity.
func (pr *PrizeRedemptions) QueryUser() *UserQuery {
	return NewPrizeRedemptionsClient(pr.config).QueryUser(pr)
}

// Update returns a builder for updating this PrizeRedemptions.
// Note that you need to call PrizeRedemptions.Unwrap() before calling this method if this PrizeRedemptions
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *PrizeRedemptions) Update() *PrizeRedemptionsUpdateOne {
	return NewPrizeRedemptionsClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the PrizeRedemptions entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *PrizeRedemptions) Unwrap() *PrizeRedemptions {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: PrizeRedemptions is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *PrizeRedemptions) String() string {
	var builder strings.Builder
	builder.WriteString("PrizeRedemptions(")
	builder.WriteString("redeemed_at=")
	builder.WriteString(pr.RedeemedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("prize_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.PrizeID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", pr.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// PrizeRedemptionsSlice is a parsable slice of PrizeRedemptions.
type PrizeRedemptionsSlice []*PrizeRedemptions
