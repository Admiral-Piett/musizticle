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

func New() *App {
	logger := logrus.New()
	logger.WithFields(logrus.Fields{"my": "fart"}).Info("Starting Sound Control App...")

	appDaos := daos.InitializeDao()
	defer appDaos.CloseDao()

	appHandler := handlers.InitializeHandlers(appDaos, logger)

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

	a.Router.HandleFunc("/artists", a.Handler.Artists()).Methods("GET", "POST")
	a.Router.HandleFunc("/songs", a.Handler.Songs()).Methods("GET", "POST")

	a.Router.HandleFunc("/import", a.Handler.Import()).Methods("POST")
}
