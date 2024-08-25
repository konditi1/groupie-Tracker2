package api

import (
	"encoding/json"
	"fmt"
	"groupie/models"
	"net/http"
)

func FetchArtists() ([]models.Artist, error) {
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error Fetching Artist", err)
		return nil, err
	}
	defer res.Body.Close()

	var artists []models.Artist
	err = json.NewDecoder(res.Body).Decode(&artists)
	if err != nil {
		fmt.Println("Error Decording Artist", err)
		return nil, err
	}
	return artists, nil
}
