package commands

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	Callback    func() error
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
	}
}
