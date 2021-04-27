package main

import (
	"github.com/Admiral-Piett/sound_control/src"
	"log"
	"net/http"
)

//TODO - environmentalize
var PORT string = ":9000"

func main() {
	app := src.New()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Printf("App up and running on localhost%s", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
