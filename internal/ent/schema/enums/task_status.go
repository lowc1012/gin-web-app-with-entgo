package enums

// TaskStatus defines the possible statuses for a Task item.
type TaskStatus string

const (
	// TaskStatusPending represents a task item that is pending.
	TaskStatusPending TaskStatus = "PENDING"
	// TaskStatusInProgress represents a task item that is in progress.
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	// TaskStatusCompleted represents a task item that has been completed.
	TaskStatusCompleted TaskStatus = "COMPLETED"
)

// Values provides list valid values for Enum.
func (TaskStatus) Values() []string {
	return []string{
		string(TaskStatusPending),
		string(TaskStatusInProgress),
		string(TaskStatusCompleted),
	}
}

func (e TaskStatus) String() string {
	return string(e)
}
