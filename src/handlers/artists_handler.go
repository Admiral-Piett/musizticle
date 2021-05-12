package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Artist struct {
	Name string `json:"name"`
}

func (h *Handler) postArtists(w http.ResponseWriter, r *http.Request) {
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
	//TODO - add artist to postgres via h.ArtistsDao.artists
	return
}

func (h *Handler) getArtists(w http.ResponseWriter, r *http.Request) {

}
