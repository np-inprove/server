// Code generated by ent, DO NOT EDIT.

package institutioninvitelink

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/np-inprove/server/internal/ent/predicate"
	"github.com/np-inprove/server/internal/entity/institution"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldLTE(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldEQ(FieldCode, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(sql.FieldContainsFold(FieldCode, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v institution.Role) predicate.InstitutionInviteLink {
	vc := v
	return predicate.InstitutionInviteLink(sql.FieldEQ(FieldRole, vc))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v institution.Role) predicate.InstitutionInviteLink {
	vc := v
	return predicate.InstitutionInviteLink(sql.FieldNEQ(FieldRole, vc))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...institution.Role) predicate.InstitutionInviteLink {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InstitutionInviteLink(sql.FieldIn(FieldRole, v...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...institution.Role) predicate.InstitutionInviteLink {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.InstitutionInviteLink(sql.FieldNotIn(FieldRole, v...))
}

// HasInstitution applies the HasEdge predicate on the "institution" edge.
func HasInstitution() predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, InstitutionTable, InstitutionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstitutionWith applies the HasEdge predicate on the "institution" edge with a given conditions (other predicates).
func HasInstitutionWith(preds ...predicate.Institution) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(func(s *sql.Selector) {
		step := newInstitutionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.InstitutionInviteLink) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.InstitutionInviteLink) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(func(s *sql.Selector) {
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
func Not(p predicate.InstitutionInviteLink) predicate.InstitutionInviteLink {
	return predicate.InstitutionInviteLink(func(s *sql.Selector) {
		p(s.Not())
	})
}
