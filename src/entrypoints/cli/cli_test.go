package cli

import (
	"tasktracker/src/commands"
	"tasktracker/src/tasks/ports"
	"tasktracker/src/tasks/ports/mocks"
	"testing"
)

func TestHandleCommand(t *testing.T) {
	successMockRepo := new(mocks.TaskRepositorySuccessfullMock)
	errorMockRepo := new(mocks.TaskRepositoryErrorMock)

	tests := []struct {
		name        string
		input       []string
		respository ports.ITaskRepository
		wantErr     bool
	}{
		{
			name:        "should run successfully for an 'add' command",
			input:       []string{"tasktracker", commands.AddCommand.String(), "Description"},
			respository: successMockRepo,
			wantErr:     false,
		},
		{
			name:        "should run successfully for an 'update' command",
			input:       []string{"tasktracker", commands.UpdateCommand.String(), "1", "Description"},
			respository: successMockRepo,
			wantErr:     false,
		},
		{
			name:        "should run successfully for a 'mark in progress' command",
			input:       []string{"tasktracker", commands.MarkInProgressCommand.String(), "1"},
			respository: successMockRepo,
			wantErr:     false,
		},
		{
			name:        "should run successfully for an 'mark done' command",
			input:       []string{"tasktracker", commands.MarkDoneCommand.String(), "1"},
			respository: successMockRepo,
			wantErr:     false,
		},
		{
			name:        "should run successfully for a 'delete' command",
			input:       []string{"tasktracker", commands.DeleteCommand.String(), "1"},
			respository: successMockRepo,
			wantErr:     false,
		},
		{
			name:        "should run successfully for a 'list' command",
			input:       []string{"tasktracker", "list"},
			respository: successMockRepo,
			wantErr:     false,
		},
		{
			name:        "should run successfully for a 'list todo' command",
			input:       []string{"tasktracker", "list", "todo"},
			respository: successMockRepo,
			wantErr:     false,
		},
		{
			name:        "should run successfully for a 'list in progress' command",
			input:       []string{"tasktracker", "list", "in-progress"},
			respository: successMockRepo,
			wantErr:     false,
		},
		{
			name:        "should run successfully for a 'list done' command",
			input:       []string{"tasktracker", "list", "done"},
			respository: successMockRepo,
			wantErr:     false,
		},
		{
			name:        "should return an error when no command is provided",
			input:       []string{"tasktracker"},
			respository: successMockRepo,
			wantErr:     true,
		},
		{
			name:        "should return an error when the given command is unknown",
			input:       []string{"tasktracker", "unknown"},
			respository: successMockRepo,
			wantErr:     true,
		},
		{
			name:        "should return an error when some error occurrs in the executed use case",
			input:       []string{"tasktracker", "list"},
			respository: errorMockRepo,
			wantErr:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := HandleCommand(test.input, test.respository)
			if (err != nil) != test.wantErr {
				t.Errorf("wantErr?: %t, got: %s", test.wantErr, err)
			}
		})
	}
}
