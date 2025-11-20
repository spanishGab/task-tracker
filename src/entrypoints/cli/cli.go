package cli

import (
	"fmt"
	"tasktracker/src/commands"
	"tasktracker/src/contracts"
	"tasktracker/src/tasks"
	usecases "tasktracker/src/tasks/useCases"
)

var createInvalidCommandError = func(command commands.CommandName, message string) error {
	return fmt.Errorf("command error '%s': %s", command, message)

}

type commandHandler struct {
	tasksRepository tasks.ITaskRepository
}

func NewCommandHandler(tasksRepository tasks.ITaskRepository) *commandHandler {
	return &commandHandler{tasksRepository: tasksRepository}
}

func (ch *commandHandler) Run(input []string) error {
	if inputLength := len(input); inputLength < 2 {
		return commands.ErrInvalidArgs
	}

	commandName := commands.CommandName(input[1])
	useCase, err := ch.chooseUseCase(commandName)
	if err != nil {
		return err
	}

	command := *commands.NewCommand(commandName, input[2:])
	result, err := useCase.Execute(command)
	if err != nil {
		return createInvalidCommandError(commandName, err.Error())
	}
	if result != nil {
		fmt.Println(*result)
	}
	return nil
}

func (ch *commandHandler) chooseUseCase(commandName commands.CommandName) (contracts.IUseCase, error) {
	var useCase contracts.IUseCase
	switch commandName {
	case commands.AddCommand:
		useCase = usecases.NewAddTask(ch.tasksRepository)
	case commands.UpdateCommand, commands.MarkDoneCommand, commands.MarkInProgressCommand:
		useCase = usecases.NewUpdateTask(ch.tasksRepository)
	case commands.DeleteCommand:
		useCase = usecases.NewDeleteTask(ch.tasksRepository)
	case commands.ListCommand:
		useCase = usecases.NewListTask(ch.tasksRepository)
	default:
		return nil, createInvalidCommandError(commandName, "unknown command")
	}
	return useCase, nil
}
