package usecases

import (
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/ports"
	"tasktracker/src/tasks/useCases/mocks"
	"testing"
)

func TestUpdateTask_parseArgs(t *testing.T) {
	tests := []struct {
		name           string
		args           []string
		taskRepository ports.ITaskRepository
		expected       *tasks.Task
		wantErr        bool
	}{
		{
			name:           "should return no error for valid args",
			args:           []string{"123", "new_value"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expected:       tasks.NewTaskWithId(123, "new_value"),
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
			name:           "should return an error for only one argument",
			args:           []string{"123"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for three args",
			args:           []string{"123", "new_value", "extra_arg"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error for non-numeric first arg",
			args:           []string{"abc", "new_value"},
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := commands.NewCommand("update", test.args)
			updteTask := NewUpdateTask(test.taskRepository)
			got, err := updteTask.parseArgs(*cmd)
			assertError(t, err, test.wantErr)
			assertTask(t, got, test.expected)
		})
	}
}
