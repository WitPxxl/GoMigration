package config

import "github.com/Witpxxl/GoMigration/command"

var flagExecDown = command.Flag{
	"down",
	"",
	"-down is used to revert a migration",
	"",
}

var flagExecUp = command.Flag{
	"up",
	"",
	"-up is used to apply a migration",
	"",
}

var configurationCommand = []command.Argument{
	{
		"exec",
		[]command.Flag{
			flagExecDown,
			flagExecUp,
		},
	},
}

func GetConfigurationCommand() []command.Argument {
	return configurationCommand
}
