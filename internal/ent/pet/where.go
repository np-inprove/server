// Code generated by ent, DO NOT EDIT.

package pet

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/np-inprove/server/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Pet {
	return predicate.Pet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Pet {
	return predicate.Pet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Pet {
	return predicate.Pet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Pet {
	return predicate.Pet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Pet {
	return predicate.Pet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Pet {
	return predicate.Pet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Pet {
	return predicate.Pet(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldName, v))
}

// SvgRaw applies equality check predicate on the "svg_raw" field. It's identical to SvgRawEQ.
func SvgRaw(v string) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldSvgRaw, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Pet {
	return predicate.Pet(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Pet {
	return predicate.Pet(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Pet {
	return predicate.Pet(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Pet {
	return predicate.Pet(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Pet {
	return predicate.Pet(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Pet {
	return predicate.Pet(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Pet {
	return predicate.Pet(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Pet {
	return predicate.Pet(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Pet {
	return predicate.Pet(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Pet {
	return predicate.Pet(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Pet {
	return predicate.Pet(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Pet {
	return predicate.Pet(sql.FieldContainsFold(FieldName, v))
}

// SvgRawEQ applies the EQ predicate on the "svg_raw" field.
func SvgRawEQ(v string) predicate.Pet {
	return predicate.Pet(sql.FieldEQ(FieldSvgRaw, v))
}

// SvgRawNEQ applies the NEQ predicate on the "svg_raw" field.
func SvgRawNEQ(v string) predicate.Pet {
	return predicate.Pet(sql.FieldNEQ(FieldSvgRaw, v))
}

// SvgRawIn applies the In predicate on the "svg_raw" field.
func SvgRawIn(vs ...string) predicate.Pet {
	return predicate.Pet(sql.FieldIn(FieldSvgRaw, vs...))
}

// SvgRawNotIn applies the NotIn predicate on the "svg_raw" field.
func SvgRawNotIn(vs ...string) predicate.Pet {
	return predicate.Pet(sql.FieldNotIn(FieldSvgRaw, vs...))
}

// SvgRawGT applies the GT predicate on the "svg_raw" field.
func SvgRawGT(v string) predicate.Pet {
	return predicate.Pet(sql.FieldGT(FieldSvgRaw, v))
}

// SvgRawGTE applies the GTE predicate on the "svg_raw" field.
func SvgRawGTE(v string) predicate.Pet {
	return predicate.Pet(sql.FieldGTE(FieldSvgRaw, v))
}

// SvgRawLT applies the LT predicate on the "svg_raw" field.
func SvgRawLT(v string) predicate.Pet {
	return predicate.Pet(sql.FieldLT(FieldSvgRaw, v))
}

// SvgRawLTE applies the LTE predicate on the "svg_raw" field.
func SvgRawLTE(v string) predicate.Pet {
	return predicate.Pet(sql.FieldLTE(FieldSvgRaw, v))
}

// SvgRawContains applies the Contains predicate on the "svg_raw" field.
func SvgRawContains(v string) predicate.Pet {
	return predicate.Pet(sql.FieldContains(FieldSvgRaw, v))
}

// SvgRawHasPrefix applies the HasPrefix predicate on the "svg_raw" field.
func SvgRawHasPrefix(v string) predicate.Pet {
	return predicate.Pet(sql.FieldHasPrefix(FieldSvgRaw, v))
}

// SvgRawHasSuffix applies the HasSuffix predicate on the "svg_raw" field.
func SvgRawHasSuffix(v string) predicate.Pet {
	return predicate.Pet(sql.FieldHasSuffix(FieldSvgRaw, v))
}

// SvgRawEqualFold applies the EqualFold predicate on the "svg_raw" field.
func SvgRawEqualFold(v string) predicate.Pet {
	return predicate.Pet(sql.FieldEqualFold(FieldSvgRaw, v))
}

// SvgRawContainsFold applies the ContainsFold predicate on the "svg_raw" field.
func SvgRawContainsFold(v string) predicate.Pet {
	return predicate.Pet(sql.FieldContainsFold(FieldSvgRaw, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, OwnerTable, OwnerPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserPets applies the HasEdge predicate on the "user_pets" edge.
func HasUserPets() predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserPetsTable, UserPetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserPetsWith applies the HasEdge predicate on the "user_pets" edge with a given conditions (other predicates).
func HasUserPetsWith(preds ...predicate.UserPet) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := newUserPetsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
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
func Not(p predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		p(s.Not())
	})
}
