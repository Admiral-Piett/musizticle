package app

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/Admiral-Piett/musizticle/app/handlers"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/Admiral-Piett/musizticle/app/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gitlab.com/avarf/getenvs"
	"net/http"
	"os"
)

type App struct {
	Router   *mux.Router
	Handler  *handlers.Handler
	Logger   *logrus.Logger
}

func New() *App {
	logger := logrus.New()
	if os.Getenv("LOG_LEVEL") == "DEBUG" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.WithFields(logrus.Fields{"it's": "de bug log!"}).Debug("We're de-buggin' now")
	logger.WithFields(logrus.Fields{"it's a": "log!"}).Info("Starting Sound Control App...")

	InitializeSettings(logger)

	appDaos := daos.InitializeDao()
	appHandler := handlers.InitializeHandlers(appDaos, logger)

	a := &App{
		Logger:   logger,
		Router:   mux.NewRouter(),
		Handler:  appHandler,
	}
	a.initRoutes()
	return a
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/api", a.Handler.Index()).Methods("GET")

	a.Router.HandleFunc("/api/auth", a.Handler.Auth).Methods("POST")
	a.Router.HandleFunc("/api/reauth", a.Handler.ReAuth).Methods("POST")

	a.Router.HandleFunc("/api/albums", a.Handler.Albums()).Methods("GET", "POST")
	a.Router.HandleFunc("/api/artists", a.Handler.Artists()).Methods("GET", "POST")
	a.Router.HandleFunc("/api/songs/{id:[0-9]+}", a.Handler.ServeSong()).Methods("GET")
	a.Router.HandleFunc("/api/songs", a.Handler.Songs()).Methods("GET", "POST")
	a.Router.HandleFunc("/api/songs/artists/{id:[0-9]+}", a.Handler.SongsByArtist()).Methods("GET")
	a.Router.HandleFunc("/api/songs/albums/{id:[0-9]+}", a.Handler.SongsByAlbum()).Methods("GET")

	//a.Router.HandleFunc("/api/search/songs/{name:[0-9]+}", a.Handler.Songs()).Methods("GET", "POST")

	a.Router.HandleFunc("/api/import", a.Handler.Import()).Methods("POST")
}

// --- CORS Proxy ---
// FIXME - environmentalize (or part of the build process?)
var CORS_ALLOW_HEADERS = "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"
var CORS_ALLOW_METHODS = "POST, GET, OPTIONS, PUT, DELETE"
var CORS_ALLOW_ORIGINS = "*"

func setupCORS(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", CORS_ALLOW_ORIGINS)
	(*w).Header().Set("Access-Control-Allow-Methods", CORS_ALLOW_METHODS)
	(*w).Header().Set("Access-Control-Allow-Headers", CORS_ALLOW_HEADERS)
}

// Proxy Handler to deal with all incoming requests in main.go.  If the Method is OPTIONS, assume this is a pre-flight
//  CORS check and return CORS headers here.
func (a *App) ProxyHandler(w http.ResponseWriter, req *http.Request) {
	setupCORS(&w, req)
	if req.Method == "OPTIONS" {
		return
	}
	a.Router.ServeHTTP(w, req)
}

func InitializeSettings(logger *logrus.Logger) {
	models.SETTINGS.Port = getenvs.GetEnvString("MUSIZTICLE_PORT", "9000")
	models.SETTINGS.SqliteDB = getenvs.GetEnvString("MUSIZTICLE_SQLITE_DB", "musizticle.db")
	models.SETTINGS.TokenExpiration, _ = getenvs.GetEnvInt("MUSIZTICLE_TOKEN_EXPIRATION", 1)

	tokenKeyLength, _ := getenvs.GetEnvInt("MUSIZTICLE_TOKEN_KEY_LENGTH", 100)
	ts := utils.GenerateRandomString(tokenKeyLength)
	logger.WithFields(logrus.Fields{"token_string": ts}).Debug("TokenStringCreated")
	models.SETTINGS.TokenKey = []byte(ts)

	// TODO: set our own key - so that we can encrypt the username/password with the gui before sending
	// So every time we start up generate a new encryption key for the id.  But the token key is something I set.
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	models.SETTINGS.PrivateKey = privateKey
	models.SETTINGS.PublicKey = &privateKey.PublicKey
}
