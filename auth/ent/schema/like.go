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

type Like struct {
	ent.Schema
}

func (Like) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Likes"},
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Like) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
	}
}

func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("account_id", uuid.UUID{}),
		field.UUID("post_id", uuid.UUID{}),
	}
}

func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("likes").
			Unique().
			Required().
			Field("account_id"),
		edge.From("post", Post.Type).
			Ref("likes").
			Unique().
			Required().
			Field("post_id"),
	}
}
