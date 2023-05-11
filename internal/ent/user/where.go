// Code generated by ent, DO NOT EDIT.

package user

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/np-inprove/server/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// FirstName applies equality check predicate on the "first_name" field. It's identical to FirstNameEQ.
func FirstName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// PasswordHash applies equality check predicate on the "password_hash" field. It's identical to PasswordHashEQ.
func PasswordHash(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordHash, v))
}

// Points applies equality check predicate on the "points" field. It's identical to PointsEQ.
func Points(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPoints, v))
}

// PointsAwardedCount applies equality check predicate on the "points_awarded_count" field. It's identical to PointsAwardedCountEQ.
func PointsAwardedCount(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPointsAwardedCount, v))
}

// PointsAwardedResetTime applies equality check predicate on the "points_awarded_reset_time" field. It's identical to PointsAwardedResetTimeEQ.
func PointsAwardedResetTime(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPointsAwardedResetTime, v))
}

// GodMode applies equality check predicate on the "god_mode" field. It's identical to GodModeEQ.
func GodMode(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGodMode, v))
}

// FirstNameEQ applies the EQ predicate on the "first_name" field.
func FirstNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFirstName, v))
}

// FirstNameNEQ applies the NEQ predicate on the "first_name" field.
func FirstNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFirstName, v))
}

// FirstNameIn applies the In predicate on the "first_name" field.
func FirstNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldFirstName, vs...))
}

// FirstNameNotIn applies the NotIn predicate on the "first_name" field.
func FirstNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFirstName, vs...))
}

// FirstNameGT applies the GT predicate on the "first_name" field.
func FirstNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldFirstName, v))
}

// FirstNameGTE applies the GTE predicate on the "first_name" field.
func FirstNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFirstName, v))
}

// FirstNameLT applies the LT predicate on the "first_name" field.
func FirstNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldFirstName, v))
}

// FirstNameLTE applies the LTE predicate on the "first_name" field.
func FirstNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFirstName, v))
}

// FirstNameContains applies the Contains predicate on the "first_name" field.
func FirstNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldFirstName, v))
}

// FirstNameHasPrefix applies the HasPrefix predicate on the "first_name" field.
func FirstNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldFirstName, v))
}

// FirstNameHasSuffix applies the HasSuffix predicate on the "first_name" field.
func FirstNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldFirstName, v))
}

// FirstNameEqualFold applies the EqualFold predicate on the "first_name" field.
func FirstNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldFirstName, v))
}

// FirstNameContainsFold applies the ContainsFold predicate on the "first_name" field.
func FirstNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldFirstName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldLastName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldEmail, v))
}

// PasswordHashEQ applies the EQ predicate on the "password_hash" field.
func PasswordHashEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswordHash, v))
}

// PasswordHashNEQ applies the NEQ predicate on the "password_hash" field.
func PasswordHashNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPasswordHash, v))
}

// PasswordHashIn applies the In predicate on the "password_hash" field.
func PasswordHashIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPasswordHash, vs...))
}

// PasswordHashNotIn applies the NotIn predicate on the "password_hash" field.
func PasswordHashNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPasswordHash, vs...))
}

// PasswordHashGT applies the GT predicate on the "password_hash" field.
func PasswordHashGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPasswordHash, v))
}

// PasswordHashGTE applies the GTE predicate on the "password_hash" field.
func PasswordHashGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPasswordHash, v))
}

// PasswordHashLT applies the LT predicate on the "password_hash" field.
func PasswordHashLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPasswordHash, v))
}

// PasswordHashLTE applies the LTE predicate on the "password_hash" field.
func PasswordHashLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPasswordHash, v))
}

// PasswordHashContains applies the Contains predicate on the "password_hash" field.
func PasswordHashContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPasswordHash, v))
}

// PasswordHashHasPrefix applies the HasPrefix predicate on the "password_hash" field.
func PasswordHashHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPasswordHash, v))
}

