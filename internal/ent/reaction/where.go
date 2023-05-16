// Code generated by ent, DO NOT EDIT.

package reaction

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/np-inprove/server/internal/ent/predicate"
)

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldUserID, v))
}

// ForumPostID applies equality check predicate on the "forum_post_id" field. It's identical to ForumPostIDEQ.
func ForumPostID(v int) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldForumPostID, v))
}

// Emoji applies equality check predicate on the "emoji" field. It's identical to EmojiEQ.
func Emoji(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldEmoji, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldUserID, vs...))
}

// ForumPostIDEQ applies the EQ predicate on the "forum_post_id" field.
func ForumPostIDEQ(v int) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldForumPostID, v))
}

// ForumPostIDNEQ applies the NEQ predicate on the "forum_post_id" field.
func ForumPostIDNEQ(v int) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldForumPostID, v))
}

// ForumPostIDIn applies the In predicate on the "forum_post_id" field.
func ForumPostIDIn(vs ...int) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldForumPostID, vs...))
}

// ForumPostIDNotIn applies the NotIn predicate on the "forum_post_id" field.
func ForumPostIDNotIn(vs ...int) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldForumPostID, vs...))
}

// EmojiEQ applies the EQ predicate on the "emoji" field.
func EmojiEQ(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldEQ(FieldEmoji, v))
}

// EmojiNEQ applies the NEQ predicate on the "emoji" field.
func EmojiNEQ(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldNEQ(FieldEmoji, v))
}

// EmojiIn applies the In predicate on the "emoji" field.
func EmojiIn(vs ...string) predicate.Reaction {
	return predicate.Reaction(sql.FieldIn(FieldEmoji, vs...))
}

// EmojiNotIn applies the NotIn predicate on the "emoji" field.
func EmojiNotIn(vs ...string) predicate.Reaction {
	return predicate.Reaction(sql.FieldNotIn(FieldEmoji, vs...))
}

// EmojiGT applies the GT predicate on the "emoji" field.
func EmojiGT(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldGT(FieldEmoji, v))
}

// EmojiGTE applies the GTE predicate on the "emoji" field.
func EmojiGTE(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldGTE(FieldEmoji, v))
}

// EmojiLT applies the LT predicate on the "emoji" field.
func EmojiLT(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldLT(FieldEmoji, v))
}

// EmojiLTE applies the LTE predicate on the "emoji" field.
func EmojiLTE(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldLTE(FieldEmoji, v))
}

// EmojiContains applies the Contains predicate on the "emoji" field.
func EmojiContains(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldContains(FieldEmoji, v))
}

// EmojiHasPrefix applies the HasPrefix predicate on the "emoji" field.
func EmojiHasPrefix(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldHasPrefix(FieldEmoji, v))
}

// EmojiHasSuffix applies the HasSuffix predicate on the "emoji" field.
func EmojiHasSuffix(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldHasSuffix(FieldEmoji, v))
}

// EmojiEqualFold applies the EqualFold predicate on the "emoji" field.
func EmojiEqualFold(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldEqualFold(FieldEmoji, v))
}

// EmojiContainsFold applies the ContainsFold predicate on the "emoji" field.
func EmojiContainsFold(v string) predicate.Reaction {
	return predicate.Reaction(sql.FieldContainsFold(FieldEmoji, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasForumPost applies the HasEdge predicate on the "forum_post" edge.
func HasForumPost() predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ForumPostColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, ForumPostTable, ForumPostColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasForumPostWith applies the HasEdge predicate on the "forum_post" edge with a given conditions (other predicates).
func HasForumPostWith(preds ...predicate.ForumPost) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		step := newForumPostStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
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
func Not(p predicate.Reaction) predicate.Reaction {
	return predicate.Reaction(func(s *sql.Selector) {
		p(s.Not())
	})
}