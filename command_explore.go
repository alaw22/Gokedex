package main

import (
    "fmt"
    "errors"
)

func commandExplore(config *Config, args ...string) error {
    if len(args) == 0 {
        return errors.New("Explore requires a location")
    }

    location := args[0]

    location_info, err := config.gokeapiClient.GetLocationPokemon(location)

    if err != nil {
        return err
    }

    fmt.Printf("Exploring %s...\n",location)
    fmt.Println("Found Pokemon:")
    for _, encounter := range location_info.PokemonEncounters {
        fmt.Printf(" - %s\n",encounter.Pokemon.Name)
    }

    return nil    
}
