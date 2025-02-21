package cli

import "fmt"

func GetCommandFrom(input []string) (*Command, error) {
	inputLength := len(input)
	if inputLength <= 1 {
		return nil, errInvalidArgs
	}
	commandName := input[1]

	var err error
	var command *Command
	switch commandName {
	case Add.String():
		if inputLength != 3 {
			return nil, errInvalidArgs
		}
		command = NewAddCommand()
		command.SetArgs(input[2:])
		err = command.ValidateAddArgs()
	case Update.String():
		if inputLength != 4 {
			return nil, errInvalidArgs
		}
		command = NewUpdateCommand()
		command.SetArgs(input[2:])
		err = command.ValidateUpdateArgs()
	case Delete.String():
		if inputLength != 3 {
			return nil, errInvalidArgs
		}
		command = NewDeleteCommand()
		command.SetArgs(input[2:])
		err = command.ValidateDeleteArgs()
	}
	if err != nil {
		return nil, fmt.Errorf("invalid command entered : %s", err.Error())
	}
	return command, nil
}
