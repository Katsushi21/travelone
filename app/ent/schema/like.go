package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

func (Like) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Friends"},
		field.ID("account_id", "post_id"),
	}
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id").
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
		field.String("post_id").
			Annotations(
				entgql.OrderField("POST_ID"),
			),
	}
}

func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("account", Account.Type).
			Unique().
			Required().
			Field("account_id"),
		edge.To("post", Post.Type).
			Unique().
			Required().
			Field("post_id"),
	}
}
