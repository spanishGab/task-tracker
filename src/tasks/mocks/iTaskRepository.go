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

func (tr *TaskRepositorySuccessfullMock) GetAllTasks() ([]tasks.Task, error) {
	return []tasks.Task{}, nil
}

func (tr *TaskRepositorySuccessfullMock) GetAllTasksByStatus(status tasks.Status) ([]tasks.Task, error) {
	return []tasks.Task{}, nil
}

func (tr *TaskRepositorySuccessfullMock) Format(tasks []tasks.Task) ([]byte, error) {
	return []byte{}, nil
}

type TaskRepositoryFailureMock struct{}

func (tr *TaskRepositoryFailureMock) CreateOne(task tasks.Task) (*tasks.Task, error) {
	return nil, fmt.Errorf("could not create task")
}

func (tr *TaskRepositoryFailureMock) DeleteOne(id uint64) error {
	return fmt.Errorf("could not delete task")
}

func (tr *TaskRepositoryFailureMock) UpdateOne(task tasks.Task) (*tasks.Task, error) {
	return nil, fmt.Errorf("could not update task")
}

func (tr *TaskRepositoryFailureMock) GetAllTasks() ([]tasks.Task, error) {
	return nil, fmt.Errorf("could not search tasks")
}

func (tr *TaskRepositoryFailureMock) GetAllTasksByStatus(status tasks.Status) ([]tasks.Task, error) {
	return nil, fmt.Errorf("could not search tasks")
}

func (tr *TaskRepositoryFailureMock) Format(tasks []tasks.Task) ([]byte, error) {
	return nil, fmt.Errorf("could not transform tasks to bytes")
}
