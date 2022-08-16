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

type Account struct {
	ent.Schema
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Accounts"},
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New).
			Annotations(
				entgql.OrderField("ID"),
			),
		field.String("email").
			MaxLen(100).
			NotEmpty().
			Unique().
			Annotations(
				entgql.OrderField("EMAIL"),
			),
		field.String("password").
			Annotations(
				entgql.OrderField("PASSWORD"),
			),
		field.Enum("type").
			NamedValues(
				"active",
				"inactive",
				"admin",
			).
			Annotations(
				entgql.OrderField("ACCOUNT_TYPE"),
			),
		field.String("name").
			Default("guest").
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.Int("age").Positive().
			Annotations(
				entgql.OrderField("AGE"),
			),
		field.Enum("gender").
			NamedValues(
				"male",
				"female",
			).
			Annotations(
				entgql.OrderField("GENDER"),
			),
		field.String("avatar").
			Annotations(
				entgql.OrderField("AVATAR"),
			),
		field.String("introduction").
			Annotations(
				entgql.OrderField("INTRODUCTION"),
			),
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", Account.Type).
			Through("friendships", FriendShip.Type),
	}
}
