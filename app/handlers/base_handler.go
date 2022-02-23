package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/Admiral-Piett/musizticle/app/utils"
	"github.com/dgrijalva/jwt-go"
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

func (h *Handler) validateHeader(w http.ResponseWriter, r *http.Request) error {
	s := r.Header.Get("Authorization")
	if s == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.UnauthorizedResponse)
		return errors.New("Authorization header not found")
	}
	authArray := strings.Split(s, "Bearer")
	// This should always be 2 exactly. If it isn't the request is malformed and that's just too bad.
	if 2 != len(authArray){
		h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: "Invalid `Authorization` header format"}).Error("ValidateHeaderError")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.UnauthorizedResponse)
		return errors.New("Authorization header invalid format")
	}
	tokenString := strings.TrimSpace(authArray[1])

	decodedToken := &utils.JwtToken{}
	token, err := jwt.ParseWithClaims(tokenString, decodedToken, tokenParsingReturn)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("ValidateHeaderFailure")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.UnauthorizedResponse)
		return err
	}

	//TODO - Send decrypted user id into handlers.
	// Set a user context for the scope of the request if possible?  Storing the user id.


	fmt.Println(s)
	fmt.Println(token)

	return nil
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

		// TODO - put in between almost all requests
		err := h.validateHeader(w, r)
		if err != nil {
			h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("Unauthorized")
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
