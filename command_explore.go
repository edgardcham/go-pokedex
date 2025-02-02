package main

import (
	"fmt"
	"github.com/edgardcham/go-pokedex/internal/pokeapi"
)

func commandExplore(config *Config, parameters []string) error {
	areaName := parameters[0]
	fmt.Printf("Exploring area %s...\n", areaName)
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s/", areaName)
	var encountersResp pokeapi.Encounters
	if err := pokeapi.FetchURL(url, &encountersResp, config.Cache); err != nil {
		return fmt.Errorf("Error fetching encounters: %v", err)
	}

	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range encountersResp.PokemonEncounters {
		fmt.Printf("- %s\n", encounter.Pokemon.Name)
	}

	return nil
}
