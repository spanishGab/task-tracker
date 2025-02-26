package commands

import (
	"fmt"
	"strconv"
)

type Delete struct {
	Command
}

func NewDeleteCommand(args []string) *Delete {
	return &Delete{
		Command: Command{name: DeleteCommand, args: args},
	}
}

func (d *Delete) Execute() (*string, error) {
	fmt.Printf("executing delete")
	if err := d.ParseArgs(); err != nil {
		return nil, fmt.Errorf("error while deleting task: %s", err.Error())
	}
	return nil, nil
}

func (d *Delete) ParseArgs() error {
	if d.args == nil || len(d.args) != 1 {
		return ErrInvalidArgs
	}
	_, err := strconv.ParseUint(d.args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("error while trying to parse argument : %s", err.Error())
	}
	return nil
}
