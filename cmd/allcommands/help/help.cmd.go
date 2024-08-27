package help

import "fmt"

// HelpCommand is a command that prints help for a given Help commmand

func HelpCommand() error {
	fmt.Println("Welcome to the Pokedex!\n\nUsage:\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}
