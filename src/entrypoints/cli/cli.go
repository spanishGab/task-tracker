package cli

import (
	"fmt"
	"os"
	"path"
	"tasktracker/src/commands"
	"tasktracker/src/database"
	"tasktracker/src/ports"
	"tasktracker/src/tasks"
	usecases "tasktracker/src/useCases"
)

var createInvalidCommandError = func(command string, message string) error {
	return fmt.Errorf("invalid command %s - %s", command, message)

}

func ReadCommand(input []string) error {
	cwd, _ := os.Getwd()
	fileHandler := database.NewFileHandler(path.Join(cwd, "..", "db", "tasks.json"))
	if err := fileHandler.Open(); err != nil {
		return fmt.Errorf("failed to open file %s: %s", fileHandler.FileName, err.Error())
	}
	defer fileHandler.Close()
	tasksRepository := tasks.NewTaskRepository(fileHandler.FileName, fileHandler)

	inputLength := len(input)
	commandName := input[1]

	var useCase ports.IUseCase
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
	// case commands.DeleteCommand.String():
	// 	if inputLength != 3 {
	// 		return commands.ErrInvalidArgs
	// 	}
	// 	delete := commands.NewDeleteCommand(input[2:])
	// 	useCase = delete
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
