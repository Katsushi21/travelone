package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// FriendShip holds the schema definition for the FriendShip entity.
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
		field.Int("account_id"),
		field.Int("friend_id"),
	}
}

func (FriendShip) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", Account.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("friend", Account.Type).
			Required().
			Unique().
			Field("friend_id"),
	}
}
