package main

import (
	"errors"
	"fmt"
)

//take in location name, fetch list of pokemon encounters for that location
func commandExplore(cfg *config, locationName string) error {
	location, err := cfg.pokeapiClient.ExploreLocationArea(locationName)
	if err != nil {
		return err
	}
	if len(location.PokemonEncounters) == 0 {
		return errors.New("no pokemon found in this location")
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}
	return nil
}
	

	