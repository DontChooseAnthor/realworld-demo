package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Article holds the schema definition for the Article entity.
type Article struct {
	ent.Schema
}

// Fields of the Article.
func (Article) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Unique(),
		field.Uint32("author"),
		field.String("description").Optional(),
		field.String("body"),
		field.JSON("tags", []string{}).
			Optional(),
		field.Time("created_at").
			Default(time.Now),
		field.Time("last_modify"),
	}
}

// Edges of the Article.
func (Article) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("articles").
			Unique().
			Field("author").
			Required(),
	}
}
