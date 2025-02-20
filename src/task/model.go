package task

import "time"

type Status string

// Task statuses
const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

type Task struct {
	id          uint64    `json:"id"`
	description string    `json:"description"`
	status      Status    `json:"status"`
	createdAt   time.Time `json:"created_at"`
	updatedAt   time.Time `json:"updated_at"`
}
