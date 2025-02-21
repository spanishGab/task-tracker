package cli

import (
	"errors"
	"fmt"
	"strconv"
)

var errInvalidArgs = errors.New("invalid arguments")

type CommandName string

const (
	Add    CommandName = "add"
	Update CommandName = "update"
	Delete CommandName = "delete"
)

func (cn CommandName) String() string {
	return string(cn)
}

type Command struct {
	name string
	args []string
}

func NewAddCommand() *Command {
	return &Command{name: "add"}
}

func NewUpdateCommand() *Command {
	return &Command{name: "update"}
}

func NewDeleteCommand() *Command {
	return &Command{name: "delete"}
}

func (c *Command) Name() string {
	return c.name
}

func (c *Command) SetArgs(args []string) {
	c.args = args
}

func (c *Command) ValidateAddArgs() error {
	if c.args == nil || len(c.args) != 1 {
		return errInvalidArgs
	}
	return nil
}

func (c *Command) ValidateUpdateArgs() error {
	if c.args == nil || len(c.args) != 2 {
		return errInvalidArgs
	}
	_, err := strconv.ParseUint(c.args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("error while trying to parse argument : %s", err.Error())
	}
	return nil
}

func (c *Command) ValidateDeleteArgs() error {
	if c.args == nil || len(c.args) != 1 {
		return errInvalidArgs
	}
	_, err := strconv.ParseUint(c.args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("error while trying to parse argument : %s", err.Error())
	}
	return nil
}
