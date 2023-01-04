package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
        field.Uint32("id").Unique(),
		field.String("username"),
		field.String("password").Default("123456"),
        field.Uint32("level").Default(0),
		field.String("bio").Optional(),
		field.String("image").Optional(),
		field.Time("created_at").Default(time.Now),
		field.Uint32("points").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("articles", Article.Type),
	}
}
