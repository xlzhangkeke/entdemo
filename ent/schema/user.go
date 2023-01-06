package schema

import (
	"net/url"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").Positive(),
		field.String("name").Unique(),
		field.Float("rank").Optional(),
		field.Bool("active").Default(false),
		field.Time("created_at").Default(time.Now),
		field.JSON("url", &url.URL{}).Optional(),
		field.JSON("strings", []string{}).Optional(),
		field.Enum("state").Values("on", "off").Optional(),
		field.UUID("uuid", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cars", Car.Type),
		// Create an inverse-edge called "groups" of type `Group`
		// and reference it to the "users" edge (in Group schema)
		// explicitly using the `Ref` method.
		edge.From("groups", Group.Type).Ref("users"),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("age", "name").Unique(),
	}
}
