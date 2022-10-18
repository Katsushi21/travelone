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
		field.UUID("account_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
		field.UUID("post_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("POST_ID"),
			),
	}
}

func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			Unique().
			Required().
			Field("account_id"),
		edge.To("post", Post.Type).
			Unique().
			Required().
			Field("post_id"),
	}
}
