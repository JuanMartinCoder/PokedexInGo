package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/JuanMartinCoder/PokedexInGo/internal/cache"
)

type Client struct {
	cache      *cache.Cache
	httpClient http.Client
}

type PokemonData struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationArea struct{}

const baseUrl = "https://pokeapi.co/api/v2"

func NewClient(interval time.Duration) Client {
	return Client{
		cache: cache.NewCache(interval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) ListLocationArea(pageUrl *string) (PokemonData, error) {
	fullUrl := baseUrl + "/location-area"

	if pageUrl != nil {
		fullUrl = *pageUrl
	}

	data, ok := c.cache.Get(fullUrl)

	if ok {
		fmt.Println("Cache hit")
		pokemonData := PokemonData{}
		err := json.Unmarshal(data, &pokemonData)
		if err != nil {
			return PokemonData{}, err
		}
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return PokemonData{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return PokemonData{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonData{}, err
	}

	pokemonData := PokemonData{}
	err = json.Unmarshal(body, &pokemonData)
	if err != nil {
		return PokemonData{}, err
	}

	c.cache.Add(fullUrl, body)

	return pokemonData, nil
}
