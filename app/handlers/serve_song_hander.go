package handlers

import (
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) serveSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			models.LogFields.ErrorMessage: err,
		}).Error("ServeSongFailure")
		http.Error(w, "Invalid ID Provided", http.StatusBadRequest)
	}
	song, err := h.Dao.FindSongById(id)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			models.LogFields.ErrorMessage: err,
		}).Error("ServeSongFailure")
		http.Error(w, "Song Not Found", http.StatusNotFound)
	}
	h.Logger.WithFields(logrus.Fields{
		models.LogFields.SongID: id,
	}).Info("ServingSong")
	http.ServeFile(w, r, song.FilePath)
}
