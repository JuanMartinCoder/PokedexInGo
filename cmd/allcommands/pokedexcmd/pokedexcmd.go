package pokedexcmd

import (
	"fmt"

	"github.com/JuanMartinCoder/PokedexInGo/api"
)

func PokedexCommand(cfg *api.Config, args ...string) error {
	pokedex := cfg.PokeClient.GetPokedex()

	fmt.Println("Pokedex:")

	if len(pokedex) == 0 {
		fmt.Println("No pokemons found")
		return nil
	}

	for _, pokemon := range pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}
	return nil
}
