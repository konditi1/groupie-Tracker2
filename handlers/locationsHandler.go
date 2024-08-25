package handlers

import (
	"fmt"
	"groupie/api"
	"groupie/models"
	"net/http"
	"strconv"
)

func LocationHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println("Error Parsing the location template", err)
		return
	}

	artistIdStr := r.URL.Query().Get("artistId")
	artistName := r.URL.Query().Get("name")
	artistImage := r.URL.Query().Get("image")
	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		fmt.Println("Invalid artist ID", err)
		return
	}

	locations, err := api.FetchLocations()
	location := locations.Index[artistId]
	if err != nil {
		fmt.Println("Error Fetching the locations Id", err)
		return
	}
	LocationsData := struct {
		Name      string
		Image     string
		Id        int
		Locations models.Locations
	}{
		Name:      artistName,
		Image:     artistImage,
		Id:        artistId,
		Locations: location,
	}

	err = tmpl.ExecuteTemplate(w, "locations.html", LocationsData)
	if err != nil {
		fmt.Println("Error Executing the locations.html", err)
	}
}
