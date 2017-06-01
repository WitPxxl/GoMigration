package command

import (
	"flag"
	"fmt"
	"os"
)

func NewCommand() *command {
	return &command{
		Flags: make(map[string]interface{}),
	}
}

type command struct {
	Argument Argument
	Flags    map[string]interface{}
}

func (cmd *command) initFlag() {
	currCommand := flag.NewFlagSet(cmd.Argument.Name, flag.ExitOnError)

	for key, fl := range cmd.Argument.ListFlag {
		currCommand.StringVar(&cmd.Argument.ListFlag[key].Value, fl.Name, fl.DefaultValue.(string), fl.HelpMessage)
	}

	currCommand.Parse(os.Args[2:])
}

func (cmd *command) CheckArgs(arguments []Argument) {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("You must used at least one argument")
		os.Exit(1)
	}

	find := false
	for _, value := range arguments {
		if value.Name == args[0] {
			find = true
			cmd.Argument = value
			break
		}
	}

	if find == false {
		fmt.Printf("Your command %s is not allowed\n", args[0])
		os.Exit(1)
	}

	cmd.initFlag()
}
