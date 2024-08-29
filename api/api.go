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

type LocationArea struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

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

func (c *Client) GetLocationArea(locationArea string) (LocationArea, error) {
	fullUrl := baseUrl + "/location-area/" + locationArea

	data, ok := c.cache.Get(fullUrl)

	if ok {
		fmt.Println("Cache hit")
		pokemonData := LocationArea{}
		err := json.Unmarshal(data, &pokemonData)
		if err != nil {
			return LocationArea{}, err
		}
	}

	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		return LocationArea{}, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return LocationArea{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}

	pokemonData := LocationArea{}
	err = json.Unmarshal(body, &pokemonData)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullUrl, body)

	return pokemonData, nil
}
