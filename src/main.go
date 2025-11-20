package main

import (
	"fmt"
	"os"
	"path"
	"tasktracker/src/database"
	"tasktracker/src/entrypoints/cli"
	"tasktracker/src/tasks"
)

func main() {
	cwd := os.Getenv("CWD")
	fileHandler := database.NewFileHandler(path.Join(cwd, "src", "database", "tasks.json"))
	fileHandler.CreatFile()

	tasksRepository := tasks.NewTaskRepository(fileHandler.FileName, fileHandler)
	handler := cli.NewCommandHandler(tasksRepository)
	if err := handler.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
