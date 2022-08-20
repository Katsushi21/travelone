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
		TimeMixin{},
	}
}

func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Immutable().
			Default(uuid.New).
			Annotations(
				entgql.OrderField("ID"),
			),
		field.String("body").
			Annotations(
				entgql.OrderField("BODY"),
			),
	}
}

func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).
			Ref("comments").
			Unique(),
		edge.From("account", Account.Type).
			Ref("comments").
			Unique(),
	}
}
