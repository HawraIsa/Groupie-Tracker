package getapi

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type Artist struct {
    Name   string `json:"name"`
    Image  string `json:"image"`
    Year   int    `json:"creationDate"`
    Album  string `json:"firstAlbum"`
    Members  []string `json:"members"`
}

func GetArtists() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
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

    var artists []Artist
	json.Unmarshal(body, &artists)

	for _, artist := range artists {
		fmt.Println(artist.Name)
	}
}

