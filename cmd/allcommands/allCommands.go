package allCommands

import (
	"fmt"

	"github.com/JuanMartinCoder/PokedexInGo/api"
	"github.com/JuanMartinCoder/PokedexInGo/cmd/allcommands/catch"
	"github.com/JuanMartinCoder/PokedexInGo/cmd/allcommands/exit"
	"github.com/JuanMartinCoder/PokedexInGo/cmd/allcommands/explore"
	"github.com/JuanMartinCoder/PokedexInGo/cmd/allcommands/inspect"
	"github.com/JuanMartinCoder/PokedexInGo/cmd/allcommands/mapb"
	"github.com/JuanMartinCoder/PokedexInGo/cmd/allcommands/mapcmd"
	"github.com/JuanMartinCoder/PokedexInGo/cmd/allcommands/pokedexcmd"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*api.Config, ...string) error
}

type ListOfCommands map[string]cliCommand

func (c *ListOfCommands) IsACommand(name string) bool {
	if _, ok := (*c)[name]; ok {
		return true
	}
	return false
}

func (c *ListOfCommands) ExcecuteCmd(name string) func(cfg *api.Config, args ...string) error {
	return (*c)[name].callback
}

func CreateCommandsList() *ListOfCommands {
	return &ListOfCommands{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCmd,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exit.CommandExit,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 location areas in the Pokemon world",
			callback:    mapcmd.MapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Similar to the map command, however, instead of displaying the next 20 locations,\nit displays the previous 20 locations.",
			callback:    mapb.MapbCommand,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "list of all the Pokémon in a given area",
			callback:    explore.Explore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Tries to catch a pokemon and add it to your pokedex",
			callback:    catch.CatchCommand,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Print a list of all the names of the Pokemons in the Pokedex",
			callback:    pokedexcmd.PokedexCommand,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Allow players to see details about a Pokemon if they have seen it before. (caught it)",
			callback:    inspect.InsepectCommand,
		},
	}
}

func helpCmd(cfg *api.Config, args ...string) error {
	lista := CreateCommandsList()
	fmt.Println("Welcome to Pokedex!\n\nThese are the Available commands:\n")
	for _, value := range *lista {
		fmt.Printf(" - %s: %s\n", value.name, value.description)
	}
	fmt.Println("\n")
	return nil
}
