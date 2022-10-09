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

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Comments"},
	}
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
	}
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("body").
			Annotations(
				entgql.OrderField("BODY"),
			),
		field.UUID("account_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
		field.UUID("post_id", uuid.UUID{}).
			Annotations(
				entgql.OrderField("POST_ID"),
			),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).
			Ref("comments").
			Required().
			Unique().
			Field("post_id"),
		edge.From("account", Account.Type).
			Ref("comments").
			Required().
			Unique().
			Field("account_id"),
	}
}