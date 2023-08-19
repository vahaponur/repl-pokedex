package main

import (
	"bufio"
	"fmt"
	commands "internal/commands"

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
