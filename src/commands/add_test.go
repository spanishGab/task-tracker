package commands

import "testing"

func TestValidateAddArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{
			name:    "should return no error for valid args",
			args:    []string{"arg1"},
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
			args:    []string{"arg1", "arg2"},
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cmd := NewAddCommand(test.args)
			err := cmd.ParseArgs()
			assertError(t, err, test.wantErr)
		})
	}
}
