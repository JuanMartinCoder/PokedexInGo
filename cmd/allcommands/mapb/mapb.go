package mapb

import (
	"fmt"

	"github.com/JuanMartinCoder/PokedexInGo/api"
)

func MapbCommand(cfg *api.Config) error {
	if cfg.PrevLocation == nil {
		return fmt.Errorf("You are in the first page. You can't go back")
	}

	data, err := cfg.PokeClient.ListLocationArea(cfg.PrevLocation)
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
