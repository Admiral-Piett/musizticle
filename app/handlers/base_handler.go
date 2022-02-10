package handlers

import (
	app_cache "github.com/Admiral-Piett/musizticle/app/cache"
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler struct {
	Dao    *daos.Dao
	AppCache *app_cache.AppCache
	Logger *logrus.Logger
}

func InitializeHandlers(dao *daos.Dao, appCache *app_cache.AppCache, logger *logrus.Logger) *Handler {
	//FIXME - Should maybe let the app method do this?
	return &Handler{Dao: dao, AppCache: appCache, Logger: logger}
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

		err := h.AppCache.ValidateSession(w, r)
		if err != nil {
			// If we can't validate the user's session there's no point in going on here.
			return
		}
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
