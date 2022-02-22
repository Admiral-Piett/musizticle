package handlers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

var errorResponse = models.ErrorResponse{Code: "UNAUTHORIZED", Message: "Unauthorized"}

func VerifyTokenWrapper(w http.ResponseWriter, r *http.Request) {

}


func encrypt(value int) ([]byte, error){
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		models.SETTINGS.PublicKey,
		[]byte(strconv.Itoa(value)),
		nil)
	if err != nil {
		return []byte{}, err
	}
	return encryptedBytes, nil
}

func generateAuthToken(user models.User) ([]byte, error) {
	now := time.Now()
	expirationTime := now.Add(time.Duration(models.SETTINGS.TokenExpiration) * time.Minute).Unix()

	encryptedId, err := encrypt(user.Id)
	if err != nil {
		return []byte{}, err
	}
	encryptedCreatedTime, err := encrypt(int(now.Unix()))
	if err != nil {
		return []byte{}, err
	}
	encryptedExpirationTime, err := encrypt(int(expirationTime))
	if err != nil {
		return []byte{}, err
	}

	tokenFields := models.JwtToken{
		UserId: encryptedId,
		CreatedAt: encryptedCreatedTime,
		ExpiresAt: encryptedExpirationTime,
	}
	token, err := json.Marshal(tokenFields)
	if err != nil {
		return []byte{}, err
	}

	jwt.NewWithClaims(jwt.SigningMethodHS256, token)

	return token, nil
}

func (h *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("PostAuthStart")
	w.Header().Set("Content-Type", "application/json")
	var creds = models.Credentials{}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{models.LogFields.ErrorMessage: err}).Error("AuthRequestFailure")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	user, err := h.Dao.GetUser(creds.Username, creds.Password)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{models.LogFields.ErrorMessage: err}).Error("AuthRequestFailure")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	_, err = generateAuthToken(user)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{models.LogFields.ErrorMessage: err}).Error("AuthRequestFailure")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorResponse)
		return
	}

	h.Logger.Info("PostAuthComplete")
}

func (h *Handler) ReAuth(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("PostReAuthStart")
	w.Header().Set("Content-Type", "application/json")
	h.Logger.Info("PostReAuthComplete")
}
