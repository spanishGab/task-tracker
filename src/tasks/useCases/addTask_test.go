package usecases

import (
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/ports"
	"tasktracker/src/tasks/useCases/mocks"
	"testing"
)

func TestAddTask_parseArgs(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		taskRepository ports.ITaskRepository
		expected       *tasks.Task
		wantErr        bool
	}{
		{
			name:           "should return no error for valid args",
			args:           []string{"arg1"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expected:       &tasks.Task{Description: "arg1"},
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
			got, err := addTask.parseArgs(*cmd)
			assertError(t, err, test.wantErr)
			assertTask(t, got, test.expected)
		})
	}
}

func assertError(t testing.TB, err error, wantError bool) {
	t.Helper()
	if (err != nil) != wantError {
		t.Errorf("wantError? = %t, got: %v", wantError, err)
	}
}

func assertTask(t testing.TB, got, expected *tasks.Task) {
	t.Helper()
	if expected == nil {
		if got != nil {
			t.Errorf("expected: %#v, got: %#v", expected, got)
		}
	} else if expected.Description != got.Description {
		t.Errorf("expected Description: %s, got Description: %s", expected.Description, got.Description)
	} else if expected.Status != got.Status {
		t.Errorf("expected Status: %s, got Status: %s", expected.Status, got.Status)
	}
}
