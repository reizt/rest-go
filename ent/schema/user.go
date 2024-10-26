package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(100).
			Unique().
			Immutable(),
		field.String("email").
			MaxLen(100),
		field.String("name").
			MaxLen(100),
		field.String("password_hash").
			MaxLen(100),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
