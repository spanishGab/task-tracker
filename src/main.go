package main

import (
	"fmt"
	"os"
	"tasktracker/src/commands"
	"tasktracker/src/entrypoints/cli"
)

func main() {
	inputLength := len(os.Args)
	if inputLength <= 1 {
		fmt.Println(commands.ErrInvalidArgs)
		return
	}
	if err := cli.ReadCommand(os.Args); err != nil {
		fmt.Println(err)
	}
}
