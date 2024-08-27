package allCommands

import (
	"github.com/JuanMartinCoder/PokedexInGo/cmd/allcommands/exit"
	"github.com/JuanMartinCoder/PokedexInGo/cmd/allcommands/help"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type ListOfCommands map[string]cliCommand

func CreateCommandsList() *ListOfCommands {
	return &ListOfCommands{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    help.HelpCommand,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exit.CommandExit,
		},
	}
}

func (c *ListOfCommands) IsACommand(name string) bool {
	if _, ok := (*c)[name]; ok {
		return true
	}
	return false
}

func (c *ListOfCommands) ExcecuteCmd(name string) func() error {
	return (*c)[name].callback
}
