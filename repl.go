package main

import (
	"fmt"
	"os"
	"bufio"
	"github.com/alaw22/Gokedex/internal/api"
)

type Config struct {
	gokeapiClient api.Client
	gokedex map[string]api.Pokemon
	nextLocationareasURL *string
	previousLocationareasURL *string
}

func StartREPL(config *Config) {
    scanner := bufio.NewScanner(os.Stdin)
    // cache := pokecache.NewCache(5*time.Second)

    /* 
        This is a REPL (Read-Evaluate-Print-Loop) much like
        any shell. It reads the users input evaluates it by 
        doing some functionality and prints the result. 
    */
    for {
        // Prompt
        fmt.Print("Pokedex > ")

        // Get input
        scanner.Scan()

        // Clean input
        clean := cleanInput(scanner.Text())
        if len(clean) == 0{
            continue
        }

        commandName := clean[0]
		arguments := clean[1:]
        command, ok := getCommands()[commandName]
        if !ok {
            fmt.Println("Unknown command")
            continue
        }

        err := command.callback(config, arguments...)
        if err != nil{
            fmt.Println(err)
        }

    }
}

type cliCommand struct {
	name 		string 
	description string
	callback 	func(config *Config, args ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"catch": cliCommand{
			name: "catch",
			description: "Attempt to catch a pokemon",
			callback: commandCatch,
		},
		"explore": cliCommand{
			name: "explore",
			description: "Explore an area",
			callback: commandExplore,
		},
		"map": cliCommand{
			name: "map",
			description: "Display the next 20 location areas",
			callback: commandMap,
		},
		"mapb": cliCommand{
			name: "mapb",
			description: "Display the previous 20 location areas",
			callback: commandMapb,
		},
		"exit": cliCommand{
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"help": cliCommand{
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
}