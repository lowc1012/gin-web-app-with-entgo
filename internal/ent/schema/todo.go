package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("title").NotEmpty(),
		field.Text("description").Optional(),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").UpdateDefault(time.Now()),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type),
	}
}
