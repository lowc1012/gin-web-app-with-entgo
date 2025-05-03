package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Text("title").NotEmpty(),
		field.Text("description").Optional().Nillable(),
		field.Int("priority").Default(0),
		field.Int("todo_id").Optional().Nillable(),
		field.Int("parent_id").Optional().Nillable(),
		field.Enum("status").NamedValues(
			"InProgress", "IN_PROGRESS",
			"Completed", "COMPLETED",
			"NotYet", "NOT_YET").Default("NOT_YET"),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").UpdateDefault(time.Now()),
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
