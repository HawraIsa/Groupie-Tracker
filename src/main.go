package main

import (
	//"groupietracker/src/getapi"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type ConcertDate struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Year         int      `json:"creationDate"`
	Album        string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertdates"`
	Relations    string   `json:"relations"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Details struct {
	Artist       Artist
	Locations    Location
	ConcertDates ConcertDate
	Relation     Relations
}

type ErrorPageData struct {
	StatusCode int
	Message    string
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/details", detailsHandler)

	link := "http://localhost:8080/"

	fmt.Println("\033[36mServer Connected...\033[0m")
	fmt.Printf("\033[36mlink on: %s\033[0m\n", link)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting the server :", err)
	}
}

/* --------------------------------------------------------------- Main Page Handler ----------------------------------------------------- */
func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		// Check if the requested URL path is one of the allowed keywords
		path := r.URL.Path
		if path != "/" && path != "/home" && path != "" && path != "/bands" {
			renderErrorPage(w, http.StatusNotFound, "Page Not Found")
			return
		}

		tmp1, _ := template.ParseFiles("html/index.html")

		w.WriteHeader(200)

		url := "https://groupietrackers.herokuapp.com/api/artists"
		response, err := http.Get(url)
		if err != nil {
			fmt.Print(err.Error())
			renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		defer response.Body.Close()

		body, readErr := io.ReadAll(response.Body)
		if readErr != nil {
			fmt.Print(err.Error())
			renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}

		var artists []Artist
		json.Unmarshal(body, &artists)

		tmp1.Execute(w, artists)
	}
}

/* --------------------------------------------------------------- Artist Details Handler ----------------------------------------------------- */
func detailsHandler(w http.ResponseWriter, r *http.Request) {
	artistName := r.URL.Query().Get("name")
	tmp1, _ := template.ParseFiles("html/details.html")

	url := "https://groupietrackers.herokuapp.com/api/artists"
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	defer response.Body.Close()

	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		fmt.Print(err.Error())
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	var artists []Artist
	json.Unmarshal(body, &artists)

	var foundArtist *Artist
	for _, artist := range artists {
		if artist.Name == artistName {
			foundArtist = &artist
			break
		}
	}

	if foundArtist == nil {
		renderErrorPage(w, http.StatusNotFound, "Artist not found")
		return
	}

	location, err := getLocation(foundArtist.Locations)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	Dates, err := getdates(foundArtist.ConcertDates)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	Relation, err := getRelations(foundArtist.Relations)
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	err = tmp1.Execute(w, Details{Artist: *foundArtist, Locations: location, ConcertDates: Dates, Relation: Relation})
	if err != nil {
		renderErrorPage(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
}

/* --------------------------------------------------------------- Render Error Page Handler ----------------------------------------------------- */

func renderErrorPage(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)

	errorPageData := ErrorPageData{
		StatusCode: statusCode,
		Message:    message,
	}

	errorTemplate, err := template.ParseFiles("html/error.html")
	if err != nil {
		fmt.Println("Error parsing error.html template:", err)
		return
	}

	err = errorTemplate.Execute(w, errorPageData)
	if err != nil {
		fmt.Println("Error executing error.html template:", err)
		return
	}
}

func getLocation(url string) (Location, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return Location{}, err
	}

	defer response.Body.Close()
	locationsData, _ := io.ReadAll(response.Body)
	var artists Location
	json.Unmarshal(locationsData, &artists)
	return artists, nil
}

func getdates(url string) (ConcertDate, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return ConcertDate{}, err
	}

	defer response.Body.Close()
	Dates, _ := io.ReadAll(response.Body)
	var artists ConcertDate
	json.Unmarshal(Dates, &artists)
	return artists, nil
}

func getRelations(url string) (Relations, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return Relations{}, err
	}

	defer response.Body.Close()
	Relation, _ := io.ReadAll(response.Body)
	var artists Relations
	json.Unmarshal(Relation, &artists)
	return artists, nil
}
