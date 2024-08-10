package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").NotEmpty(),
		field.Enum("status").NamedValues(
			"InProgress", "IN_PROGRESS",
			"Completed", "COMPLETED", "NotYet", "NOT_YET").Default("NOT_YET"),
		field.Int("priority").Default(0),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.Time("updated_at").Default(time.Now()),
	}

}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
