package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/sound_control/src/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
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

func (h *Handler) GetSongsByArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.getSongsByArtistId(w, r)
	}
}

func (h *Handler) getSongsByArtistId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetSongsByArtistIdFailure")
		http.Error(w, "Invalid ID Provided", http.StatusBadRequest)
	}
	songs, err := h.Dao.FindSongsByArtistId(id)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetSongsByArtistIdFailure")
		http.Error(w, "Song Not Found", http.StatusNotFound)
	}
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetSongsByArtistIdFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
}

func (h *Handler) GetSongsByAlbum() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.getSongsByArtistId(w, r)
	}
}

func (h *Handler) getSongsByAlbumId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetSongsByAlbumIdFailure")
		http.Error(w, "Invalid ID Provided", http.StatusBadRequest)
	}
	songs, err := h.Dao.FindSongsByArtistId(id)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetSongsByAlbumIdFailure")
		http.Error(w, "Song Not Found", http.StatusNotFound)
	}
	err = json.NewEncoder(w).Encode(songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
		}).Error("GetSongsByAlbumIdFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
}


