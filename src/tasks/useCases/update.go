package usecases

import (
	"fmt"
	"strconv"
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
)

const updateResult = "Task updated successfully (ID: %d)"

type UpdateTask struct {
	repository tasks.ITaskRepository
}

func NewUpdateTask(reposotory tasks.ITaskRepository) *UpdateTask {
	return &UpdateTask{
		repository: reposotory,
	}
}

func (u *UpdateTask) Execute(command commands.Command) (*string, error) {
	fmt.Println("executing update")
	task, err := u.parseArgs(command)
	if err != nil {
		return nil, fmt.Errorf("error while updating task: %s", err.Error())
	}
	updatedTask, err := u.repository.UpdateOne(*task)
	if err != nil {
		return nil, fmt.Errorf(taskUpdateFailed, err.Error())
	}
	result := fmt.Sprintf(updateResult, updatedTask.ID)
	return &result, nil
}

func (u *UpdateTask) parseArgs(command commands.Command) (*tasks.Task, error) {
	if command.Args() == nil || len(command.Args()) < 1 {
		return nil, errInvalidCommand
	}
	taskId, err := strconv.ParseUint(command.Args()[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error while trying to parse argument : %s", err.Error())
	}
	switch command.Name() {
	case commands.UpdateCommand:
		if len(command.Args()) < 2 {
			return nil, fmt.Errorf("invalid arguments")
		}
		return &tasks.Task{
			ID:          taskId,
			Description: command.Args()[1],
		}, nil
	case commands.MarkInProgressCommand:
		return &tasks.Task{
			ID:     taskId,
			Status: tasks.InProgress,
		}, nil
	case commands.MarkDoneCommand:
		return &tasks.Task{
			ID:     taskId,
			Status: tasks.Done,
		}, nil
	default:
		return nil, errInvalidCommand
	}
}
