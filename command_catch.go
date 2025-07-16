package main

import (
	"fmt"
	"errors"
	"math/rand"
)

const (
	experienceScaler = 80.0
)

func commandCatch(c *Config, args ...string) error {
	if len(args) == 0 {
		return errors.New("catch requires a pokemon name")
	}

	pokemonName := args[0]

	pokemon, err := c.gokeapiClient.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n",pokemonName)
	probability := experienceScaler/float64(pokemon.BaseExperience)
	// fmt.Printf("probability: %v\n",probability)
	if chance := rand.Float64(); chance <= probability {
		// fmt.Printf("chance = %v\n",chance)
		// You caught it!
		fmt.Printf("%s was caught!\n",pokemonName)
		// Add to gokedex map
		c.gokedex[pokemonName] = pokemon
	} else {
		// fmt.Printf("chance = %v\n",chance)
		fmt.Printf("%s escaped!\n",pokemonName)
	}

	return nil

}