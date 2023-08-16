package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("--Pokedex Usage--")
	fmt.Println("")
	coms := getCliCommands()
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

func getCliCommands() map[string]cliCommand {
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
	}
}
