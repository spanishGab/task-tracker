package usecases

import (
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/ports"
	"tasktracker/src/tasks/ports/mocks"
	"testing"
)

func TestDeleteTask_parseArgs(t *testing.T) {
	tests := []struct {
		name           string
		cmdName        commands.CommandName
		args           []string
		taskRepository ports.ITaskRepository
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
			got, err := deleteTask.parseArgs(*cmd)
			assertError(t, err, test.wantErr)
			assertTask(t, &tasks.Task{ID: got}, &tasks.Task{ID: test.expected})
		})
	}
}
