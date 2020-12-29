package schema

import "github.com/facebookincubator/ent"

// Medical holds the schema definition for the Medical entity.
type Medical struct {
	ent.Schema
}

// Fields of the Medical.
func (Medical) Fields() []ent.Field {
	return nil
}

// Edges of the Medical.
func (Medical) Edges() []ent.Edge {
	return nil
}
