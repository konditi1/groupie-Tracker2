package handlers

import (
	"fmt"
	"groupie/api"
	"groupie/models"
	"net/http"
	"strconv"
)

func RelationHandler(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		fmt.Println("Error Parsing Relations Template", err)
		return
	}

	artistName := r.URL.Query().Get("name")
	artistIdStr := r.URL.Query().Get("artistId")
	artistId, err := strconv.Atoi(artistIdStr)
	if err != nil {
		fmt.Println("Error incorrect Artist Id", err)
		return
	}
	relations, err := api.FetchRelation()
	if err != nil {
		fmt.Println("Error Fetching Relations Data", err)
		return
	}

	relation := relations.Index[artistId]
	relationsData := struct {
		Name     string
		Relation models.Relation
	}{
		Name:     artistName,
		Relation: relation,
	}
	fmt.Println(relation)
	execErr := tmpl.ExecuteTemplate(w, "relation.html", relationsData)
	if execErr != nil {
		fmt.Println("Error Executing relation html", execErr)
	}
}
