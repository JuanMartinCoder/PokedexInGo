package help

import (
	"fmt"
)

func HelpCommand(list map[string]interface{}) error {
	fmt.Println("Welcome to the Pokedex!\n\n")
	for _, v := range list {
		fmt.Printf("%s: %s", v.name, v.description)
	}
	return nil
}
