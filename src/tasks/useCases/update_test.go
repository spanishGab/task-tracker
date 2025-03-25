package usecases

import (
	"fmt"
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/mocks"
	"testing"
)

func TestUpdateTask_parseCommand(t *testing.T) {
	tests := []struct {
		name           string
		command        *commands.Command
		taskRepository tasks.ITaskRepository
		expected       *tasks.Task
		wantErr        bool
	}{
		{
			name:           "should return no error for valid update command",
			command:        commands.NewCommand(commands.UpdateCommand, []string{"123", "new_value"}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expected: &tasks.Task{
				ID:          123,
				Description: "new_value",
			},
			wantErr: false,
		},
		{
			name:           "should return no error for valid mark in progress command",
			command:        commands.NewCommand(commands.MarkInProgressCommand, []string{"123"}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expected: &tasks.Task{
				ID:     123,
				Status: tasks.InProgress,
			},
			wantErr: false,
		},
		{
			name:           "should return no error for valid mark done command",
			command:        commands.NewCommand(commands.MarkDoneCommand, []string{"123"}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expected: &tasks.Task{
				ID:     123,
				Status: tasks.Done,
			},
			wantErr: false,
		},
		{
			name:           "should return an error for invalid command name",
			command:        commands.NewCommand("invalid_command", []string{"123"}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for non-numeric first arg",
			command:        commands.NewCommand(commands.UpdateCommand, []string{"abc", "new_value"}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for nil args",
			command:        commands.NewCommand(commands.UpdateCommand, nil),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for no args",
			command:        commands.NewCommand(commands.UpdateCommand, []string{}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for only one argument in update command",
			command:        commands.NewCommand(commands.UpdateCommand, []string{"123"}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			updteTask := NewUpdateTask(test.taskRepository)
			got, err := updteTask.parseCommand(*test.command)
			assertError(t, err, test.wantErr)
			assertTask(t, got, test.expected)
		})
	}
}
func TestUpdateTask_Execute(t *testing.T) {
	tests := []struct {
		name           string
		command        *commands.Command
		taskRepository tasks.ITaskRepository
		expectedResult *string
		wantErr        bool
	}{
		{
			name:           "shouls update a task succeessfully for an update command",
			command:        commands.NewCommand(commands.UpdateCommand, []string{"123", "new_value"}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expectedResult: func() *string { s := fmt.Sprintf(updateResult, 123); return &s }(),
			wantErr:        false,
		},
		{
			name:           "shouls update a task succeessfully for a mark in progress command",
			command:        commands.NewCommand(commands.MarkInProgressCommand, []string{"123"}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expectedResult: func() *string { s := fmt.Sprintf(updateResult, 123); return &s }(),
			wantErr:        false,
		},
		{
			name:           "shouls update a task succeessfully for a mark done command",
			command:        commands.NewCommand(commands.MarkDoneCommand, []string{"123"}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expectedResult: func() *string { s := fmt.Sprintf(updateResult, 123); return &s }(),
			wantErr:        false,
		},
		{
			name:           "should return an error when failing to parse the given command",
			command:        commands.NewCommand(commands.UpdateCommand, []string{}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error when repository fails to update",
			command:        commands.NewCommand(commands.UpdateCommand, []string{"123"}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			updateTask := NewUpdateTask(test.taskRepository)
			got, err := updateTask.Execute(*test.command)
			assertError(t, err, test.wantErr)
			assertResultString(t, got, test.expectedResult)
		})
	}
}
