package main

import (
	"bufio"
	"fmt"
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
		val, ok := getCliCommands()[readed]
		if !ok {
			fmt.Println("Unknown error occured exiting...")
		}
		val.callback()
	}

}
