package main

import (
	"fmt"
	"errors"
	// "github.com/alaw22/Gokedex/internal/api"
)


func commandInspect(c *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("inspect requires a pokemon name")
	}

	pokemonName := args[0]

	if pokemon, exists := c.gokedex[pokemonName]; exists {
		fmt.Println("Name:",pokemonName)
		fmt.Println("Height:",pokemon.Height)
		fmt.Println("Weight:",pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%s: %v\n",stat.Stat.Name,stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, _type := range pokemon.Types {
			fmt.Printf("  - %s\n",_type.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}

	return nil
}