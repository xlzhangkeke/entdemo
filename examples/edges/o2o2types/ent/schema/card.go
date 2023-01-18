package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.Time("expired"),
		field.String("number"),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		// We add the "Required" method to the builder
		// to make this edge required on entity creation.
		// i.e. Card cannot be created without its owner.
		edge.From("owner", User.Type).Ref("card").Unique().Required(),
	}
}
