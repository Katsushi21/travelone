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

type Marker struct {
	ent.Schema
}

func (Marker) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Markers"},
	}
}

func (Marker) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (Marker) Fields() []ent.Field {
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
		field.String("lat").
			Annotations(
				entgql.OrderField("LAT"),
			),
		field.String("lng").
			Annotations(
				entgql.OrderField("LNG"),
			),
		field.String("post_id").
			Annotations(
				entgql.OrderField("POST_ID"),
			),
	}
}

func (Marker) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).
			Ref("marker").
			Unique().
			Field("post_id"),
	}
}
