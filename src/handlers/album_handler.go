package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/sound_control/src/daos"
	"github.com/Admiral-Piett/sound_control/src/utils"
	"github.com/sirupsen/logrus"
	"net/http"
)

type AlbumsHandler struct {
	Dao *daos.Dao
}


func (h *Handler) getAlbums(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	songs, err := h.Dao.FetchAllAlbums()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetAlbumsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetAlbumsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
}

func (h *Handler) postAlbums(w http.ResponseWriter, r *http.Request) {

}