// PasswordHashHasSuffix applies the HasSuffix predicate on the "password_hash" field.
func PasswordHashHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPasswordHash, v))
}

// PasswordHashEqualFold applies the EqualFold predicate on the "password_hash" field.
func PasswordHashEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPasswordHash, v))
}

// PasswordHashContainsFold applies the ContainsFold predicate on the "password_hash" field.
func PasswordHashContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPasswordHash, v))
}

// PointsEQ applies the EQ predicate on the "points" field.
func PointsEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPoints, v))
}

// PointsNEQ applies the NEQ predicate on the "points" field.
func PointsNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPoints, v))
}

// PointsIn applies the In predicate on the "points" field.
func PointsIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldPoints, vs...))
}

// PointsNotIn applies the NotIn predicate on the "points" field.
func PointsNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPoints, vs...))
}

// PointsGT applies the GT predicate on the "points" field.
func PointsGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldPoints, v))
}

// PointsGTE applies the GTE predicate on the "points" field.
func PointsGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPoints, v))
}

// PointsLT applies the LT predicate on the "points" field.
func PointsLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldPoints, v))
}

// PointsLTE applies the LTE predicate on the "points" field.
func PointsLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPoints, v))
}

// PointsAwardedCountEQ applies the EQ predicate on the "points_awarded_count" field.
func PointsAwardedCountEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPointsAwardedCount, v))
}

// PointsAwardedCountNEQ applies the NEQ predicate on the "points_awarded_count" field.
func PointsAwardedCountNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPointsAwardedCount, v))
}

// PointsAwardedCountIn applies the In predicate on the "points_awarded_count" field.
func PointsAwardedCountIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldPointsAwardedCount, vs...))
}

// PointsAwardedCountNotIn applies the NotIn predicate on the "points_awarded_count" field.
func PointsAwardedCountNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPointsAwardedCount, vs...))
}

// PointsAwardedCountGT applies the GT predicate on the "points_awarded_count" field.
func PointsAwardedCountGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldPointsAwardedCount, v))
}

// PointsAwardedCountGTE applies the GTE predicate on the "points_awarded_count" field.
func PointsAwardedCountGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPointsAwardedCount, v))
}

// PointsAwardedCountLT applies the LT predicate on the "points_awarded_count" field.
func PointsAwardedCountLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldPointsAwardedCount, v))
}

// PointsAwardedCountLTE applies the LTE predicate on the "points_awarded_count" field.
func PointsAwardedCountLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPointsAwardedCount, v))
}

// PointsAwardedResetTimeEQ applies the EQ predicate on the "points_awarded_reset_time" field.
func PointsAwardedResetTimeEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPointsAwardedResetTime, v))
}

// PointsAwardedResetTimeNEQ applies the NEQ predicate on the "points_awarded_reset_time" field.
func PointsAwardedResetTimeNEQ(v time.Time) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPointsAwardedResetTime, v))
}

// PointsAwardedResetTimeIn applies the In predicate on the "points_awarded_reset_time" field.
func PointsAwardedResetTimeIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldIn(FieldPointsAwardedResetTime, vs...))
}

// PointsAwardedResetTimeNotIn applies the NotIn predicate on the "points_awarded_reset_time" field.
func PointsAwardedResetTimeNotIn(vs ...time.Time) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPointsAwardedResetTime, vs...))
}

// PointsAwardedResetTimeGT applies the GT predicate on the "points_awarded_reset_time" field.
func PointsAwardedResetTimeGT(v time.Time) predicate.User {
	return predicate.User(sql.FieldGT(FieldPointsAwardedResetTime, v))
}

// PointsAwardedResetTimeGTE applies the GTE predicate on the "points_awarded_reset_time" field.
func PointsAwardedResetTimeGTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPointsAwardedResetTime, v))
}

