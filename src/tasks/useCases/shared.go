package usecases

import "errors"

var errInvalidCommand = errors.New("invalid command")

const (
	taskCreationFailed = "error while adding new task: %s"
	taskDeletionFailed = "error while deleting task: %s"
	taskUpdateFailed   = "error while updating task: %s"
	tasksListingFailed = "error while listing tasks: %s"
)
