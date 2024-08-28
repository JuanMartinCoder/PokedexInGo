package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Client struct {
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

const baseUrl = "https://pokeapi.co/api/v2"

func NewClient() Client {
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Client) ListLocationArea(pageUrl *string) (PokemonData, error) {
	fullUrl := baseUrl + "/location-area?"

	if pageUrl != nil {
		fullUrl = *pageUrl
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

	return pokemonData, nil
}
