package cli

import (
	"fmt"
	"tasktracker/src/commands"
)

var createInvalidCommandError = func(command string, message string) error {
	return fmt.Errorf("invalid command %s - %s", command, message)

}

func ReadCommand(input []string) error {
	inputLength := len(input)
	if inputLength <= 1 {
		return commands.ErrInvalidArgs
	}
	commandName := input[1]

	var command commands.ICommand
	switch commandName {
	case commands.AddCommand.String():
		if inputLength != 3 {
			return commands.ErrInvalidArgs
		}
		add := commands.NewAddCommand(input[2:])
		command = add
	case commands.UpdateCommand.String():
		if inputLength != 4 {
			return commands.ErrInvalidArgs
		}
		update := commands.NewUpdateCommand(input[2:])
		command = update
	case commands.DeleteCommand.String():
		if inputLength != 3 {
			return commands.ErrInvalidArgs
		}
		delete := commands.NewDeleteCommand(input[2:])
		command = delete
	default:
		return createInvalidCommandError(commandName, "unknown command")
	}
	result, err := command.Execute()
	if err != nil {
		return createInvalidCommandError(commandName, err.Error())
	}
	if result != nil {
		fmt.Println(*result)
	}
	return nil
}
