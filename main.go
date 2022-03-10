package main

import (
	"fmt"
	"github.com/Admiral-Piett/musizticle/app"
	"github.com/Admiral-Piett/musizticle/app/models"
	"log"
	"net/http"
)


func main() {
	musizticle := app.New()
	defer musizticle.Handler.Dao.CloseDao()

	http.HandleFunc("/", musizticle.ProxyHandler)

	musizticle.Logger.Info("App up and running on http://localhost:", models.SETTINGS.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", models.SETTINGS.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
