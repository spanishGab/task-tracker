package ports

import "tasktracker/src/commands"

type IUseCase interface {
	Execute(command commands.Command) (*string, error)
}
