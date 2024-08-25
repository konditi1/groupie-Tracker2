package handlers

import (
	"groupie/models"
	"strconv"
	"fmt"
	"groupie/api"
	"net/http"
)

func DatesHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println("Error Parsing dates html", err)
		return
	}

	artistName := r.URL.Query().Get("name")
	artistIdStr := r.URL.Query().Get("artistId")
	artistId, err:= strconv.Atoi(artistIdStr)
	if err != nil {
		fmt.Println("Error incorrect Artist Id", err)
		return
	}

	dates, err := api.FetchDates()
	if err != nil {
		fmt.Println("Error Fetching the Dates data for the Dates template", err)
		return
	}

	date := dates.Index[artistId]
	DatesData := struct {
		Dates models.Dates
		Name string
	}{
		Dates : date,
		Name : artistName,
	}
	errExe := tmpl.ExecuteTemplate(w, "dates.html", DatesData)
	if errExe != nil {
		fmt.Println("Error Executing the date.html", errExe)
	}
}
