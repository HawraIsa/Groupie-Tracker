package getapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetArtistDetails() {
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

	GetArtistName(body)
	GetArtistImage(body)
	GetArtistAlbum(body)
	GetArtistMembers(body)
	GetArtistYear(body)
}

func GetArtistName(b []byte) {
	var artists []Artist
	json.Unmarshal(b, &artists)

	firstArtist := artists[0]
	fmt.Printf("Artist name: %s \n", firstArtist.Name)
}

func GetArtistImage(b []byte) {
	var artists []Artist
	json.Unmarshal(b, &artists)

	firstArtist := artists[0]
	fmt.Printf("Artist image link: %s \n", firstArtist.Image)
}

func GetArtistMembers(b []byte) {
	var artists []Artist
	json.Unmarshal(b, &artists)

	if len(artists) > 0 {
		firstArtist := artists[0]
		for i, member := range firstArtist.Members {
			fmt.Printf("Artist member %d: %s\n", i+1, member)
		}
	} else {
		fmt.Println("No artists found.")
	}
	//firstArtist := artists[0]

	// for _, artist := range artists {
	// 	for i, member := range artist.Members {
	// 		fmt.Printf("Artist member %d: %s\n", i+1, member.Name)
	// 		fmt.Printf("Artist member role : %s\n", member.Role)
	// 	}
	// }
}

func GetArtistYear(b []byte) {
	var artists []Artist
	json.Unmarshal(b, &artists)

	firstArtist := artists[0]
	fmt.Printf("Artist Creation Year: %d \n", firstArtist.Year)
}

func GetArtistAlbum(b []byte) {
	var artists []Artist
	json.Unmarshal(b, &artists)

	firstArtist := artists[0]
	fmt.Printf("Artist album: %s \n", firstArtist.Album)
}

/* ----------------------------- Get artist details by name OR id --------------------------------------- */

func GetArtistDetailsByName(artistName string) ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}

	var artists []Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, err
	}

	var foundArtists []Artist

	for _, artist := range artists {
		if artist.Name == artistName {
			foundArtists = append(foundArtists, artist)
		}
	}

	if len(foundArtists) > 0 {
		return foundArtists, nil
	}

	return nil, fmt.Errorf("no artist details found for name: %s", artistName)
}

func GetArtistDetailsByID(artistID int) (*Artist, error) {
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%d", artistID)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return nil, readErr
	}

	var artist Artist
	if err := json.Unmarshal(body, &artist); err != nil {
		return nil, err
	}

	return &artist, nil
}
