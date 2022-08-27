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

type Post struct {
	ent.Schema
}

func (Post) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Posts"},
	}
}

func (Post) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New).
			Annotations(
				entgql.OrderField("ID"),
			),
		field.String("title").
			Annotations(
				entgql.OrderField("TITLE"),
			),
		field.String("body").
			Annotations(
				entgql.OrderField("BODY"),
			),
		field.String("img").
			Annotations(
				entgql.OrderField("IMG"),
			),
		field.UUID("account_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
	}
}

func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("comments", Comment.Type),
		edge.To("marker", Marker.Type).
			Unique(),
		edge.From("account", Account.Type).
			Ref("posts").
			Required().
			Unique().
			Field("account_id"),
	}
}
