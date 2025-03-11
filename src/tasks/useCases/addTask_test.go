package usecases

import (
	"tasktracker/src/commands"
	"tasktracker/src/tasks/ports"
	"tasktracker/src/tasks/useCases/mocks"
	"testing"
)

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
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        false,
		},
		{
			name:           "should return an error for no args",
			args:           []string{},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for nil args",
			args:           nil,
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for multiple args",
			args:           []string{"arg1", "arg2"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
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
