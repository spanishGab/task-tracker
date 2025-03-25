package usecases

import (
	"fmt"
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
)

const (
	emptyTasksResult = "No tasks to list"
)

type ListTask struct {
	repository tasks.ITaskRepository
}

func NewListTask(repository tasks.ITaskRepository) *ListTask {
	return &ListTask{
		repository: repository,
	}
}

func (lt *ListTask) Execute(command commands.Command) (*string, error) {
	fmt.Println("listing tasks")
	status, err := lt.parseCommand(command)
	if err != nil {
		return nil, fmt.Errorf(tasksListingFailed, err.Error())
	}

	var tasks []tasks.Task
	if status == nil {
		tasks, err = lt.repository.GetAllTasks()
		if err != nil {
			return nil, fmt.Errorf(tasksListingFailed, err.Error())
		}
	} else {
		tasks, err = lt.repository.GetAllTasksByStatus(*status)
		if err != nil {
			return nil, fmt.Errorf(tasksListingFailed, err.Error())
		}
	}
	result, err := lt.repository.Format(tasks)

	formattedResult := emptyTasksResult
	if len(result) > 0 {
		formattedResult = string(result)
	}
	if err != nil {
		return nil, fmt.Errorf(tasksListingFailed, err.Error())
	}
	return &formattedResult, nil
}

func (lt *ListTask) parseCommand(command commands.Command) (*tasks.Status, error) {
	if command.Args() == nil || command.Name() != commands.ListCommand {
		return nil, errInvalidCommand
	}

	if len(command.Args()) <= 0 {
		return nil, nil
	}

	var status tasks.Status
	switch command.Args()[0] {
	case tasks.Done.String():
		status = tasks.Done
	case tasks.Todo.String():
		status = tasks.Todo
	case tasks.InProgress.String():
		status = tasks.InProgress
	default:
		return nil, fmt.Errorf("invalid status '%s'", command.Args()[0])
	}
	return &status, nil
}
