package handlers

import (
	"fmt"
	"net/http"
)

type IndexHandler struct {

}

func (h *IndexHandler) Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello there big boy")
	}
}