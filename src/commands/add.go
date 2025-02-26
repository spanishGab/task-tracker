package commands

import "fmt"

type Add struct {
	Command
}

func NewAddCommand(args []string) *Add {
	return &Add{
		Command: Command{name: AddCommand, args: args},
	}
}

func (a *Add) Execute() (*string, error) {
	fmt.Printf("executing add")
	if err := a.ParseArgs(); err != nil {
		return nil, fmt.Errorf("error while adding new task: %s", err.Error())
	}
	return nil, nil
}

func (a *Add) ParseArgs() error {
	if a.args == nil || len(a.args) != 1 {
		return ErrInvalidArgs
	}
	a.task.SetDescription(a.args[0])
	return nil
}
