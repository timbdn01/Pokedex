package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type RespLocationArea struct {
	Name            string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// List all pokemon within a location area
func (c *Client) ExploreLocationArea(locationAreaName string) (RespLocationArea, error) {
	url := baseURL + "/location-area/" + locationAreaName
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationArea{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationArea{}, err
	}
	locationAreaResp := RespLocationArea{}
	err = json.Unmarshal(dat, &locationAreaResp)
	if err != nil {
		return RespLocationArea{}, err
	}
	return locationAreaResp, nil
}