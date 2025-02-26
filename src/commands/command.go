package commands

import (
	"errors"
	"tasktracker/src/tasks"
)

var ErrInvalidArgs = errors.New("invalid arguments")

type CommandName string

const (
	AddCommand    CommandName = "add"
	UpdateCommand CommandName = "update"
	DeleteCommand CommandName = "delete"
)

func (cn CommandName) String() string {
	return string(cn)
}

type ICommand interface {
	Execute() (*string, error)
}

type Command struct {
	name CommandName
	args []string
	task tasks.Task
}

func NewCommand(name CommandName, args []string, task tasks.Task) *Command {
	return &Command{
		name: name,
		args: args,
		task: task,
	}
}

func (c *Command) Name() CommandName {
	return c.name
}

func (c *Command) Args() []string {
	return c.args
}

func (c *Command) Task() tasks.Task {
	return c.task
}
