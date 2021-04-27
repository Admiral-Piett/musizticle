package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/sound_control/src/daos"
	"log"
	"net/http"
)

func Artists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getArtists(w, r)
		case "POST":
			postArtists(w, r)
		}
	}
}

type Artist struct {
	Name string `json:"name"`
}

func postArtists(w http.ResponseWriter, r *http.Request) {
	req := Artist{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("ERROR - %s", err)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		log.Printf("ERROR - Invalid name field: `%s`", req.Name)
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	artistsTable := daos.ArtistDao{}
	//FIXME - this isn't working
	err := artistsTable.Open()
	if err != nil {
		log.Printf("ERROR - Cannot connect to database: `%s`", err)
		http.Error(w, "Internl server error", http.StatusInternalServerError)
		return
	}
	defer artistsTable.Close()
	return
}

func getArtists(w http.ResponseWriter, r *http.Request) {

}
