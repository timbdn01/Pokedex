package main

import (
	"fmt"
)

//catch pokemon at a location
func commandCatch(cfg *config, pokemonName string) error {
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	pokemon, err := cfg.pokeapiClient.CatchPokemon(pokemonName)
	if err != nil {
		return err
	}
	if pokemon == nil {
		fmt.Printf("%s escaped!\n", pokemonName)
	} else {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		fmt.Printf("You may now inspect %s using the 'inspect %s' command.\n", pokemon.Name, pokemon.Name)
	}
	trackPokedex(cfg,pokemon.Name)
	return nil
}