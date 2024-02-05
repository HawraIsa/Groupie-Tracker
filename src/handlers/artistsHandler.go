package groupietracker

// import (
// 	//"groupietracker"
// 	"encoding/json"
// 	"fmt"
// 	//"log"
// 	"net/http"
// 	//"os"
// )

// type Artist struct {
//     Name   string `json:"name"`
//     Image  string `json:"image"`
//     Year   int    `json:"year"`
//     Album  string `json:"firstAlbum"`
//     Members []struct {
//         Name string `json:"name"`
//         Role string `json:"role"`
//     } `json:"members"`
// }

// func fetchArtists() ([]Artist, error) {
// 	url := "https://groupietrackers.herokuapp.com/api/artists"

// 	var artist Artist
	
// 	err := getJson(url, &artist)
// 	if err != nil {
// 		fmt.Printf("error getting artist data %s\n", err.Error())
// 	} else {
// 		fmt.Printf("artist %s\n", artist.Name)
// 	}
// 	// defer response.Body.Close()

// 	// var artists []Artist
// 	// err = json.NewDecoder(response.Body).Decode(&artists)
// 	// if err != nil {
// 	// 	return nil, err
// 	// }

// 	// return artists, nil
// }

// func getJson(url string, target interface {}) error {
// 	response, err := client.get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer response.Body.Close()

// 	return json.NewDecoder(response.Body).Decode(target)
// }

// func artistsHandler(w http.ResponseWriter, r *http.Request) {
// 	// Fetch data from the artists API endpoint
	

//     // response, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
//     // if err != nil {
//     //     fmt.Print(err.Error())
//     //     os.Exit(1)
//     // }


// 	// Serialize the data as JSON and send the response
// 	//json.NewEncoder(w).Encode(artists)
// }