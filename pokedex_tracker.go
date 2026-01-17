package main

import (
	"github.com/timbdn01/pokedexcli/internal/pokeapi"
	"fmt"
)

var pokedex = make(map[string]pokeapi.Pokemon)

// use a map[string]Pokemon to track caught pokemon
func trackPokedex(cfg *config, pokemonName string) error {
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}
	if _, exists := pokedex[pokemon.Name]; exists {
		return nil // already tracked
	}
	pokedex[pokemon.Name] = *pokemon
	return nil

}

// getPokedex prints a list of caught pokemon names
func commandPokedex(cfg *config, args string) error {
	fmt.Println("Your Pokedex:")
	if len(pokedex) == 0 {
		fmt.Println("  (no Pokemon caught yet)")
		return nil
	}
	for name := range pokedex {
		fmt.Println(name)
	}
	return nil
}