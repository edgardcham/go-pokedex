package main

import (
	"bufio"
	"fmt"
	"github.com/edgardcham/go-pokedex/internal/pokeapi"
	"github.com/edgardcham/go-pokedex/internal/pokecache"
	"os"
	"strings"
	"time"
)

func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	config := NewConfig()
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(config, words[1:])
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display the next 20 map locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 map locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the area for pokemon",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt capturing a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon in the Pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display Pokemon in the Pokedex",
			callback:    commandPokedex,
		},
	}
}

// Config to handle pagination
type Config struct {
	NextURL string
	PrevURL string
	Cache   *pokecache.Cache
	Pokedex map[string]pokeapi.Pokemon
}

func NewConfig() *Config {
	return &Config{
		NextURL: "https://pokeapi.co/api/v2/location-area?limit=20&offset=0",
		PrevURL: "",
		Cache:   pokecache.NewCache(time.Second * 5),
		Pokedex: make(map[string]pokeapi.Pokemon),
	}
}
