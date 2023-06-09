// Code generated by ent, DO NOT EDIT.

package jwtrevocation

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/np-inprove/server/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldLTE(FieldID, id))
}

// Jti applies equality check predicate on the "jti" field. It's identical to JtiEQ.
func Jti(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldEQ(FieldJti, v))
}

// Expiry applies equality check predicate on the "expiry" field. It's identical to ExpiryEQ.
func Expiry(v time.Time) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldEQ(FieldExpiry, v))
}

// JtiEQ applies the EQ predicate on the "jti" field.
func JtiEQ(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldEQ(FieldJti, v))
}

// JtiNEQ applies the NEQ predicate on the "jti" field.
func JtiNEQ(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldNEQ(FieldJti, v))
}

// JtiIn applies the In predicate on the "jti" field.
func JtiIn(vs ...string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldIn(FieldJti, vs...))
}

// JtiNotIn applies the NotIn predicate on the "jti" field.
func JtiNotIn(vs ...string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldNotIn(FieldJti, vs...))
}

// JtiGT applies the GT predicate on the "jti" field.
func JtiGT(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldGT(FieldJti, v))
}

// JtiGTE applies the GTE predicate on the "jti" field.
func JtiGTE(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldGTE(FieldJti, v))
}

// JtiLT applies the LT predicate on the "jti" field.
func JtiLT(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldLT(FieldJti, v))
}

// JtiLTE applies the LTE predicate on the "jti" field.
func JtiLTE(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldLTE(FieldJti, v))
}

// JtiContains applies the Contains predicate on the "jti" field.
func JtiContains(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldContains(FieldJti, v))
}

// JtiHasPrefix applies the HasPrefix predicate on the "jti" field.
func JtiHasPrefix(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldHasPrefix(FieldJti, v))
}

// JtiHasSuffix applies the HasSuffix predicate on the "jti" field.
func JtiHasSuffix(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldHasSuffix(FieldJti, v))
}

// JtiEqualFold applies the EqualFold predicate on the "jti" field.
func JtiEqualFold(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldEqualFold(FieldJti, v))
}

// JtiContainsFold applies the ContainsFold predicate on the "jti" field.
func JtiContainsFold(v string) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldContainsFold(FieldJti, v))
}

// ExpiryEQ applies the EQ predicate on the "expiry" field.
func ExpiryEQ(v time.Time) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldEQ(FieldExpiry, v))
}

// ExpiryNEQ applies the NEQ predicate on the "expiry" field.
func ExpiryNEQ(v time.Time) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldNEQ(FieldExpiry, v))
}

// ExpiryIn applies the In predicate on the "expiry" field.
func ExpiryIn(vs ...time.Time) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldIn(FieldExpiry, vs...))
}

// ExpiryNotIn applies the NotIn predicate on the "expiry" field.
func ExpiryNotIn(vs ...time.Time) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldNotIn(FieldExpiry, vs...))
}

// ExpiryGT applies the GT predicate on the "expiry" field.
func ExpiryGT(v time.Time) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldGT(FieldExpiry, v))
}

// ExpiryGTE applies the GTE predicate on the "expiry" field.
func ExpiryGTE(v time.Time) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldGTE(FieldExpiry, v))
}

// ExpiryLT applies the LT predicate on the "expiry" field.
func ExpiryLT(v time.Time) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldLT(FieldExpiry, v))
}

// ExpiryLTE applies the LTE predicate on the "expiry" field.
func ExpiryLTE(v time.Time) predicate.JWTRevocation {
	return predicate.JWTRevocation(sql.FieldLTE(FieldExpiry, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.JWTRevocation) predicate.JWTRevocation {
	return predicate.JWTRevocation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.JWTRevocation) predicate.JWTRevocation {
	return predicate.JWTRevocation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.JWTRevocation) predicate.JWTRevocation {
	return predicate.JWTRevocation(func(s *sql.Selector) {
		p(s.Not())
	})
}
