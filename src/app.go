package src

import (
	"bytes"
	"fmt"
	"github.com/Admiral-Piett/sound_control/src/daos"
	"github.com/Admiral-Piett/sound_control/src/handlers"
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/gorilla/mux"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

type App struct {
	Router *mux.Router
	Handlers *handlers.Handlers
}

func New() *App {
	initializePostgres()

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

func initializePostgres() {
	_, scriptPath, _, _ := runtime.Caller(0)
	directory := filepath.Join(filepath.Dir(scriptPath), "..")
	runtimePath := directory + "/postgres/runtime"
	dataPath := directory + "/postgres/data"
	os.MkdirAll(runtimePath, 0755)
	os.MkdirAll(dataPath, 0755)

	fmt.Println(directory)
	// Start up the embedded postgres instance
	logger := &bytes.Buffer{}
	port, _ := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	postgres := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		Username(os.Getenv("POSTGRES_USER")).
		Password(os.Getenv("POSTGRES_PASSWORD")).
		Database(os.Getenv("POSTGRES_DB")).
		Version(embeddedpostgres.V12).
		RuntimePath(fmt.Sprintf("")).
		DataPath(dataPath).
		Port(uint32(port)).
		StartTimeout(10 * time.Second).
		Logger(logger))
	go func() {
		err := postgres.Start()
		defer postgres.Stop()
		if err != nil {
			panic(err)
		}
	}()
}
