package usecases

import (
	"fmt"
	"strconv"
	"tasktracker/src/commands"
	"tasktracker/src/tasks"
	"tasktracker/src/tasks/ports"
)

const updateResult = "Task updated successfully (ID: %d)"

type UpdateTask struct {
	repository ports.ITaskRepository
}

func NewUpdateTask(reposotory ports.ITaskRepository) *UpdateTask {
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
	updatedTask, _ := u.repository.UpdateOne(*task)
	result := fmt.Sprintf(updateResult, updatedTask.ID)
	return &result, nil
}

func (u *UpdateTask) parseArgs(command commands.Command) (*tasks.Task, error) {
	if command.Args() == nil || len(command.Args()) != 2 {
		return nil, errInvalidCommand
	}
	taskId, err := strconv.ParseUint(command.Args()[0], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error while trying to parse argument : %s", err.Error())
	}
	return tasks.NewTaskWithId(taskId, command.Args()[1]), nil
}
