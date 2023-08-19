package commands

import (
	"errors"
	"fmt"
	"os"

	pokeapi "github.com/vahaponur/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	Callback    func() error
}
type MapConfig struct {
	Next     *string
	Previous *string
}
type LocationArea = pokeapi.LocationArea

var firstLocationAreaURL = "https://pokeapi.co/api/v2/location-area/"
var currentLocationAreaURL = "https://pokeapi.co/api/v2/location-area/"
var config MapConfig = MapConfig{Next: &currentLocationAreaURL, Previous: nil}

func commandMap() error {
	if config.Next == nil {
		return errors.New("This is the end of the map")
	}
	locArea := pokeapi.GetLocationArea(*config.Next)
	for _, area := range locArea.Results {
		fmt.Println(area.Name)
	}

	config.Previous = locArea.Previous
	config.Next = locArea.Next

	return nil
}
func commandMapb() error {
	if config.Previous == nil {
		*config.Next = firstLocationAreaURL
		return errors.New("No data to show, this the beginning of the map")
	}
	locArea := pokeapi.GetLocationArea(*config.Previous)
	for _, area := range locArea.Results {
		fmt.Println(area.Name)
	}

	config.Previous = locArea.Previous
	config.Next = locArea.Next

	return nil
}
func commandHelp() error {
	fmt.Println("--Pokedex Usage--")
	fmt.Println("")
	coms := GetCliCommands()
	for key, val := range coms {
		fmt.Printf("%v : %v\n", key, val.description)
		fmt.Println("")
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

func GetCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get Location Areas (Next 20 for each call)",
			Callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get 20 Previous Location Areas (Throws error if map in the first call)",
			Callback:    commandMapb,
		},
	}
}
