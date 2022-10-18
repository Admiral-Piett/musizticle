package handlers

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/Admiral-Piett/musizticle/app/mocks"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func generateRequestInfo(method, url string, body interface{}) (*httptest.ResponseRecorder, *http.Request) {
	if url == "" {
		url = "/health-check"
	}
	if method == "" {
		method = "GET"
	}

	var req *http.Request
	var err error
	if body != nil {
		b, _ := json.Marshal(body)
		request_body := bytes.NewBuffer(b)
		req, err = http.NewRequest(method, url, request_body)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		panic(err)
	}
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	return rr, req
}

func TestMain(m *testing.M) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)

	models.SETTINGS.SqliteDB = "dao_tests.db"
	models.SETTINGS.SqliteDriver = "sqlite3"
	models.SETTINGS.PrivateKey = privateKey
	models.SETTINGS.PublicKey = &privateKey.PublicKey
	models.SETTINGS.TokenExpiration = 10

	code := m.Run()
	os.Exit(code)
}

func Test_InitializeHandlers_success(t *testing.T) {
	dao := daos.InitializeDao()

	handler := InitializeHandlers(dao, logrus.New())

	assert.Equal(t, dao, handler.Dao)
	assert.IsType(t, &logrus.Logger{}, handler.Logger)
}

func Test_validateHeader_success(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	user := models.User{Id: 12345}
	response, _ := generateAuthToken(user)
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AuthToken))

	ctx, err := h.validateHeader(w, r)

	assert.Nil(t, err)
	assert.Equal(t, 12345, ctx.Value("UserId"))
}

func Test_validateHeader_no_authorization_header(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	ctx, err := h.validateHeader(w, r)

	assert.Error(t, err)
	assert.IsType(t, context.Background(), ctx)
}

func Test_validateHeader_authorization_header_no_token(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)

	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	r.Header.Add("Authorization", "Bearer")

	ctx, err := h.validateHeader(w, r)

	assert.Error(t, err)
	assert.IsType(t, context.Background(), ctx)
}

func Test_validateHeader_authorization_header_too_many_parts(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)

	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	r.Header.Add("Authorization", "Bearer Bearer")

	ctx, err := h.validateHeader(w, r)

	assert.Error(t, err)
	assert.IsType(t, context.Background(), ctx)
}

func Test_validateHeader_bad_token(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)

	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	r.Header.Add("Authorization", "Bearer bad-token")

	ctx, err := h.validateHeader(w, r)

	assert.Error(t, err)
	assert.IsType(t, context.Background(), ctx)
}

func Test_validateHeader_bad_user_id_encryption(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)

	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	pubkey := models.SETTINGS.PublicKey
	privateKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	models.SETTINGS.PublicKey = &privateKey.PublicKey
	defer func() {
		models.SETTINGS.PublicKey = pubkey
	}()

	user := models.User{Id: 12345}
	response, _ := generateAuthToken(user)
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", response.AuthToken))

	ctx, err := h.validateHeader(w, r)

	assert.Error(t, err)
	assert.IsType(t, context.Background(), ctx)
}
