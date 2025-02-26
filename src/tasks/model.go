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
	id          uint64
	description string
	status      Status
	createdAt   time.Time
	updatedAt   time.Time
}

func NewTask(description string) *Task {
	return &Task{
		description: description,
		status:      Todo,
	}
}

func (t *Task) ID() uint64 {
	return t.id
}

func (t *Task) Description() string {
	return t.description
}

func (t *Task) Status() Status {
	return t.status
}

func (t *Task) CreatedAt() time.Time {
	return t.createdAt
}

func (t *Task) UpdatedAt() time.Time {
	return t.updatedAt
}

func (t *Task) SetID(id uint64) {
	t.id = id
}

func (t *Task) SetDescription(description string) {
	t.description = description
}

func (t *Task) SetStatus(status Status) {
	t.status = status
}

func (t *Task) SetCreatedAt(createdAt time.Time) {
	t.createdAt = createdAt
}

func (t *Task) SetUpdatedAt(updatedAt time.Time) {
	t.updatedAt = updatedAt
}
