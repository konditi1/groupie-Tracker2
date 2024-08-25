package route

import (
	"groupie/handlers"
	"net/http"
)

func Routes() {
	http.HandleFunc("/", handlers.ArtistHandler)
	http.HandleFunc("/locations", handlers.LocationHandler)
	http.HandleFunc("/dates", handlers.DatesHandler)
	http.HandleFunc("/relations", handlers.RelationHandler)
}
