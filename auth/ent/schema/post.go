package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type Post struct {
	ent.Schema
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Posts"},
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
	}
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("body"),
		field.String("img"),
		field.UUID("account_id", uuid.UUID{}),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
		edge.To("marker", Marker.Type).
			Unique(),
		edge.From("account", Account.Type).
			Ref("posts").
			Required().
			Unique().
			Field("account_id"),
		edge.To("likes", Like.Type),
	}
}
