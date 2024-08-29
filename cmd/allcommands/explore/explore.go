package explore

import (
	"fmt"

	"github.com/JuanMartinCoder/PokedexInGo/api"
)

func Explore(cfg *api.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Explore command takes 1 argument <area-name>")
	}
	locationArea := args[0]

	data, err := cfg.PokeClient.GetLocationArea(locationArea)
	if err != nil {
		return err
	}

	fmt.Printf("Pokemonds in %s:\n", locationArea)
	for _, v := range data.PokemonEncounters {
		fmt.Printf("- %s\n", v.Pokemon.Name)
	}

	return nil
}
