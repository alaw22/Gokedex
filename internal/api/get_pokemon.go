package api

import (
	"io"
	"net/http"
	"encoding/json"
)


func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	url := GokedexURL + "/pokemon/" + pokemonName

	if data, exists := c.gokeCache.Get(url); exists {
		pokemon := Pokemon{}
		err := json.Unmarshal(data,&pokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET",url,nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(data,&pokemon)
	if err != nil {
		return Pokemon{}, nil
	}
	
	// Add to cache
	c.gokeCache.Add(url,data)

	return pokemon, nil 
}