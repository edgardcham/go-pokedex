package main

import (
	"fmt"
	"github.com/edgardcham/go-pokedex/internal/pokeapi"
)

func commandMapb(config *Config, parameters []string) error {
	prevURL := config.PrevURL
	if prevURL == "" {
		fmt.Println("You're on the first page")
		return nil
	}
	var mapResp pokeapi.MapResponse
	if err := pokeapi.FetchURL(prevURL, &mapResp, config.Cache); err != nil {
		return fmt.Errorf("Error fetching map: %v", err)
	}

	for _, area := range mapResp.Results {
		fmt.Println(area.Name)
	}

	config.NextURL = mapResp.Next
	if mapResp.Previous != nil {
		config.PrevURL = *mapResp.Previous
	} else {
		config.PrevURL = ""
	}

	return nil
}
