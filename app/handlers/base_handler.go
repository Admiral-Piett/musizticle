package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/Admiral-Piett/musizticle/app/utils"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

var LogFields = models.LogFieldStruct{
	AlbumId:      "album_id",
	ArtistId:     "artist_id",
	ErrorMessage: "error_message",
	FilePath:     "file_path",
	SongID:       "song_id",
	RequestBody:  "request_body",
	Size:         "size",
	StackContext: "stack_context",
}

type Handler struct {
	Dao    *daos.Dao
	Logger *logrus.Logger
}

func InitializeHandlers(dao *daos.Dao, logger *logrus.Logger) *Handler {
	//FIXME - Should maybe let the app method do this?
	return &Handler{Dao: dao, Logger: logger}
}

func tokenParsingReturn(token *jwt.Token) (interface{}, error) {
	return models.SETTINGS.TokenKey, nil
}

func (h *Handler) validateHeader(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	s := r.Header.Get("Authorization")
	if s == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.UnauthorizedResponse)
		return r.Context(), errors.New("Authorization header not found")
	}
	authArray := strings.Split(s, "Bearer")
	// This should always be 2 exactly. If it isn't the request is malformed and that's just too bad.
	if 2 != len(authArray){
		h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: "Invalid `Authorization` header format"}).Error("ValidateHeaderError")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.UnauthorizedResponse)
		return r.Context(), errors.New("Authorization header invalid format")
	}
	tokenString := strings.TrimSpace(authArray[1])

	decodedToken := &utils.JwtToken{}
	_, err := jwt.ParseWithClaims(tokenString, decodedToken, tokenParsingReturn)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("ValidateHeaderFailure")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.UnauthorizedResponse)
		return r.Context(), err
	}

	// If we can't decrypt the token then we know it's invalid, so stop.
	userId, err := utils.Decrypt(decodedToken.UserId)
	if err != nil {
		return r.Context(), err
	}

	return context.WithValue(r.Context(), "UserId", userId), nil
}

// HTTP Method Routers
func (h *Handler) Albums() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// FIXME: can I make a wrapper out of this?
		ctx, err := h.validateHeader(w, r)
		if err != nil {
			h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("Unauthorized")
			return
		}
		h.getAlbums(w, r.WithContext(ctx))
		switch r.Method {
		case "GET":
			h.getAlbums(w, r.WithContext(ctx))
		case "POST":
			http.Error(w, "NOT_FOUND", http.StatusNotFound)
			return
		}
	}
}

func (h *Handler) Artists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx, err := h.validateHeader(w, r)
		if err != nil {
			h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("Unauthorized")
			return
		}

		switch r.Method {
		case "GET":
			h.getArtists(w, r.WithContext(ctx))
		case "POST":
			http.Error(w, "NOT_FOUND", http.StatusNotFound)
			return
		}
	}
}

func (h *Handler) ServeSong() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := h.validateHeader(w, r)
		if err != nil {
			h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("Unauthorized")
			return
		}
		h.serveSong(w, r.WithContext(ctx))
	}
}

func (h *Handler) Songs() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx, err := h.validateHeader(w, r)
		if err != nil {
			h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("Unauthorized")
			return
		}
		switch r.Method {
		case "GET":
			h.getSongs(w, r.WithContext(ctx))
		case "POST":
			http.Error(w, "NOT_FOUND", http.StatusNotFound)
			return
		}
	}
}

func (h *Handler) SongsByArtist() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx, err := h.validateHeader(w, r)
		if err != nil {
			h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("Unauthorized")
			return
		}
		switch r.Method {
		case "GET":
			h.getSongsByArtistId(w, r.WithContext(ctx))
		case "POST":
			http.Error(w, "NOT_FOUND", http.StatusNotFound)
			return
		}
	}
}

func (h *Handler) SongsByAlbum() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		ctx, err := h.validateHeader(w, r)
		if err != nil {
			h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("Unauthorized")
			return
		}
		h.getSongsByAlbumId(w, r.WithContext(ctx))
		switch r.Method {
		case "GET":
			h.getSongsByAlbumId(w, r.WithContext(ctx))
		case "POST":
			http.Error(w, "NOT_FOUND", http.StatusNotFound)
			return
		}
	}
}

func (h *Handler) Import() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, err := h.validateHeader(w, r)
		if err != nil {
			h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("Unauthorized")
			return
		}

		h.songImport(w, r.WithContext(ctx))
	}
}
