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

// Mute holds the schema definition for the Mute entity.
type Mute struct {
	ent.Schema
}

func (Mute) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Mutes"},
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Mute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
	}
}

func (Mute) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("account_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
		field.UUID("mute_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("MUTE_ID"),
			),
	}
}

func (Mute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			Required().
			Unique().
			Field("account_id"),
		edge.To("mute", Account.Type).
			Required().
			Unique().
			Field("mute_id"),
	}
}
