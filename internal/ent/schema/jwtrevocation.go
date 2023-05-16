package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// JWTRevocation holds the schema definition for the JWTRevocation entity.
type JWTRevocation struct {
	ent.Schema
}

// Fields of the JWTRevocation.
func (JWTRevocation) Fields() []ent.Field {
	return []ent.Field{
		field.String("jti").
			NotEmpty().
			Immutable().
			Comment("ID of the revoked JWT"),
		field.Time("expiry").
			Immutable().
			Comment("Expiry of the JWT.\nThe revocation can be deleted after the expiry has passed"),
	}
}

// Edges of the JWTRevocation.
func (JWTRevocation) Edges() []ent.Edge {
	return nil
}
