package handlers

import (
	"groupie/api"
	"fmt"
	"net/http"
	"html/template"
)
var tmpl, err = template.ParseGlob("./templates/*.html")
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println("Error Passing the Artist template", err)
	}

	artist, err := api.FetchArtists() 
	if err != nil {
		fmt.Println("Error fetching the Artistis data", err)
	}
	er := tmpl.ExecuteTemplate(w, "index.html", artist)
	if er != nil {
		fmt.Println("Error Rendoring the index Template", err)
	}
}