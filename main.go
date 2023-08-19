package main

import (
	"bufio"
	"fmt"
	commands "internal/commands"
	"internal/pokeapi"
	"net/url"
	"os"
)

func main() {
	locArea := pokeapi.GetNextLocationArea(url.URL{RawPath: "https://pokeapi.co/api/v2/location-area/"})
	fmt.Println(locArea.Next)
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		ok := reader.Scan()
		if ok != true {
			fmt.Println("Error occured")
			break
		}
		readed := reader.Text()
		val, ok := commands.GetCliCommands()[readed]
		if !ok {
			fmt.Println("Command does not exist ")
			continue
		}
		val.Callback()
	}

}
