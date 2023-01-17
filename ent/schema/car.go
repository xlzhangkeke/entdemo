package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Float("amount").SchemaType(map[string]string{
			dialect.MySQL:    "decimal(6,2)", // Override MySQL.
			dialect.Postgres: "numeric",      // Override Postgres.
		}),
		field.Time("registered_at"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		// setting the edge to unique, ensure
		// that a car can have only one owner.
		edge.From("owner", User.Type).Ref("cars").Unique(),
	}
}
