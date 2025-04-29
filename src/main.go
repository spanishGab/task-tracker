package main

import (
	"fmt"
	"os"
	"path"
	"tasktracker/src/commands"
	"tasktracker/src/database"
	"tasktracker/src/entrypoints/cli"
	"tasktracker/src/tasks"
)

func main() {
	inputLength := len(os.Args)
	if inputLength <= 1 {
		fmt.Println(commands.ErrInvalidArgs)
		return
	}
	cwd := os.Getenv("CWD")
	fileHandler := database.NewFileHandler(path.Join(cwd, "src", "database", "tasks.json"))
	fileHandler.CreatFile()

	tasksRepository := tasks.NewTaskRepository(fileHandler.FileName, fileHandler)
	if err := cli.HandleCommand(os.Args, tasksRepository); err != nil {
		fmt.Println(err)
	}
}
