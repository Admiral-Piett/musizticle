package handlers

import (
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/mocks"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_getAlbums_success(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.getAlbums(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_getAlbums_GetAllAlbums_returns_error(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	d := &mocks.DaoMock{}
	d.GetAllAlbumsMock = func() ([]models.Album, error) {
		return []models.Album{}, fmt.Errorf("boom")
	}
	h := &Handler{
		Dao:    d,
		Logger: logrus.New(),
	}

	h.getAlbums(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_getAlbums_EncodeResponse_error(t *testing.T) {
	EncodeResponse = func(w http.ResponseWriter, value interface{}) error {
		return fmt.Errorf("boom")
	}
	defer func() {
		EncodeResponse = encodeResponse
	}()

	w, r := generateRequestInfo("", "", nil)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.getAlbums(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
