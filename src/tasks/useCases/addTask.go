package usecases

import (
	"fmt"
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
)

const addResult = "Task addeed successfully (ID: %d)"

type AddTask struct {
	repository tasks.ITaskRepository
}

func NewAddTask(repository tasks.ITaskRepository) *AddTask {
	return &AddTask{
		repository: repository,
	}
}

func (a *AddTask) Execute(command commands.Command) (*string, error) {
	fmt.Println("executing add")
	task, err := a.parseCommand(command)
	if err != nil {
		return nil, fmt.Errorf(taskCreationFailed, err.Error())
	}
	newTask, err := a.repository.CreateOne(*task)
	if err != nil {
		return nil, fmt.Errorf(taskCreationFailed, err.Error())
	}
	result := fmt.Sprintf(addResult, newTask.ID)
	return &result, nil
}

func (a *AddTask) parseCommand(command commands.Command) (*tasks.Task, error) {
	if command.Args() == nil ||
		len(command.Args()) != 1 ||
		command.Name() != commands.AddCommand {
		return nil, errInvalidCommand
	}
	return &tasks.Task{
		Description: command.Args()[0],
		Status:      tasks.Todo,
	}, nil
}
