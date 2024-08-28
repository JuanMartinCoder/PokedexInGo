package mapcmd

import (
	"fmt"

	"github.com/JuanMartinCoder/PokedexInGo/api"
)

func MapCommand(cfg *api.Config) error {
	data, err := cfg.PokeClient.ListLocationArea(cfg.NextLocation)
	if err != nil {
		return err
	}

	for _, v := range data.Results {
		fmt.Printf("- %s\n", v.Name)
	}

	cfg.NextLocation = data.Next
	cfg.PrevLocation = data.Previous

	return nil
}
