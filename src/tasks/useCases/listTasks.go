package usecases

import (
	"fmt"
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/ports"
)

type ListTask struct {
	repository ports.ITaskRepository
}

func NewListTask(repository ports.ITaskRepository) *ListTask {
	return &ListTask{
		repository: repository,
	}
}

func (a *ListTask) Execute(command commands.Command) (*string, error) {
	fmt.Println("listing tasks")
	task, err := a.parseArgs(command)
	if err != nil {
		return nil, fmt.Errorf(tasksListingFailed, err.Error())
	}

	var tasks []tasks.Task
	var formattedResult string
	if task.Status == "" {
		tasks, err = a.repository.GetAllTasks()
		if err != nil {
			return nil, fmt.Errorf(tasksListingFailed, err.Error())
		}
	}
	// TODO: finish status listing and add unit-tests
	result, err := a.repository.TasksToBytes(tasks)
	formattedResult = string(result)
	if err != nil {
		return nil, fmt.Errorf(tasksListingFailed, err.Error())
	}
	return &formattedResult, nil
}

func (a *ListTask) parseArgs(command commands.Command) (*tasks.Task, error) {
	if command.Args() == nil || command.Name() != commands.ListCommand {
		return nil, errInvalidCommand
	}

	if len(command.Args()) <= 0 {
		return &tasks.Task{}, nil
	}

	switch command.Args()[0] {
	case "done":
		return &tasks.Task{Status: tasks.Done}, nil
	case "todo":
		return &tasks.Task{Status: tasks.Todo}, nil
	case "in-progress":
		return &tasks.Task{Status: tasks.InProgress}, nil
	default:
		return nil, fmt.Errorf("invalid status: %s", command.Args()[0])
	}
}
