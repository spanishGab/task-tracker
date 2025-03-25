package usecases

import (
	"fmt"
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/mocks"

	"testing"
)

func TestAddTask_parseCommand(t *testing.T) {
	tests := []struct {
		name           string
		cmdName        commands.CommandName
		args           []string
		taskRepository tasks.ITaskRepository
		expected       *tasks.Task
		wantErr        bool
	}{
		{
			name:           "should return no error for valid args",
			cmdName:        "add",
			args:           []string{"arg1"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expected:       &tasks.Task{Description: "arg1", Status: tasks.Todo},
			wantErr:        false,
		},
		{
			name:           "should return an error for invalid command name",
			cmdName:        "create",
			args:           []string{"123"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for no args",
			cmdName:        "add",
			args:           []string{},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for nil args",
			cmdName:        "add",
			args:           nil,
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for multiple args",
			cmdName:        "add",
			args:           []string{"arg1", "arg2"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := commands.NewCommand(test.cmdName, test.args)
			addTask := NewAddTask(test.taskRepository)
			got, err := addTask.parseCommand(*cmd)
			assertError(t, err, test.wantErr)
			assertTask(t, got, test.expected)
		})
	}
}
func TestAddTask_Execute(t *testing.T) {
	tests := []struct {
		name           string
		cmdName        commands.CommandName
		args           []string
		taskRepository tasks.ITaskRepository
		expectedResult *string
		wantErr        bool
	}{
		{
			name:           "should add a task successfully",
			cmdName:        "add",
			args:           []string{"arg1"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expectedResult: func() *string { s := fmt.Sprintf(addResult, 0); return &s }(),
			wantErr:        false,
		},
		{
			name:           "should return an error when failing to parse the given command",
			cmdName:        "create",
			args:           []string{"123"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error when repository fails",
			cmdName:        "add",
			args:           []string{"arg1"},
			taskRepository: &mocks.TaskRepositoryFailureMock{},
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := commands.NewCommand(test.cmdName, test.args)
			addTask := NewAddTask(test.taskRepository)
			got, err := addTask.Execute(*cmd)
			assertError(t, err, test.wantErr)
			assertResultString(t, got, test.expectedResult)
		})
	}
}

func assertResultString(t testing.TB, got, expected *string) {
	t.Helper()
	if expected == nil {
		if got != nil {
			t.Errorf("expected: %#v, got: %#v", expected, got)
		}
	} else if *expected != *got {
		t.Errorf("expected: %s, got: %s", *expected, *got)
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
