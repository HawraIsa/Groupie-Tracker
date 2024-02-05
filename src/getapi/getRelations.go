package getapi

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

// type Relation struct {
//     Artists   []Artist   `json:"artists"`
//     Locations []Location `json:"locations"`
//     Dates     []Date     `json:"dates"`
// }

type Response struct {
	Index []struct {
		ID         int                 `json:"id"`
		DatesLocat map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func GetRelations() {
	url := "https://groupietrackers.herokuapp.com/api/relation"
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

    // var relations []Relation
	// json.Unmarshal(body, &relations)

	// // for _, relation := range relations {
	// // 	fmt.Printf("Relation : %s"), relation.
	// // }

	// Unmarshal the JSON data into the Response struct
	
	var data Response
	json.Unmarshal(body, &data)

	for i, item := range data.Index {
		fmt.Printf("Relation %d:\n", i+1)
		for location, dates := range item.DatesLocat {
			fmt.Printf("Location: %s\n", location)
			for _, date := range dates {
				fmt.Printf("  Date: %s\n", date)
			}
		}
		fmt.Println()
	}
}

