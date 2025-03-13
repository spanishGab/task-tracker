package usecases

import (
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/ports"
	"tasktracker/src/tasks/useCases/mocks"
	"testing"
)

func TestDeleteTask_parseArgs(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		taskRepository ports.ITaskRepository
		expected       *tasks.Task
		wantErr        bool
	}{
		{
			name:           "should return no error for valid args",
			args:           []string{"123"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expected:       &tasks.Task{ID: 123},
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
			args:           []string{"123", "extra_arg"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for non-numeric id",
			args:           []string{"abc"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := commands.NewCommand("delete", test.args)
			deleteTask := NewDeleteTask(test.taskRepository)
			got, err := deleteTask.parseArgs(*cmd)
			assertError(t, err, test.wantErr)
			assertTask(t, got, test.expected)
		})
	}
}
