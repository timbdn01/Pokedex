package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"math/rand"
	"errors"
)

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`

}

func (c *Client) GetPokemon(pokemonName string) (*Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch pokemon data")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, err
	}
	return &pokemon, nil
}

// take in pokemon name, attempt to catch it using a random chance compared to its base experience
func (c *Client) CatchPokemon(pokemonName string) (*Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch pokemon data")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		return nil, err
	}
	catchChance := rand.Intn(256)
	if catchChance > pokemon.BaseExperience {
		return &pokemon, nil
	} else {
		return nil, errors.New("pokemon escaped")
	}
}