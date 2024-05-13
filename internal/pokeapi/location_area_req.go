package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas (pageURL *string) (LocationAreaResponse,error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil{
		fullURL = *pageURL
	}

	dat, ok := c.cache.Get(fullURL)

	if ok{
		locationsAreasResp := LocationAreaResponse{}
		err := json.Unmarshal(dat, &locationsAreasResp)
		if err != nil {
			return LocationAreaResponse{}, err
		}
		fmt.Println("Cache hit")
		return locationsAreasResp, nil
	}
	fmt.Println("Cache miss")


	req,err := http.NewRequest("GET",fullURL,nil)
	
	if err != nil{
		return LocationAreaResponse{},err
	}

	resp,err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399{
		return LocationAreaResponse{}, fmt.Errorf("bad status code %v", resp.StatusCode)
	}

	data,err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	locationsAreasResp := LocationAreaResponse{}

	err = json.Unmarshal(data, &locationsAreasResp)

	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(fullURL,data)

	return locationsAreasResp, nil

}