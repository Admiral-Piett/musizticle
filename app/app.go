package app

import (
	"github.com/Admiral-Piett/sound_control/app/daos"
	"github.com/Admiral-Piett/sound_control/app/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io/fs"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	Handler *handlers.Handler
	Logger *logrus.Logger
	FrontEnd *fs.FS
}

func New(dao *daos.Dao, distFS fs.FS) *App {
	logger := logrus.New()
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.WithFields(logrus.Fields{"it's a": "fart"}).Info("Starting Sound Control App...")

	appHandler := handlers.InitializeHandlers(dao, logger)

	a := &App{
		Logger: logger,
		Router:   mux.NewRouter(),
		Handler: appHandler,
		FrontEnd: &distFS,
	}
	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.Handle("/", http.FileServer(http.FS(*a.FrontEnd))).Methods("GET")

	a.Router.HandleFunc("/api/albums", a.Handler.Albums()).Methods("GET", "POST")
	a.Router.HandleFunc("/api/artists", a.Handler.Artists()).Methods("GET", "POST")
	a.Router.HandleFunc("/api/songs/{id:[0-9]+}", a.Handler.ServeSong()).Methods("GET")
	a.Router.HandleFunc("/api/songs", a.Handler.Songs()).Methods("GET", "POST")
	a.Router.HandleFunc("/api/songs/artists/{id:[0-9]+}", a.Handler.GetSongsByArtist()).Methods("GET")
	a.Router.HandleFunc("/api/songs/albums/{id:[0-9]+}", a.Handler.GetSongsByArtist()).Methods("GET")

	//a.Router.HandleFunc("/api/search/songs/{name:[0-9]+}", a.Handler.Songs()).Methods("GET", "POST")

	a.Router.HandleFunc("/api/import", a.Handler.Import()).Methods("POST")
}
