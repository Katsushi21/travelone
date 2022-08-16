package schema

import "entgo.io/ent"

// Mute holds the schema definition for the Mute entity.
type Mute struct {
	ent.Schema
}

// Fields of the Mute.
func (Mute) Fields() []ent.Field {
	return nil
}

// Edges of the Mute.
func (Mute) Edges() []ent.Edge {
	return nil
}
