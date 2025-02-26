package commands

import "testing"

func TestValidateUpdateArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "should return no error for valid args",
			args:    []string{"123", "new_value"},
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
			name:    "should return an error for only one argument",
			args:    []string{"123"},
			wantErr: true,
		},
		{
			name:    "should return an error for three args",
			args:    []string{"123", "new_value", "extra_arg"},
			wantErr: true,
		},
		{
			name:    "should return an error for non-numeric first arg",
			args:    []string{"abc", "new_value"},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := NewUpdateCommand(test.args)
			err := cmd.ParseArgs()
			assertError(t, err, test.wantErr)
		})
	}
}
