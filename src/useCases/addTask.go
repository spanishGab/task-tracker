package usecases

import (
	"errors"
	"fmt"
	"tasktracker/src/commands"
	"tasktracker/src/ports"
	"tasktracker/src/tasks"
)

var ErrInvalidCommand = errors.New("invalid command")

const (
	taskCreationFailed = "error while adding new task: %s"
)

type AddTask struct {
	repository ports.ITaskRepository
}

func NewAddTask(repository ports.ITaskRepository) *AddTask {
	return &AddTask{
		repository: repository,
	}
}

func (a *AddTask) Execute(command commands.Command) (*string, error) {
	fmt.Println("executing add")
	if err := a.parseArgs(command); err != nil {
		return nil, fmt.Errorf(taskCreationFailed, err.Error())
	}
	newTask, err := a.repository.CreateOne(*tasks.NewTask(command.Args()[0]))
	if err != nil {
		return nil, fmt.Errorf(taskCreationFailed, err.Error())
	}
	result := fmt.Sprintf("Task addeed successfully (ID: %d)", newTask.ID)
	return &result, nil
}

func (a *AddTask) parseArgs(command commands.Command) error {
	if command.Args() == nil || len(command.Args()) != 1 {
		return ErrInvalidCommand
	}
	return nil
}
