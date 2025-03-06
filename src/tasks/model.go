package tasks

import (
	"time"
)

type Status string

// Task status
const (
	Todo       Status = "todo"
	InProgress Status = "in-progress"
	Done       Status = "done"
)

type Task struct {
	ID          uint64    `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewTask(description string) *Task {
	return &Task{
		Description: description,
		Status:      Todo,
	}
}
