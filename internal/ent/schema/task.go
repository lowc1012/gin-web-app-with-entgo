package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.New()),
		field.Text("title").NotEmpty(),
		field.Text("description").Optional(),
		field.Int("priority"),
		field.UUID("todo_id", uuid.Nil).Optional(),
		field.UUID("parent_id", uuid.Nil).Optional(),
		field.String("status").NotEmpty(),
		field.Time("created_at").Immutable(),
		field.Time("updated_at").UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("todo", Todo.Type).Ref("tasks").Field("todo_id").Unique(), // O2M or M2O
		edge.To("parent", Task.Type).Unique().Field("parent_id").From("children"),
	}
}
