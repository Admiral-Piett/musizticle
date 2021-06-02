package main

import (
	"embed"
	"fmt"
	"github.com/Admiral-Piett/sound_control/app"
	"github.com/Admiral-Piett/sound_control/app/daos"
	"io/fs"
	"log"
	"net/http"
	"os"
)

//go:embed web/dist
var nextFS embed.FS

//TODO - environmentalize
var PORT string = fmt.Sprintf(":%s", os.Getenv("PORT"))

func main() {
	// Root at the `dist` folder in the web app.
	distFS, err := fs.Sub(nextFS, "web/dist")
	if err != nil {
		log.Fatal(err)
	}

	appDaos := daos.InitializeDao()
	defer appDaos.CloseDao()
	app := app.New(appDaos, distFS)

	http.HandleFunc("/", app.Router.ServeHTTP)

	app.Logger.WithField("port", os.Getenv("PORT")).Info("App up and running on localhost", PORT)
	err = http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
