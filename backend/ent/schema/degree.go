package schema

import "github.com/facebookincubator/ent"

// Degree holds the schema definition for the Degree entity.
type Degree struct {
	ent.Schema
}

// Fields of the Degree.
func (Degree) Fields() []ent.Field {
	return []ent.Field{
		
	}
 
}

// Edges of the Degree.
func (Degree) Edges() []ent.Edge {
	return nil
}
