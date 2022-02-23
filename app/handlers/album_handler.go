package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AlbumsHandler struct {
	Dao *daos.Dao
}

func (h *Handler) getAlbums(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("GetAlbumsStart")
	songs, err := h.Dao.GetAllAlbums()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetAlbumsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetAlbumsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	h.Logger.Info("GetAlbumsComplete")
}

func (h *Handler) postAlbums(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("PostAlbumsStart")
	h.Logger.Info("PostAlbumsComplete")
}
