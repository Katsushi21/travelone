package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type FriendShip struct {
	ent.Schema
}

func (FriendShip) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "FriendShips"},
	}
}

func (FriendShip) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (FriendShip) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id").
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
		field.String("friend_id").
			Annotations(
				entgql.OrderField("FRIEND_ID"),
			),
	}
}

func (FriendShip) Edges() []ent.Edge {
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
