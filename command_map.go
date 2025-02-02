package main

import (
	"fmt"
	"github.com/edgardcham/go-pokedex/internal/pokeapi"
)

func commandMap(config *Config, parameters []string) error {
	// Get NextURL from Config
	nextURL := config.NextURL
	var mapResp pokeapi.MapResponse
	if err := pokeapi.FetchURL(nextURL, &mapResp, config.Cache); err != nil {
		return fmt.Errorf("Error fetching map: %v", err)
	}
	// Print each location name
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
