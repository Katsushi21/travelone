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

type Request struct {
	ent.Schema
}

func (Request) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Requests"},
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Request) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
	}
}

func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("account_id", uuid.UUID{}),
		field.UUID("request_id", uuid.UUID{}),
		field.Enum("status").Values(
			"IN_PROCESS",
			"ACCEPT",
			"DENY",
			"BREAK_IN_PROCESS",
			"BREAK_IN_ACCEPT",
			"BREAK_DENY",
		),
	}
}

func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			Required().
			Unique().
			Field("account_id"),
		edge.To("request", Account.Type).
			Required().
			Unique().
			Field("request_id"),
	}
}
