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

type Friend struct {
	ent.Schema
}

func (Friend) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Friends"},
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}

func (Friend) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
	}
}

func (Friend) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("account_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
		field.UUID("friend_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("FRIEND_ID"),
			),
	}
}

func (Friend) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			Required().
			Unique().
			Field("account_id"),
		edge.To("friend", Account.Type).
			Required().
			Unique().
			Field("friend_id"),
	}
}
