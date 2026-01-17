package main

import (
	"errors"
	"fmt"
)

// Inspect caught pokemon details
func commandInspect(cfg *config, pokemonName string) error {
	pokemon, exists := pokedex[pokemonName]
	if !exists {
		return errors.New("You haven't caught that Pokemon yet!")
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, s := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Printf("Types:\n")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}