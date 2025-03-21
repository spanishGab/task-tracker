package commands

import (
	"errors"
)

var ErrInvalidArgs = errors.New("invalid arguments")

type CommandName string

const (
	AddCommand            CommandName = "add"
	UpdateCommand         CommandName = "update"
	DeleteCommand         CommandName = "delete"
	MarkInProgressCommand CommandName = "mark-in-progress"
	MarkDoneCommand       CommandName = "mark-done"
	ListCommand           CommandName = "list"
)

func (cn CommandName) String() string {
	return string(cn)
}

type Command struct {
	name CommandName
	args []string
}

func NewCommand(name CommandName, args []string) *Command {
	return &Command{
		name: name,
		args: args,
	}
}

func (c *Command) Name() CommandName {
	return c.name
}

func (c *Command) Args() []string {
	return c.args
}
