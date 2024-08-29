# Pokedex CLI in Go 


## What is the purpose of this project?

the purpose is to build a Pokedex in a command-line REPL using Golang and the Pokedex API.

### What is REPL?

A REPL (Read-Eval-Print-Loop) is a programming environment that allows users to enter commands and see the results of their commands as they are being executed.

### What is Pokedex API?

Pokedex is a free and open-source API that provides access to Pokémon data. It is a great way to learn about the Pokémon universe and how to use the API.

## Learning Goals

- Practice writing a CLI application in Golang
- Practice making HTTP Requests in Go 
- Learn how to parse JSON in Go  
- Learn about caching and how to use it to improve performance

## Techonologies Used

- [Go](https://golang.org/)
- [Go HTTP Client](https://golang.org/pkg/net/http/)
- [JSON](https://golang.org/pkg/encoding/json/)
- [Pokedex API](https://pokeapi.co/)

## Commands

- help - Shows a list of Commands
- exit - Exits the CLI 
- map  - Displays the names of 20 location areas in the Pokemon world 
- mapb - Similar to the map command, however, instead of displaying the next 20 locations, it displays the previous 20 locations. It's a way to go back.
- explore [area-name] - List all the Pokemon in a given area 
- catch [pokemon-name] - Catch a Pokemon 
- inspect [pokemon-name] - Inspect a Pokemon's stats

## Structure of the Project
The basic idea was to make as separated as possible all the concerns of the application. For example, take the allCommands/ folder, 
inside we have separated in different files each command to do exactly what the need to.

![imagen](https://github.com/user-attachments/assets/12ba43ab-c9d4-4dae-9a64-52a688651a52)

