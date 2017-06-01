package main

import (
	"fmt"

	"github.com/Witpxxl/GoMigration/command"
	"github.com/Witpxxl/GoMigration/config"
)

var arguments = config.GetConfigurationCommand()

var cmd = command.NewCommand()

func main() {
	cmd.CheckArgs(arguments)

	fmt.Println("You just called the", cmd.Argument.Name, "function")
	fmt.Println("With args: ")

	for _, f := range cmd.Argument.ListFlag {
		println("\t-> ", f.Name, "=", f.Value)
	}
}
