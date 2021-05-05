package src

import (
	"github.com/Admiral-Piett/sound_control/src/daos"
	"github.com/Admiral-Piett/sound_control/src/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	Handlers *handlers.Handlers
}

func New() *App {
	appDaos := daos.InitializeDaos()
	defer appDaos.CloseAllDaos()
	appHandlers := handlers.InitializeHandlers(appDaos)

	a := &App{
		Router:   mux.NewRouter(),
		Handlers: appHandlers,
	}
	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.Handlers.IndexHandler.Index()).Methods("GET")
	a.Router.HandleFunc("/artists", a.Handlers.ArtistsHandler.Artists()).Methods("GET", "POST")
}
