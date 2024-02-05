package getapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Date struct {
	ConcertDates []string `json:"dates"`
}

func GetDates() {
	url := "https://groupietrackers.herokuapp.com/api/dates"
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

	// var dates []Date
	// json.Unmarshal(body, &dates)

	// for i, date := range dates {
	// 		fmt.Printf("Date %d: %s\n", i+1, date.ConcertDates[0])
	// }

	var dates struct {
		Index []struct {
			ID    int      `json:"id"`
			Dates []string `json:"dates"`
		} `json:"index"`
	}
	json.Unmarshal(body, &dates)

	for i, date := range dates.Index {
		fmt.Printf("Date %d: %s\n", i+1, date.Dates[0])
	}
}

/* --------------------------------------------- Get dates by Artist ----------------------------------------- */
func GetDatesByArtist(artist string) ([]string, error) {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}

	var dates struct {
		Index []struct {
			ID    int      `json:"id"`
			Dates []string `json:"dates"`
		} `json:"index"`
	}
	if err := json.Unmarshal(body, &dates); err != nil {
		return nil, err
	}

	var foundDates []string

	for _, date := range dates.Index {
		for _, d := range date.Dates {
			if d == artist {
				foundDates = append(foundDates, d)
			}
		}
	}

	if len(foundDates) > 0 {
		return foundDates, nil
	}

	return nil, fmt.Errorf("no dates found for artist: %s", artist)
}
