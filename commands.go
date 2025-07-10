package main

import (

        "os"
        "fmt"
)

type cliCommand struct {
	name 		string 
	description string
	callback 	func() error
}

type config struct {
    next     string
    previous string
}

var commandRegistry = map[string]cliCommand{}


func commandMap(c config) error {
    return nil
}

func commandExit(c config) error {
    fmt.Println("Closing the Pokedex... Goodbye!")
    os.Exit(0)
    return fmt.Errorf("Unable to exit loop...")
}

func commandHelp(c config) error {
    fmt.Println("Welcome to the Pokedex!")
    fmt.Println("Usage:\n")
    for _, command := range commandRegistry{
        fmt.Printf("%s: %s\n",command.name,command.description)
    }
    return nil
}


func init() {
    commandRegistry["map"] = cliCommand{
        name: "map",
        description: "Display 20 location areas at a time",
        callback: commandMap,
    }
    commandRegistry["exit"] = cliCommand{
        name: "exit",
        description: "Exit the Pokedex",
        callback: commandExit,
    }
    commandRegistry["help"] = cliCommand{
        name: "help",
        description: "Displays a help message",
        callback: commandHelp,
    }
}