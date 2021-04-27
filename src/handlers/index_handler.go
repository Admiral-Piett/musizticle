package handlers

import (
	"fmt"
	"net/http"
)

func IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there big boy")
	}
}