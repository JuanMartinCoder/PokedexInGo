package catch

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/JuanMartinCoder/PokedexInGo/api"
)

func CatchCommand(cfg *api.Config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("Explore command takes 1 argument <pockemon-name>")
	}
	pokemon := args[0]

	data, err := cfg.PokeClient.GetPokemonStats(pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s\n", pokemon)
	random := float64(rand.Intn(data.BaseExperience))
	threshold := math.Floor(float64(data.BaseExperience) * 0.25)

	if random > threshold {
		cfg.PokeClient.AddPokemon(data)
		// Catch the pockemon
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s scaped!\n", pokemon)
	}

	return nil
}
