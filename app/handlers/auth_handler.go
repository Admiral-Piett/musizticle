package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/Admiral-Piett/musizticle/app/utils"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func VerifyTokenWrapper(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("PostAuthStart")
	w.Header().Set("Content-Type", "application/json")
	var creds = models.Credentials{}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || (creds == models.Credentials{}) {
		h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("AuthRequestFailure")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.UnauthorizedResponse)
		return
	}

	user, err := h.Dao.GetUser(creds.Username, creds.Password)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("AuthRequestFailure")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.UnauthorizedResponse)
		return
	}

	response, err := generateAuthToken(user)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{LogFields.ErrorMessage: err}).Error("AuthRequestFailure")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.UnauthorizedResponse)
		return
	}
	err = json.NewEncoder(w).Encode(response)
	h.Logger.Info("PostAuthComplete")
}

func generateAuthToken(user models.User) (models.AuthResponse, error) {
	response := models.AuthResponse{}
	now := time.Now()
	expirationTime := now.Add(time.Duration(models.SETTINGS.TokenExpiration) * time.Minute)

	encryptedId, err := utils.Encrypt(user.Id)
	if err != nil {
		return response, err
	}

	tokenFields := utils.JwtToken{
		UserId: encryptedId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			NotBefore: now.Unix(),
			IssuedAt: now.Unix(),
		},
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenFields)
	tokenString, err := token.SignedString(models.SETTINGS.TokenKey)
	if err != nil {
		return response, err
	}

	response.AuthToken = tokenString
	expirationString, _ := expirationTime.UTC().MarshalText()
	response.ExpirationTime = string(expirationString)

	return response, err
}
