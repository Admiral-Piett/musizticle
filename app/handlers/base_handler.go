package handlers

import (
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	Dao    *daos.Dao
	Logger *logrus.Logger
}

func InitializeHandlers(dao *daos.Dao, logger *logrus.Logger) *Handler {
	//FIXME - Should maybe let the app method do this?
	return &Handler{Dao: dao, Logger: logger}
}

// HTTP Method Routers
func (h *Handler) Albums() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "GET":
			h.getAlbums(w, r)
		case "POST":
			h.postAlbums(w, r)
		}
	}
}

func (h *Handler) Artists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "GET":
			h.getArtists(w, r)
		case "POST":
			h.postArtists(w, r)
		}
	}
}

func (h *Handler) ServeSong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			h.serveSong(w, r)
		default:
			return
		}
	}
}

func (h *Handler) Songs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case "GET":
			h.getSongs(w, r)
		case "POST":
			h.postSongs(w, r)
		}
	}
}