// PointsAwardedResetTimeLT applies the LT predicate on the "points_awarded_reset_time" field.
func PointsAwardedResetTimeLT(v time.Time) predicate.User {
	return predicate.User(sql.FieldLT(FieldPointsAwardedResetTime, v))
}

// PointsAwardedResetTimeLTE applies the LTE predicate on the "points_awarded_reset_time" field.
func PointsAwardedResetTimeLTE(v time.Time) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPointsAwardedResetTime, v))
}

// PointsAwardedResetTimeIsNil applies the IsNil predicate on the "points_awarded_reset_time" field.
func PointsAwardedResetTimeIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPointsAwardedResetTime))
}

// PointsAwardedResetTimeNotNil applies the NotNil predicate on the "points_awarded_reset_time" field.
func PointsAwardedResetTimeNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPointsAwardedResetTime))
}

// GodModeEQ applies the EQ predicate on the "god_mode" field.
func GodModeEQ(v bool) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGodMode, v))
}

// GodModeNEQ applies the NEQ predicate on the "god_mode" field.
func GodModeNEQ(v bool) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGodMode, v))
}

// HasCourse applies the HasEdge predicate on the "course" edge.
func HasCourse() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CourseTable, CourseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCourseWith applies the HasEdge predicate on the "course" edge with a given conditions (other predicates).
func HasCourseWith(preds ...predicate.Course) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCourseStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInstitution applies the HasEdge predicate on the "institution" edge.
func HasInstitution() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, InstitutionTable, InstitutionPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInstitutionWith applies the HasEdge predicate on the "institution" edge with a given conditions (other predicates).
func HasInstitutionWith(preds ...predicate.Institution) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newInstitutionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRedemptions applies the HasEdge predicate on the "redemptions" edge.
func HasRedemptions() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, RedemptionsTable, RedemptionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRedemptionsWith applies the HasEdge predicate on the "redemptions" edge with a given conditions (other predicates).
func HasRedemptionsWith(preds ...predicate.Redemption) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newRedemptionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasForumPosts applies the HasEdge predicate on the "forum_posts" edge.
func HasForumPosts() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ForumPostsTable, ForumPostsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasForumPostsWith applies the HasEdge predicate on the "forum_posts" edge with a given conditions (other predicates).
func HasForumPostsWith(preds ...predicate.ForumPost) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newForumPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPet applies the HasEdge predicate on the "pet" edge.
func HasPet() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, PetTable, PetPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPetWith applies the HasEdge predicate on the "pet" edge with a given conditions (other predicates).
func HasPetWith(preds ...predicate.Pet) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, GroupsTable, GroupsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newGroupsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReactedPosts applies the HasEdge predicate on the "reacted_posts" edge.
func HasReactedPosts() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ReactedPostsTable, ReactedPostsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReactedPostsWith applies the HasEdge predicate on the "reacted_posts" edge with a given conditions (other predicates).
func HasReactedPostsWith(preds ...predicate.ForumPost) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newReactedPostsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUserPets applies the HasEdge predicate on the "user_pets" edge.
func HasUserPets() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserPetsTable, UserPetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserPetsWith applies the HasEdge predicate on the "user_pets" edge with a given conditions (other predicates).
func HasUserPetsWith(preds ...predicate.UserPet) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newUserPetsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroupUsers applies the HasEdge predicate on the "group_users" edge.
func HasGroupUsers() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, GroupUsersTable, GroupUsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupUsersWith applies the HasEdge predicate on the "group_users" edge with a given conditions (other predicates).
func HasGroupUsersWith(preds ...predicate.GroupUser) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newGroupUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReactions applies the HasEdge predicate on the "reactions" edge.
func HasReactions() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ReactionsTable, ReactionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReactionsWith applies the HasEdge predicate on the "reactions" edge with a given conditions (other predicates).
func HasReactionsWith(preds ...predicate.Reaction) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newReactionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
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
func Not(p predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		p(s.Not())
	})
}
