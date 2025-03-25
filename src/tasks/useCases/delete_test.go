package usecases

import (
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/mocks"
	"testing"
)

func TestDeleteTask_parseCommand(t *testing.T) {
	tests := []struct {
		name           string
		cmdName        commands.CommandName
		args           []string
		taskRepository tasks.ITaskRepository
		expected       uint64
		wantErr        bool
	}{
		{
			name:           "should return no error for valid args",
			cmdName:        "delete",
			args:           []string{"123"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expected:       123,
			wantErr:        false,
		},
		{
			name:           "should return an error for invalid command name",
			cmdName:        "del",
			args:           []string{"123"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for no args",
			cmdName:        "delete",
			args:           []string{},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for nil args",
			cmdName:        "delete",
			args:           nil,
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for multiple args",
			cmdName:        "delete",
			args:           []string{"123", "extra_arg"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for non-numeric id",
			cmdName:        "delete",
			args:           []string{"abc"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := commands.NewCommand(test.cmdName, test.args)
			deleteTask := NewDeleteTask(test.taskRepository)
			got, err := deleteTask.parseCommand(*cmd)
			assertError(t, err, test.wantErr)
			assertTask(t, &tasks.Task{ID: got}, &tasks.Task{ID: test.expected})
		})
	}
}
func TestDeleteTask_Execute(t *testing.T) {
	tests := []struct {
		name           string
		cmdName        commands.CommandName
		args           []string
		taskRepository tasks.ITaskRepository
		expected       *string
		wantErr        bool
	}{
		{
			name:           "should delete a task successfully",
			cmdName:        "delete",
			args:           []string{"123"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expected:       &deleteResult,
			wantErr:        false,
		},
		{
			name:           "should return an error when failing to parse the given command",
			cmdName:        "delete",
			args:           []string{"abc"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error when repository fails to delete",
			cmdName:        "delete",
			args:           []string{"123"},
			taskRepository: &mocks.TaskRepositoryFailureMock{},
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := commands.NewCommand(test.cmdName, test.args)
			deleteTask := NewDeleteTask(test.taskRepository)
			got, err := deleteTask.Execute(*cmd)
			assertError(t, err, test.wantErr)
			if !test.wantErr {
				assertResultString(t, got, test.expected)
			}
		})
	}
}
