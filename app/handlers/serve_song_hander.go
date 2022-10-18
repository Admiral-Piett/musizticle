package handlers

import (
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// Call to ServeFile swap
var serveFile = http.ServeFile

func (h *Handler) serveSong(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("ServeSongFailure")
		http.Error(w, "Invalid ID Provided", http.StatusBadRequest)
	}
	song, err := h.Dao.FindSongById(id)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("ServeSongFailure")
		http.Error(w, "Song Not Found", http.StatusNotFound)
	}
	h.Logger.WithFields(logrus.Fields{
		LogFields.SongID: id,
	}).Info("ServingSong")
	serveFile(w, r, fmt.Sprintf("%s%s", models.SETTINGS.LibraryPath, song.FilePath))
}
