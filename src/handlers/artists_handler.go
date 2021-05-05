package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/sound_control/src/daos"
	"log"
	"net/http"
)

type ArtistHandler struct {
	ArtistsDao *daos.ArtistsDao
}

type Artist struct {
	Name string `json:"name"`
}

// HTTP Method Router
func (h *ArtistHandler) Artists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			h.getArtists(w, r)
		case "POST":
			h.postArtists(w, r)
		}
	}
}

func (h *ArtistHandler) postArtists(w http.ResponseWriter, r *http.Request) {
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

func (h *ArtistHandler) getArtists(w http.ResponseWriter, r *http.Request) {

}
