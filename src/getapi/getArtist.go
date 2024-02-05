package getapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func GetArtist() {
	// API url, make HTTP 'GET' request
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return
	} // response body is closed at the end of the request, regardless of any error that may occur
	// to avoid the need to explicitly close the response body and reduce the risk of forgetting to close it.
	defer response.Body.Close()
	// process response body
	body, readErr := io.ReadAll(response.Body) // data in json **response body is in bytes
	if readErr != nil {
		fmt.Print(err.Error())
	}
	// create struct ,  decode JSON data into the struct
	var artists []Artist
	json.Unmarshal(body, &artists)

	if len(artists) > 0 { // testing if it will print the first artist **works
		firstArtist := artists[0]
		fmt.Println(firstArtist.Name)
	}

	/* TO DO : search for a particular user based on name or id ---- Acheived down ---- */
}

func GetArtistByName(name string) (*Artist, error) {

	// API url, make HTTP 'GET' request
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	// Read the response body
	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}

	// Unmarshal the JSON data into the artists slice
	var artists []Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, err
	}

	// Search for the artist by name
	for _, artist := range artists {
		if artist.Name == name {
			// Create a copy of the artist before returning its address
			copyArtist := artist
			return &copyArtist, nil
		}
	}

	// If the artist is not found, return an error
	return nil, fmt.Errorf("artist not found: %s", name)
	// // API url, make HTTP 'GET' request
	// url := "https://groupietrackers.herokuapp.com/api/artists"
	// response, err := http.Get(url)
	// if err != nil {
	// 	return nil, err
	// }
	// defer response.Body.Close()

	// // Read the response body
	// body, readErr := io.ReadAll(response.Body)
	// if readErr != nil {
	// 	return nil, readErr
	// }

	// // Unmarshal the JSON data into the artists slice
	// var artists []Artist
	// if err := json.Unmarshal(body, &artists); err != nil {
	// 	return nil, err
	// }

	// // Search for the artist by name
	// for _, artist := range artists {
	// 	if artist.Name == name {
	// 		return &artist, nil
	// 	}
	// }

	// // If the artist is not found, return an error
	// return nil, fmt.Errorf("artist not found: %s", name)
}
