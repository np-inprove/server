// Code generated by ent, DO NOT EDIT.

package studyplan

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/np-inprove/server/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldEQ(FieldName, v))
}

// ShareCode applies equality check predicate on the "share_code" field. It's identical to ShareCodeEQ.
func ShareCode(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldEQ(FieldShareCode, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldContainsFold(FieldName, v))
}

// ShareCodeEQ applies the EQ predicate on the "share_code" field.
func ShareCodeEQ(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldEQ(FieldShareCode, v))
}

// ShareCodeNEQ applies the NEQ predicate on the "share_code" field.
func ShareCodeNEQ(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldNEQ(FieldShareCode, v))
}

// ShareCodeIn applies the In predicate on the "share_code" field.
func ShareCodeIn(vs ...string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldIn(FieldShareCode, vs...))
}

// ShareCodeNotIn applies the NotIn predicate on the "share_code" field.
func ShareCodeNotIn(vs ...string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldNotIn(FieldShareCode, vs...))
}

// ShareCodeGT applies the GT predicate on the "share_code" field.
func ShareCodeGT(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldGT(FieldShareCode, v))
}

// ShareCodeGTE applies the GTE predicate on the "share_code" field.
func ShareCodeGTE(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldGTE(FieldShareCode, v))
}

// ShareCodeLT applies the LT predicate on the "share_code" field.
func ShareCodeLT(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldLT(FieldShareCode, v))
}

// ShareCodeLTE applies the LTE predicate on the "share_code" field.
func ShareCodeLTE(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldLTE(FieldShareCode, v))
}

// ShareCodeContains applies the Contains predicate on the "share_code" field.
func ShareCodeContains(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldContains(FieldShareCode, v))
}

// ShareCodeHasPrefix applies the HasPrefix predicate on the "share_code" field.
func ShareCodeHasPrefix(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldHasPrefix(FieldShareCode, v))
}

// ShareCodeHasSuffix applies the HasSuffix predicate on the "share_code" field.
func ShareCodeHasSuffix(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldHasSuffix(FieldShareCode, v))
}

// ShareCodeIsNil applies the IsNil predicate on the "share_code" field.
func ShareCodeIsNil() predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldIsNull(FieldShareCode))
}

// ShareCodeNotNil applies the NotNil predicate on the "share_code" field.
func ShareCodeNotNil() predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldNotNull(FieldShareCode))
}

// ShareCodeEqualFold applies the EqualFold predicate on the "share_code" field.
func ShareCodeEqualFold(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldEqualFold(FieldShareCode, v))
}

// ShareCodeContainsFold applies the ContainsFold predicate on the "share_code" field.
func ShareCodeContainsFold(v string) predicate.StudyPlan {
	return predicate.StudyPlan(sql.FieldContainsFold(FieldShareCode, v))
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.StudyPlan {
	return predicate.StudyPlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.StudyPlan {
	return predicate.StudyPlan(func(s *sql.Selector) {
		step := newAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMilestones applies the HasEdge predicate on the "milestones" edge.
func HasMilestones() predicate.StudyPlan {
	return predicate.StudyPlan(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MilestonesTable, MilestonesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMilestonesWith applies the HasEdge predicate on the "milestones" edge with a given conditions (other predicates).
func HasMilestonesWith(preds ...predicate.Milestone) predicate.StudyPlan {
	return predicate.StudyPlan(func(s *sql.Selector) {
		step := newMilestonesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.StudyPlan) predicate.StudyPlan {
	return predicate.StudyPlan(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.StudyPlan) predicate.StudyPlan {
	return predicate.StudyPlan(func(s *sql.Selector) {
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
func Not(p predicate.StudyPlan) predicate.StudyPlan {
	return predicate.StudyPlan(func(s *sql.Selector) {
		p(s.Not())
	})
}
