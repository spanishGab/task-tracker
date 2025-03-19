package usecases

import (
	"fmt"
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/ports"
)

const addResult = "Task addeed successfully (ID: %d)"

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
	task, err := a.parseArgs(command)
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

func (a *AddTask) parseArgs(command commands.Command) (*tasks.Task, error) {
	if command.Args() == nil || len(command.Args()) != 1 {
		return nil, errInvalidCommand
	}
	return &tasks.Task{
		Description: command.Args()[0],
		Status:      tasks.Done,
	}, nil
}
