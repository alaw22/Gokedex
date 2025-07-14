package api

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreaRespBase, error){
	url := GokedexURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if data, exists := c.gokeCache.Get(url); exists {
		fmt.Println("USING CACHE BABY")
		location_area_base := LocationAreaRespBase{}
		err := json.Unmarshal(data, &location_area_base)
		if err != nil {
			return LocationAreaRespBase{}, err
		}
		return location_area_base, nil
	}

	// Make Request
	req, err := http.NewRequest("GET",url,nil)

	if err != nil {
        return LocationAreaRespBase{}, err
	}

	// Ask http.Client to do our request
	res, err := c.httpClient.Do(req)

	if err != nil {
        return LocationAreaRespBase{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)

	if err != nil {
        return LocationAreaRespBase{}, err
	}

	// Unmarshal Data into LocationAreaRespBase struct
	location_area_base := LocationAreaRespBase{}

	err = json.Unmarshal(data, &location_area_base)
	
	if err != nil {
		return LocationAreaRespBase{}, err
	}

	// Add to cache
	c.gokeCache.Add(url,data)

	return location_area_base, nil

}