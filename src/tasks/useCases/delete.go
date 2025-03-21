package usecases

import (
	"fmt"
	"strconv"
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/ports"
)

var deleteResult = "Task deleted successfully"

type DeleteTask struct {
	repository ports.ITaskRepository
}

func NewDeleteTask(repository ports.ITaskRepository) *DeleteTask {
	return &DeleteTask{
		repository: repository,
	}
}

func (d *DeleteTask) Execute(command commands.Command) (*string, error) {
	fmt.Println("executing delete")
	task, err := d.parseArgs(command)
	if err != nil {
		return nil, fmt.Errorf(taskDeletionFailed, err.Error())
	}
	err = d.repository.DeleteOne(task.ID)
	if err != nil {
		return nil, fmt.Errorf(taskDeletionFailed, err.Error())
	}
	return &deleteResult, nil
}

func (d *DeleteTask) parseArgs(command commands.Command) (*tasks.Task, error) {
	if command.Args() == nil ||
		len(command.Args()) != 1 ||
		command.Name() != commands.DeleteCommand {
		return nil, errInvalidCommand
	}
	taskId, err := strconv.ParseUint(command.Args()[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error while trying to parse argument : %s", err.Error())
	}
	return &tasks.Task{ID: taskId}, nil
}
