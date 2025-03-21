package ports

import "tasktracker/src/tasks"

type ITaskRepository interface {
	CreateOne(task tasks.Task) (*tasks.Task, error)
	DeleteOne(id uint64) error
	UpdateOne(task tasks.Task) (*tasks.Task, error)
	GetAllTasks() ([]tasks.Task, error)
	TasksToBytes(tasks []tasks.Task) ([]byte, error)
}
