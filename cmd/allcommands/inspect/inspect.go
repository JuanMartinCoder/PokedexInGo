package inspect

import (
	"fmt"

	"github.com/JuanMartinCoder/PokedexInGo/api"
)

func InsepectCommand(cfg *api.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Explore command takes 1 argument <pokemon-name>")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.PokeClient.GetPokemon(pokemonName)
	if !ok {
		return fmt.Errorf("You have not caught that pokemons yet", pokemonName)
	}

	fmt.Println("Name: ", pokemon.Name)
	fmt.Println("Height: ", pokemon.Height)
	fmt.Println("Weight: ", pokemon.Weight)
	fmt.Println("Stats:")
	for i := range pokemon.Stats {
		fmt.Printf("  - %s: %v\n", pokemon.Stats[i].Stat.Name, pokemon.Stats[i].BaseStat)
	}
	fmt.Println("Types:")
	for i := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemon.Types[i].Type.Name)
	}

	return nil
}
