package handlers

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) getAlbums(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("GetAlbumsStart")
	songs, err := h.Dao.GetAllAlbums()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetAlbumsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	err = EncodeResponse(w, songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetAlbumsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	h.Logger.Info("GetAlbumsComplete")
}
