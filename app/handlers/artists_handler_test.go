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

func Test_getArtists_success(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.getArtists(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_getArtists_GetAllAlbums_returns_error(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	d := &mocks.DaoMock{}
	d.GetAllArtistsMock = func() ([]models.Artist, error) {
		return []models.Artist{}, fmt.Errorf("boom")
	}
	h := &Handler{
		Dao:    d,
		Logger: logrus.New(),
	}

	h.getArtists(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_getArtists_EncodeResponse_error(t *testing.T) {
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

	h.getArtists(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
