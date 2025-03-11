package ports

import "tasktracker/src/tasks"

type ITaskRepository interface {
	CreateOne(task tasks.Task) (*tasks.Task, error)
	DeleteOne(id uint64) error
}
