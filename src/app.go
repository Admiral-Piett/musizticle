package src

import (
	"github.com/Admiral-Piett/sound_control/src/daos"
	"github.com/Admiral-Piett/sound_control/src/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type App struct {
	Router *mux.Router
	Handler *handlers.Handler
	Logger *logrus.Logger
}

func New(dao *daos.Dao) *App {
	logger := logrus.New()
	logger.WithFields(logrus.Fields{"my": "fart"}).Info("Starting Sound Control App...")

	appHandler := handlers.InitializeHandlers(dao, logger)

	a := &App{
		Logger: logger,
		Router:   mux.NewRouter(),
		Handler: appHandler,
	}
	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.Handler.Index()).Methods("GET")

	a.Router.HandleFunc("/api/albums", a.Handler.Albums()).Methods("GET", "POST")
	a.Router.HandleFunc("/api/artists", a.Handler.Artists()).Methods("GET", "POST")
	a.Router.HandleFunc("/api/songs/{id:[0-9]+}", a.Handler.ServeSong()).Methods("GET")
	a.Router.HandleFunc("/api/songs", a.Handler.Songs()).Methods("GET", "POST")
	a.Router.HandleFunc("/api/songs/artists/{id:[0-9]+}", a.Handler.GetSongsByArtist()).Methods("GET")
	a.Router.HandleFunc("/api/songs/albums/{id:[0-9]+}", a.Handler.GetSongsByArtist()).Methods("GET")

	//a.Router.HandleFunc("/api/search/songs/{name:[0-9]+}", a.Handler.Songs()).Methods("GET", "POST")

	a.Router.HandleFunc("/api/import", a.Handler.Import()).Methods("POST")
}
