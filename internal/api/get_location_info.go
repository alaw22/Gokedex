package api

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func (c *Client) GetLocationPokemon(locationName string) (LocationAreaRespInfo, error) {

	url := GokedexURL + "/location-area/" + locationName

	if data, exists := c.gokeCache.Get(url); exists {
		fmt.Println("USING CACHE BABY")
		location_info := LocationAreaRespInfo{}
		err := json.Unmarshal(data, &location_info)
		if err != nil {
			return LocationAreaRespInfo{}, err
		}
		return location_info, nil
	}
	

	// Make Request
	req, err := http.NewRequest("GET",url,nil)

	if err != nil {
        return LocationAreaRespInfo{}, err
	}

	// Ask http.Client to do our request
	res, err := c.httpClient.Do(req)

	if err != nil {
        return LocationAreaRespInfo{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
        return LocationAreaRespInfo{}, err
	}

	// Unmarshal Data into LocationAreaRespInfo struct
	location_info := LocationAreaRespInfo{}

	err = json.Unmarshal(data, &location_info)
	
	if err != nil {
		return LocationAreaRespInfo{}, err
	}

	// Add to cache
	c.gokeCache.Add(url,data)

	return location_info, nil

}