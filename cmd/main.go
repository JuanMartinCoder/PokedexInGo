package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/JuanMartinCoder/PokedexInGo/api"
	allCommands "github.com/JuanMartinCoder/PokedexInGo/cmd/allcommands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cmds := allCommands.CreateCommandsList()

	cfg := api.NewConfig(time.Minute)

	fmt.Printf("pokedex > ")
	for scanner.Scan() {
		// map the text written by the user to a command
		userInput := scanner.Text()
		// fmt.Println(userInput)
		if cmds.IsACommand(userInput) {
			if err := cmds.ExcecuteCmd(userInput)(&cfg); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Command not found")
		}

		fmt.Printf("pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
