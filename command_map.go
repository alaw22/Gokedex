package main

import (
	"errors"
	"fmt"
)

func commandMap(config *Config, args ...string) error {

	locAreaResp, err := config.gokeapiClient.GetLocationAreas(config.nextLocationareasURL)
	if err != nil {
		return err
	}

	config.nextLocationareasURL = locAreaResp.Next
	config.previousLocationareasURL = locAreaResp.Previous

	for _, result := range locAreaResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}


func commandMapb(config *Config, args ...string) error {
	if config.previousLocationareasURL == nil {
		return errors.New("You are on the first page, no turning back")
	}

	locAreaResp, err := config.gokeapiClient.GetLocationAreas(config.previousLocationareasURL)
	if err != nil {
		return err
	}

	config.nextLocationareasURL = locAreaResp.Next
	config.previousLocationareasURL = locAreaResp.Previous

	for _, result := range locAreaResp.Results {
		fmt.Println(result.Name)
	}

	return nil
}
