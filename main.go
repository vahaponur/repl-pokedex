package main

import (
	"bufio"
	"fmt"
	commands "internal/commands"
	"strings"

	"os"
)

func main() {

	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		ok := reader.Scan()
		if ok != true {
			fmt.Println("Error occured")
			break
		}
		readed := reader.Text()
		if strings.Contains(readed, "explore") {
			index := strings.Index(readed, " ")
			if index == -1 {
				fmt.Println("Invalid input")
				return
			}
			areaName := readed[index+1:]
			val, ok := commands.GetCliCommands()["explore"]
			if !ok {
				fmt.Println("Command does not exist ")
				continue
			}
			err := val.Callback(areaName)
			if err != nil {
				fmt.Println(err)
			}
			// Extract the substring after "explore"
			continue

		}
		val, ok := commands.GetCliCommands()[readed]
		if !ok {
			fmt.Println("Command does not exist ")
			continue
		}
		err := val.Callback()
		if err != nil {
			fmt.Println(err)
		}
	}

}
