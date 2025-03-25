package usecases

import (
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/mocks"
	"testing"
)

func TestListTask_parseCommand(t *testing.T) {
	done := tasks.Done
	todo := tasks.Todo
	inProgress := tasks.InProgress
	tests := []struct {
		name     string
		cmdName  commands.CommandName
		args     []string
		expected *tasks.Status
		wantErr  bool
	}{
		{
			name:     "should return no error for valid args with status 'done'",
			cmdName:  commands.ListCommand,
			args:     []string{tasks.Done.String()},
			expected: &done,
			wantErr:  false,
		},
		{
			name:     "should return no error for valid args with status 'todo'",
			cmdName:  commands.ListCommand,
			args:     []string{tasks.Todo.String()},
			expected: &todo,
			wantErr:  false,
		},
		{
			name:     "should return no error for valid args with status 'in-progress'",
			cmdName:  commands.ListCommand,
			args:     []string{tasks.InProgress.String()},
			expected: &inProgress,
			wantErr:  false,
		},
		{
			name:     "should return no error for no status args",
			cmdName:  commands.ListCommand,
			args:     []string{},
			expected: nil,
			wantErr:  false,
		},
		{
			name:     "should return an error for invalid status",
			cmdName:  commands.ListCommand,
			args:     []string{"invalid_status"},
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "should return an error for invalid command name",
			cmdName:  "invalid_command",
			args:     []string{tasks.Done.String()},
			expected: nil,
			wantErr:  true,
		},
		{
			name:     "should return an error for nil args",
			cmdName:  commands.ListCommand,
			args:     nil,
			expected: nil,
			wantErr:  true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := commands.NewCommand(test.cmdName, test.args)
			listTask := NewListTask(nil)
			got, err := listTask.parseCommand(*cmd)
			if (err != nil) != test.wantErr {
				t.Errorf("error = %v, wantErr? %v", err, test.wantErr)
				return
			}
			if test.expected != nil {
				if got == nil {
					t.Errorf("got a nil value, expected = %v", *test.expected)
				} else if *got != *test.expected {
					t.Errorf("got = %v, expected = %v", *got, *test.expected)
				}
			} else if got != nil {
				t.Errorf("got = %v, expected a nil value", got)
			}
		})
	}
}

func TestListTask_Execute(t *testing.T) {
	emptyResult := emptyTasksResult

	tests := []struct {
		name           string
		command        *commands.Command
		taskRepository tasks.ITaskRepository
		expectedResult *string
		wantErr        bool
	}{
		{
			name:           "should list all tasks succeessfully for a list command",
			command:        commands.NewCommand(commands.ListCommand, []string{}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expectedResult: &emptyResult,
			wantErr:        false,
		},
		{
			name:           "shouls list all tasks succeessfully for a list done command",
			command:        commands.NewCommand(commands.ListCommand, []string{tasks.Done.String()}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expectedResult: &emptyResult,
			wantErr:        false,
		},
		{
			name:           "shouls list all tasks succeessfully for a list todo command",
			command:        commands.NewCommand(commands.ListCommand, []string{tasks.Todo.String()}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expectedResult: &emptyResult,
			wantErr:        false,
		},
		{
			name:           "shouls list all tasks succeessfully for a list in-progress command",
			command:        commands.NewCommand(commands.ListCommand, []string{tasks.InProgress.String()}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			expectedResult: &emptyResult,
			wantErr:        false,
		},
		{
			name:           "should return an error when failing to parse the given command",
			command:        commands.NewCommand(commands.CommandName("lis"), []string{}),
			taskRepository: &mocks.TaskRepositorySuccessfullMock{},
			wantErr:        true,
		},
		{
			name:           "should return an error when repository fails to update",
			command:        commands.NewCommand(commands.ListCommand, []string{"123"}),
			taskRepository: &mocks.TaskRepositoryFailureMock{},
			wantErr:        true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			updateTask := NewListTask(test.taskRepository)
			got, err := updateTask.Execute(*test.command)
			assertError(t, err, test.wantErr)
			assertResultString(t, got, test.expectedResult)
		})
	}
}
