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
		field.UUID("account_id", uuid.UUID{}).
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
			Required().
			Unique().
			Field("account_id"),
	}
}
