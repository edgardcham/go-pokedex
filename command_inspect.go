package main

import (
	"fmt"
)

func commandInspect(config *Config, parameters []string) error {
	// If the pokemon is not in the pokedex, return "You have not caught that pokemon"
	// Otherwise display:
	/*
		Name: pidgey
		Height: 3
		Weight: 18
		Stats:
		  -hp: 40
		  -attack: 45
		  -defense: 40
		  -special-attack: 35
		  -special-defense: 35
		  -speed: 56
		Types:
		  - normal
		  - flying

	*/

	pokemonName := parameters[0]
	pokemon, ok := config.Pokedex[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}
