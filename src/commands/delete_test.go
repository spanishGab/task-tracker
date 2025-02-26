package commands

import "testing"

func TestValidateDeleteArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "should return no error for valid args",
			args:    []string{"123"},
			wantErr: false,
		},
		{
			name:    "should return an error for no args",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "should return an error for nil args",
			args:    nil,
			wantErr: true,
		},
		{
			name:    "should return an error for multiple args",
			args:    []string{"123", "extra_arg"},
			wantErr: true,
		},
		{
			name:    "should return an error for non-numeric id",
			args:    []string{"abc"},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := NewDeleteCommand(test.args)
			err := cmd.ParseArgs()
			assertError(t, err, test.wantErr)
		})
	}
}
