package main

import (
	"fmt"
)

func commandPokedex(c *Config, args ...string) error {
	if len(c.gokedex) == 0 {
		fmt.Println("no pokemon caught, go catch 'em all!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemonName, _ := range c.gokedex {
		fmt.Printf(" - %s\n",pokemonName)
	}
	
	return nil
}
