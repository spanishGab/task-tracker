package usecases

import "errors"

var errInvalidCommand = errors.New("invalid command")

const (
	taskCreationFailed = "error while adding new task: %s"
	taskDeletionFailed = "error while deleting task: %s"
)
