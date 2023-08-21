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
		fields := strings.Fields(readed)
		command := fields[0]
		options := make([]string, 0, 0)
		if len(fields) > 1 {
			for _, option := range fields[1:] {
				options = append(options, option)
			}

		}
		val, ok := commands.GetCliCommands()[command]
		if !ok {
			fmt.Println("Command does not exist ")
			continue
		}
		err := val.Callback(options...)
		if err != nil {
			fmt.Println(err)
		}
	}

}
