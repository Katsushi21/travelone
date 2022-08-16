package schema

import "entgo.io/ent"

// Marker holds the schema definition for the Marker entity.
type Marker struct {
	ent.Schema
}

// Fields of the Marker.
func (Marker) Fields() []ent.Field {
	return nil
}

// Edges of the Marker.
func (Marker) Edges() []ent.Edge {
	return nil
}
