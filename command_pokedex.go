package main

import "fmt"

func commandPokedex(config *Config, parameters []string) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range config.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
