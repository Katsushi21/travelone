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
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
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
			MaxLen(256).
			NotEmpty().
			Unique(),
		field.String("password").
			Sensitive(),
		field.Enum("type").
			Values(
				"ACTIVE",
				"INACTIVE",
				"ADMIN",
			),
		field.String("name"),
		field.Int("age").
			Positive(),
		field.Enum("gender").
			Values(
				"MALE",
				"FEMALE",
				"NONE",
			),
		field.String("avatar"),
		field.String("introduction"),
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
		edge.To("friends", Account.Type).
			Through("friendship", Friend.Type),
		edge.To("mutes", Account.Type).
			Through("muteTarget", Mute.Type),
		edge.To("requests", Account.Type).
			Through("requestTarget", Request.Type),
		edge.To("likes", Like.Type),
		edge.To("session", Session.Type),
	}
}
