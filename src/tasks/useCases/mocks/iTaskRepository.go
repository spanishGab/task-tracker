package mocks

import (
	"fmt"
	"tasktracker/src/tasks"
)

type TaskRepositorySuccessfullMock struct{}

func (tr *TaskRepositorySuccessfullMock) CreateOne(task tasks.Task) (*tasks.Task, error) {
	return &task, nil
}

func (tr *TaskRepositorySuccessfullMock) DeleteOne(id uint64) error {
	return nil
}

func (tr *TaskRepositorySuccessfullMock) UpdateOne(task tasks.Task) (*tasks.Task, error) {
	return &task, nil
}

type TaskRepositoryErrorMock struct{}

func (tr *TaskRepositoryErrorMock) CreateOne(task tasks.Task) (*tasks.Task, error) {
	return nil, fmt.Errorf("could not create task")
}

func (tr *TaskRepositoryErrorMock) DeleteOne(id uint64) error {
	return fmt.Errorf("could not delete task")
}

func (tr *TaskRepositoryErrorMock) UpdateOne(task tasks.Task) (*tasks.Task, error) {
	return nil, fmt.Errorf("could not update task")
}
