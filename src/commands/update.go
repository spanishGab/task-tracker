package commands

import (
	"fmt"
	"strconv"
)

type Update struct {
	Command
}

func NewUpdateCommand(args []string) *Update {
	return &Update{
		Command: Command{name: UpdateCommand, args: args},
	}
}

func (u *Update) Execute() (*string, error) {
	fmt.Println("executing update")
	if err := u.ParseArgs(); err != nil {
		return nil, fmt.Errorf("error while updating task: %s", err.Error())
	}
	return nil, nil
}

func (u *Update) ParseArgs() error {
	if u.args == nil || len(u.args) < 2 {
		return ErrInvalidArgs
	}
	_, err := strconv.ParseUint(u.args[0], 10, 64)
	if err != nil {
		return fmt.Errorf("error while trying to parse argument : %s", err.Error())
	}
	return nil
}
