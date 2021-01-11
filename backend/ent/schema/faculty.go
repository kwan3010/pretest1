package schema

import "github.com/facebookincubator/ent"

// Faculty holds the schema definition for the Faculty entity.
type Faculty struct {
	ent.Schema
}

// Fields of the Faculty.
func (Faculty) Fields() []ent.Field {
	return nil
}

// Edges of the Faculty.
func (Faculty) Edges() []ent.Edge {
	return nil
}
