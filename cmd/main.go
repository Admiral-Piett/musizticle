package main

import (
	"fmt"
	"github.com/Admiral-Piett/sound_control/src"
	"github.com/Admiral-Piett/sound_control/src/daos"
	"log"
	"net/http"
	"os"
)

//TODO - environmentalize
var PORT string = fmt.Sprintf(":%s", os.Getenv("PORT"))

func main() {
	appDaos := daos.InitializeDao()
	defer appDaos.CloseDao()
	app := src.New(appDaos)

	http.HandleFunc("/", app.Router.ServeHTTP)

	app.Logger.WithField("port", os.Getenv("PORT")).Info("App up and running on localhost", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
