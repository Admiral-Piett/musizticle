package main

import (
	"fmt"
	"github.com/Admiral-Piett/musizticle/app"
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/Admiral-Piett/musizticle/app/models"
	"log"
	"net/http"
)


func main() {
	app.InitializeSettings()

	appDaos := daos.InitializeDao()
	defer appDaos.CloseDao()
	app := app.New(appDaos)

	http.HandleFunc("/", app.ProxyHandler)

	app.Logger.Info("App up and running on http://localhost:", models.SETTINGS.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", models.SETTINGS.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
