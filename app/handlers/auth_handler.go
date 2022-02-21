package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/musizticle/app/models"
	"net/http"
)

func VerifyTokenWrapper(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	var creds = models.Credentials{}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}

func (h *Handler) ReAuth(w http.ResponseWriter, r *http.Request) {

}
