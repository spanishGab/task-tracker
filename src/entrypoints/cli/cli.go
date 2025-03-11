package cli

import (
	"fmt"
	"tasktracker/src/commands"
	generalports "tasktracker/src/ports"
	taskports "tasktracker/src/tasks/ports"
	usecases "tasktracker/src/tasks/useCases"
)

var createInvalidCommandError = func(command string, message string) error {
	return fmt.Errorf("invalid command %s - %s", command, message)

}

func ReadCommand(input []string, tasksRepository taskports.ITaskRepository) error {
	inputLength := len(input)
	commandName := input[1]

	var useCase generalports.IUseCase
	var command commands.Command
	switch commandName {
	case commands.AddCommand.String():
		if inputLength != 3 {
			return commands.ErrInvalidArgs
		}
		command = *commands.NewCommand(commands.AddCommand, input[2:])
		useCase = usecases.NewAddTask(tasksRepository)
	// case commands.UpdateCommand.String():
	// 	if inputLength != 4 {
	// 		return commands.ErrInvalidArgs
	// 	}
	// 	update := commands.NewUpdateCommand(input[2:])
	// 	useCase = update
	case commands.DeleteCommand.String():
		if inputLength != 3 {
			return commands.ErrInvalidArgs
		}
		command = *commands.NewCommand(commands.DeleteCommand, input[2:])
		useCase = usecases.NewDeleteTask(tasksRepository)
	default:
		return createInvalidCommandError(commandName, "unknown command")
	}
	result, err := useCase.Execute(command)
	if err != nil {
		return createInvalidCommandError(commandName, err.Error())
	}
	if result != nil {
		fmt.Println(*result)
	}
	return nil
}
