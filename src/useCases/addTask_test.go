package usecases

import (
	"fmt"
	"tasktracker/src/commands"
	"tasktracker/src/ports"
	"tasktracker/src/tasks"
	"testing"
)

type TaskRepositorySuccessfullMock struct{}

func (tr *TaskRepositorySuccessfullMock) CreateOne(task tasks.Task) (*tasks.Task, error) {
	return &task, nil
}

type TaskRepositoryErrorMock struct{}

func (tr *TaskRepositoryErrorMock) CreateOne(task tasks.Task) (*tasks.Task, error) {
	return nil, fmt.Errorf("could not create task")
}

func TestAdd_ParseArgs(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		taskRepository ports.ITaskRepository
		wantErr        bool
	}{
		{
			name:           "should return no error for valid args",
			args:           []string{"arg1"},
			taskRepository: &TaskRepositorySuccessfullMock{},
			wantErr:        false,
		},
		{
			name:           "should return an error for no args",
			args:           []string{},
			taskRepository: &TaskRepositoryErrorMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for nil args",
			args:           nil,
			taskRepository: &TaskRepositoryErrorMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for multiple args",
			args:           []string{"arg1", "arg2"},
			taskRepository: &TaskRepositoryErrorMock{},
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := commands.NewCommand("add", test.args)
			addTask := NewAddTask(test.taskRepository)
			err := addTask.parseArgs(*cmd)
			assertError(t, err, test.wantErr)
		})
	}
}

func assertError(t testing.TB, err error, wantError bool) {
	if (err != nil) != wantError {
		t.Errorf("wantError? = %t, got: %v", wantError, err)
	}
}
