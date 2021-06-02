package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there big boy")
	}
}