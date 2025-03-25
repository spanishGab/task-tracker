package cli

import (
	"fmt"
	"tasktracker/src/commands"
	"tasktracker/src/contracts"
	"tasktracker/src/tasks"
	usecases "tasktracker/src/tasks/useCases"
)

var createInvalidCommandError = func(command string, message string) error {
	return fmt.Errorf("invalid command '%s': %s", command, message)

}

func HandleCommand(input []string, tasksRepository tasks.ITaskRepository) error {
	inputLength := len(input)

	var useCase contracts.IUseCase
	var command commands.Command
	if inputLength < 2 {
		return commands.ErrInvalidArgs
	}

	commandName := input[1]
	switch commandName {
	case commands.AddCommand.String():
		command = *commands.NewCommand(commands.AddCommand, input[2:])
		useCase = usecases.NewAddTask(tasksRepository)
	case commands.UpdateCommand.String(), commands.MarkDoneCommand.String(), commands.MarkInProgressCommand.String():
		command = *commands.NewCommand(commands.CommandName(commandName), input[2:])
		useCase = usecases.NewUpdateTask(tasksRepository)
	case commands.DeleteCommand.String():
		command = *commands.NewCommand(commands.DeleteCommand, input[2:])
		useCase = usecases.NewDeleteTask(tasksRepository)
	case commands.ListCommand.String():
		command = *commands.NewCommand(commands.ListCommand, input[2:])
		useCase = usecases.NewListTask(tasksRepository)
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
