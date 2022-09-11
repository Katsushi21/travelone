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
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
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
		field.UUID("account_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
		field.UUID("request_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("REQUEST_ID"),
			),
		field.Enum("status").Values(
			"inProcess",
			"accept",
			"deny",
			"breakInProcess",
			"breakAccept",
			"breakDeny",
		).
			Annotations(
				entgql.OrderField("STATUS"),
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
