package handlers

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func (h *Handler) getSongs(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("GetSongsStart")
	songs, err := h.Dao.GetAllSongs()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetSongsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	err = EncodeResponse(w, songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetSongsFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	h.Logger.Info("GetSongsComplete")
}

func (h *Handler) getSongsByArtistId(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("GetSongsByArtistsIdStart")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetSongsByArtistIdFailure")
		http.Error(w, "Invalid ID Provided", http.StatusBadRequest)
	}
	songs, err := h.Dao.FindSongsByArtistId(id)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetSongsByArtistIdFailure")
		http.Error(w, "Song Not Found", http.StatusNotFound)
	}
	err = EncodeResponse(w, songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetSongsByArtistIdFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	h.Logger.Info("GetSongsByArtistsIdComplete")
}

func (h *Handler) getSongsByAlbumId(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("GetSongsByAlbumIdStart")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetSongsByAlbumIdFailure")
		http.Error(w, "Invalid ID Provided", http.StatusBadRequest)
	}
	songs, err := h.Dao.FindSongsByAlbumId(id)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetSongsByAlbumIdFailure")
		http.Error(w, "Song Not Found", http.StatusNotFound)
	}
	err = EncodeResponse(w, songs)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
		}).Error("GetSongsByAlbumIdFailure")
		http.Error(w, "General Error", http.StatusInternalServerError)
	}
	h.Logger.Info("GetSongsByAlbumIdComplete")
}
