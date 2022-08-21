package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Request struct {
	ent.Schema
}

func (Request) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Requests"},
	}
}

func (Request) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id").
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
		field.String("request_id").
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
