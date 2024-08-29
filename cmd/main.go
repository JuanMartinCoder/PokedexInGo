package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		cleanInput := cleanInput(userInput)
		if len(cleanInput) == 0 {
			continue
		}
		command := cleanInput[0]
		args := []string{}

		if len(cleanInput) > 1 {
			args = cleanInput[1:]
		}
		if cmds.IsACommand(command) {
			if err := cmds.ExcecuteCmd(command)(&cfg, args...); err != nil {
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

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
