package main

import (
	"fmt"
	"github.com/edgardcham/go-pokedex/internal/pokeapi"
	"math/rand"
)

func commandCatch(config *Config, parameters []string) error {
	pokemonName := parameters[0]

	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s/", pokemonName)
	var pokemonResp pokeapi.Pokemon
	if err := pokeapi.FetchURL(url, &pokemonResp, config.Cache); err != nil {
		return fmt.Errorf("Error fetching Pokemon data: %v", err)
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	baseExp := pokemonResp.BaseExperience
	rawChance := 150 - baseExp
	if rawChance < 5 {
		rawChance = 5
	} else if rawChance > 95 {
		rawChance = 95
	}

	numGenerated := rand.Intn(100)

	if numGenerated < rawChance {
		fmt.Printf("%s was caught!\n", pokemonName)
		config.Pokedex[pokemonResp.Name] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}
