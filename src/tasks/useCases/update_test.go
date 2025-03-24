package usecases

import (
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/ports"
	"tasktracker/src/tasks/ports/mocks"
	"testing"
)

func TestUpdateTask_parseArgs(t *testing.T) {
	tests := []struct {
		name           string
		command        *commands.Command
		taskRepository ports.ITaskRepository
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
			got, err := updteTask.parseArgs(*test.command)
			assertError(t, err, test.wantErr)
			assertTask(t, got, test.expected)
		})
	}
}
