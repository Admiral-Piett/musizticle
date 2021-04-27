package src

import (
	"github.com/Admiral-Piett/sound_control/src/handlers"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}
	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", handlers.IndexHandler()).Methods("GET")
	a.Router.HandleFunc("/artists", handlers.Artists()).Methods("GET", "POST")
}
