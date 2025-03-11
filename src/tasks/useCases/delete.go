package usecases

import (
	"fmt"
	"strconv"
	"tasktracker/src/commands"
	"tasktracker/src/tasks/ports"
)

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
	taskId, err := d.parseArgs(command)
	if err != nil {
		return nil, fmt.Errorf(taskDeletionFailed, err.Error())
	}
	err = d.repository.DeleteOne(taskId)
	if err != nil {
		return nil, fmt.Errorf(taskDeletionFailed, err.Error())
	}
	return nil, nil
}

func (d *DeleteTask) parseArgs(command commands.Command) (uint64, error) {
	if command.Args() == nil || len(command.Args()) != 1 {
		return 0, errInvalidCommand
	}
	taskId, err := strconv.ParseUint(command.Args()[0], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error while trying to parse argument : %s", err.Error())
	}
	return taskId, nil
}
