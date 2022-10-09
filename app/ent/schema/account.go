package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		UuidMixin{},
		TimeMixin{},
	}
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").
			MaxLen(100).
			NotEmpty().
			Unique().
			Annotations(
				entgql.OrderField("EMAIL"),
			),
		field.String("password").
			Sensitive().
			Annotations(
				entgql.OrderField("PASSWORD"),
			),
		field.Enum("type").
			Values(
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
			Values(
				"male",
				"female",
				"none",
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
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
		edge.To("friends", Account.Type).
			Through("friendships", Friend.Type),
		edge.To("mutes", Mute.Type),
		edge.To("requests", Account.Type).
			Through("requestTargets", Request.Type),
		edge.To("likes", Like.Type),
		edge.To("session", Session.Type),
	}
}
