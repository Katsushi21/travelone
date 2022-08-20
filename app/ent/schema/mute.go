package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Mute holds the schema definition for the Mute entity.
type Mute struct {
	ent.Schema
}

func (Mute) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Mutes"},
	}
}

func (Mute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Mute.
func (Mute) Fields() []ent.Field {
	return []ent.Field{
		field.String("account_id").
			Annotations(
				entgql.OrderField("ACCOUNT_ID"),
			),
		field.String("mute_id").
			Annotations(
				entgql.OrderField("MUTE_ID"),
			),
	}
}

// Edges of the Mute.
func (Mute) Edges() []ent.Edge {
	return nil
}
