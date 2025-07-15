package main

import 
(
    "time"
    "github.com/alaw22/Gokedex/internal/api"
)


func main() {

    gokeClient := api.NewClient(5 * time.Second, 5 * time.Second)

    config := &Config{
        gokedex: make(map[string]api.Pokemon),
        gokeapiClient: gokeClient,
    }

    StartREPL(config)
}
