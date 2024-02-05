package getapi

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type Location struct {
    LocationName string `json:"name"`
}

func GetLocations() {
	url := "https://groupietrackers.herokuapp.com/api/locations"
    response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
    defer response.Body.Close()

    body, readErr := io.ReadAll(response.Body) // data in json //response body is byte
    if readErr != nil {
        fmt.Print(err.Error())
    }

    // var locations []Location
	// json.Unmarshal(body, &locations)

	// for _, location := range locations {
	// 	fmt.Println(location.Name)
	// }

	var locations struct {
		Index []struct {
			ID    int    `json:"id"`
			Locations []string `json:"locations"`
		} `json:"index"`
	}
	json.Unmarshal(body, &locations)

	for i, location := range locations.Index {
		fmt.Printf(" %d: %s\n", i+1, location.Locations[0])
	}
}
/* ------------------------------ Get location based on a certain date --------------------------------------------- */

func GetLocationsByDate(date string) ([]string, error) {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}

	var locations struct {
		Index []struct {
			ID        int      `json:"id"`
			Locations []string `json:"locations"`
			Dates     []string `json:"dates"`
		} `json:"index"`
	}
	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, err
	}

	var foundLocations []string

	for _, location := range locations.Index {
		for i, d := range location.Dates {
			if d == date {
				foundLocations = append(foundLocations, location.Locations[i])
			}
		}
	}

	if len(foundLocations) > 0 {
		return foundLocations, nil
	}

	return nil, fmt.Errorf("no locations found for date: %s", date)
}
