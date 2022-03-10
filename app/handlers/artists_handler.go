package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Artist struct {
	Name string `json:"name"`
}

func (h *Handler) getArtists(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("GetArtistsStart")
	songs, err := h.Dao.GetAllArtists()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetArtistsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetArtistsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	h.Logger.Info("GetArtistsComplete")
}
