package entrypoint

import (
	"errors"
)

var errInvlidArgs = errors.New("invalid arguments")

type Args string

// CLI args
const (
	add    Args = "add"
	update Args = "update"
	delete      = "delete"
	list        = "list"
)

func ReadUserInput(input []string) error {
	if len(input) <= 1 {
		return errInvlidArgs
	}
	command := input[1]

	switch command {
	case string(add):
		if len(input) <= 2 {
			return errInvlidArgs
		}
		// TODO: finalizar
	}
}
