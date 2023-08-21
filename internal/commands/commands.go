package commands

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	pokeapi "github.com/vahaponur/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	Callback    func(dynamic_optional ...string) error
}
type MapConfig struct {
	Next     *string
	Previous *string
}
type LocationArea = pokeapi.LocationArea
type AreaDetails = pokeapi.AreaDetails
type Pokemon = pokeapi.Pokemon

var pokedex map[string]Pokemon = make(map[string]pokeapi.Pokemon)
var firstLocationAreaURL = "https://pokeapi.co/api/v2/location-area/"
var currentLocationAreaURL = "https://pokeapi.co/api/v2/location-area/"
var config MapConfig = MapConfig{Next: &currentLocationAreaURL, Previous: nil}

func commandMap(dynamic_optional ...string) error {
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
func commandCatch(dynamic_optional ...string) error {
	pokeName := dynamic_optional[0]
	if _, ok := pokedex[pokeName]; ok {
		return errors.New(fmt.Sprintf("%v already caughted", pokeName))
	}
	pokemon := pokeapi.GetPokemonFromName(pokeName)
	source := rand.NewSource(time.Now().UnixMilli())
	r1 := rand.New(source)
	possibilityDecider := r1.Intn(pokemon.BaseExperience)
	fmt.Printf("Throwing pokeball to %v\n", pokeName)
	if possibilityDecider < 80 {
		pokedex[pokeName] = pokemon
		fmt.Printf("%v caughted\n", pokeName)
	} else {
		fmt.Printf("%v escaped\n", pokeName)
	}
	return nil
}
func commandInspect(dynamic_optional ...string) error {
	pokeName := dynamic_optional[0]
	pokemon, ok := pokedex[pokeName]
	if !ok {
		return errors.New(fmt.Sprintf("You have not caught that pokemon yet"))
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typea := range pokemon.Types {
		fmt.Printf(" -%v\n", typea.Type.Name)
	}
	return nil
}
func commandMapb(dynamic_optional ...string) error {
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
func commandHelp(dynamic_optional ...string) error {
	fmt.Println("--Pokedex Usage--")
	fmt.Println("")
	coms := GetCliCommands()
	for key, val := range coms {
		fmt.Printf("%v : %v\n", key, val.description)
		fmt.Println("")
	}
	return nil
}

func commandExit(dynamic_optional ...string) error {
	os.Exit(0)
	return nil
}
func commandExplore(dynamic_optional ...string) error {
	pokes := pokeapi.GetPokemonsFromLocArea(dynamic_optional[0])
	fmt.Println("Exploring Pokemons")
	for _, encounter := range pokes.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil

}
func GetCliCommands(dynamic_optional ...string) map[string]cliCommand {
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
		"explore": {
			name:        "explore <location_area_name>",
			description: "Get Pokemons to encounter in that location area",
			Callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Throws a Pokeball to given pokemon, success rate depens on base experience\n (Higher is harder)",
			Callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Inspects a pokemon if it is on the Pokedex",
			Callback:    commandInspect,
		},
	}
}
