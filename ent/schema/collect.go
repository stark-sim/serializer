package schema

import (
	"entgo.io/ent"
)

type Collect struct {
	ent.Schema
}

func (Collect) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the ServerOrder.
func (Collect) Edges() []ent.Edge {
	return []ent.Edge{}
}
