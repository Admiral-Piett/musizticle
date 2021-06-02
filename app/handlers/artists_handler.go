package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/sound_control/app/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Artist struct {
	Name string `json:"name"`
}

func (h *Handler) postArtists(w http.ResponseWriter, r *http.Request) {
	req := Artist{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("PostArtistFailure")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	if req.Name == "" {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.RequestBody: req,
		}).Error("PostArtistFailure - Invalid Name Field")
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	//TODO - add artist to postgres via h.ArtistsDao.artists
	w.WriteHeader(http.StatusCreated)
	return
}

func (h *Handler) getArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	songs, err := h.Dao.FetchAllArtists()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetArtistsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetArtistsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
}
