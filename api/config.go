package api

type Config struct {
	PokeClient   Client
	NextLocation *string
	PrevLocation *string
}

func NewConfig() Config {
	return Config{
		PokeClient:   NewClient(),
		NextLocation: nil,
		PrevLocation: nil,
	}
}
