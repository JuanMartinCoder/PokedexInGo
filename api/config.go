package api

import "time"

type Config struct {
	PokeClient   Client
	NextLocation *string
	PrevLocation *string
}

func NewConfig(interval time.Duration) Config {
	return Config{
		PokeClient:   NewClient(interval),
		NextLocation: nil,
		PrevLocation: nil,
	}
}
