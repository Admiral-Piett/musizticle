package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/mocks"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

var validRequest = models.Credentials{
	Username: "HanSolo",
	Password: "millenium4eva",
}

func Test_Auth_success(t *testing.T) {
	w, r := generateRequestInfo("", "", validRequest)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.Auth(w, r)

	assert.Equal(t, http.StatusOK, w.Code)

	response := &models.AuthResponse{}
	b, _ := io.ReadAll(r.Body)
	json.Unmarshal(b, response)
	assert.IsType(t, "", response.AuthToken)
	assert.IsType(t, "", response.ExpirationTime)
	assert.Equal(t, "", response.ReauthToken)
}

func Test_Auth_empty_request_body_returns_unauthorized(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.Auth(w, r)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "UNAUTHORIZED")
}

func Test_Auth_invalid_request_returns_unauthorized(t *testing.T) {
	w, r := generateRequestInfo("", "", struct {
		Field string
	}{Field: "Iamgarbage"})

	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.Auth(w, r)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "UNAUTHORIZED")
}

func Test_Auth_GetUser_error_returns_unauthorized(t *testing.T) {
	w, r := generateRequestInfo("", "", validRequest)

	d := &mocks.DaoMock{}
	d.GetUserMock = func(username, password string) (models.User, error) {
		return models.User{}, fmt.Errorf("boom")
	}

	h := &Handler{
		Dao:    d,
		Logger: logrus.New(),
	}

	h.Auth(w, r)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "UNAUTHORIZED")
}

func Test_Auth_generateAuthToken_error_returns_unauthorized(t *testing.T) {
	tokenFunction = func(user models.User) (models.AuthResponse, error) {
		return models.AuthResponse{}, fmt.Errorf("boom")
	}
	defer func() {
		tokenFunction = generateAuthToken
	}()
	w, r := generateRequestInfo("", "", validRequest)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.Auth(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "UNAUTHORIZED")
}

func Test_Auth_EncodeResponse_error(t *testing.T) {
	EncodeResponse = func(w http.ResponseWriter, value interface{}) error {
		return fmt.Errorf("boom")
	}
	defer func() {
		EncodeResponse = encodeResponse
	}()

	w, r := generateRequestInfo("", "", validRequest)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.Auth(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

}
