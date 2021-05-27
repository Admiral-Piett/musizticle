package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/sound_control/src/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) getSongs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	songs, err := h.Dao.FetchAllSongs()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetSongsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetSongsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
}

func (h *Handler) postSongs(w http.ResponseWriter, r *http.Request) {

}
