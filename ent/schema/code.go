package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Code holds the schema definition for the Code entity.
type Code struct {
	ent.Schema
}

// Fields of the Code.
func (Code) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			MaxLen(100).
			Unique().
			Immutable(),
		field.String("email").
			MaxLen(100).
			Immutable(),
		field.String("action").
			MaxLen(100).
			Immutable(),
		field.String("value_hash").
			MaxLen(100).
			Immutable(),
		field.Int("expires_at").
			Immutable(),
		field.Int("created_at").
			Immutable(),
	}
}

// Edges of the Code.
func (Code) Edges() []ent.Edge {
	return nil
}
