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
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (Marker) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UuidMixin{},
		TimeMixin{},
	}
}

func (Marker) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.String("lat"),
		field.String("lng"),
		field.UUID("post_id", uuid.UUID{}),
	}
}

func (Marker) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("post", Post.Type).
			Ref("marker").
			Required().
			Unique().
			Field("post_id"),
	}
}
