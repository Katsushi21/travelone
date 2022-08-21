package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Session struct {
	ent.Schema
}

func (Session) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Sessions"},
	}
}

func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id").
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
		field.String("session").
			Annotations(
				entgql.OrderField("SESSION"),
			),
	}
}

func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).
			Ref("session").
			Unique().
			Field("account_id"),
	}
}
