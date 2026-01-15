package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	scanner := bufio.NewScanner(os.Stdin)

	

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		input := scanner.Text()
		commands := cleanInput(input)
		if len(commands) == 0 {
			continue
		}
		command := commands[0]
		if cmd, exists := initCommands()[command]; exists {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error executing command %q: %v\n", command, err)
			}
		} else {
			fmt.Printf("Please enter a valid command. You entered: %s\n", command)
		}
	}
}