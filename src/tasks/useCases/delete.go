package usecases

import (
	"fmt"
	"strconv"
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
)

var deleteResult = "Task deleted successfully"

type DeleteTask struct {
	repository tasks.ITaskRepository
}

func NewDeleteTask(repository tasks.ITaskRepository) *DeleteTask {
	return &DeleteTask{
		repository: repository,
	}
}

func (d *DeleteTask) Execute(command commands.Command) (*string, error) {
	fmt.Println("executing delete")
	id, err := d.parseArgs(command)
	if err != nil {
		return nil, fmt.Errorf(taskDeletionFailed, err.Error())
	}
	err = d.repository.DeleteOne(id)
	if err != nil {
		return nil, fmt.Errorf(taskDeletionFailed, err.Error())
	}
	return &deleteResult, nil
}

func (d *DeleteTask) parseArgs(command commands.Command) (uint64, error) {
	if command.Args() == nil ||
		len(command.Args()) != 1 ||
		command.Name() != commands.DeleteCommand {
		return 0, errInvalidCommand
	}
	taskId, err := strconv.ParseUint(command.Args()[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error while trying to parse argument : %s", err.Error())
	}
	return taskId, nil
}
