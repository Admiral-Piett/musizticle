package handlers

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) getArtists(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("GetArtistsStart")
	songs, err := h.Dao.GetAllArtists()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetArtistsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	err = EncodeResponse(w, songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetArtistsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	h.Logger.Info("GetArtistsComplete")
}
