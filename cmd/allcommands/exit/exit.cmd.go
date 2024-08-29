package exit

import (
	"os"

	"github.com/JuanMartinCoder/PokedexInGo/api"
)

func CommandExit(cfg *api.Config, args ...string) error {
	os.Exit(0)
	return nil
}
