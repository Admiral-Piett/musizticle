package handlers

import (
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/mocks"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_getSongs_success(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.getSongs(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_getSongs_GetAllSongs_error(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	d := &mocks.DaoMock{}
	d.GetAllSongsMock = func() ([]models.ListSong, error) {
		return []models.ListSong{}, fmt.Errorf("boom")
	}
	h := &Handler{
		Dao:    d,
		Logger: logrus.New(),
	}

	h.getSongs(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_getSongs_EncodeResponse_error(t *testing.T) {
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

	h.getSongs(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_getSongsByArtistId_success(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	h.getSongsByArtistId(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_getSongsByArtistId_invalid_param(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.getSongsByArtistId(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_getSongsByArtistId_FindSongsByArtistId_error(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	d := &mocks.DaoMock{}
	d.FindSongsByArtistIdMock = func(id int) ([]models.ListSong, error) {
		return []models.ListSong{}, fmt.Errorf("boom")
	}
	h := &Handler{
		Dao:    d,
		Logger: logrus.New(),
	}

	h.getSongsByArtistId(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func Test_getSongsByArtistId_EncodeResponse_error(t *testing.T) {
	EncodeResponse = func(w http.ResponseWriter, value interface{}) error {
		return fmt.Errorf("boom")
	}
	defer func() {
		EncodeResponse = encodeResponse
	}()

	w, r := generateRequestInfo("", "", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.getSongsByArtistId(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func Test_getSongsByAlbumId_success(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.getSongsByAlbumId(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_getSongsByAlbumId_invalid_param(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)

	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.getSongsByAlbumId(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_getSongsByAlbumId_FindSongsByAlbumId_error(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	d := &mocks.DaoMock{}
	d.FindSongsByAlbumIdMock = func(id int) ([]models.ListSong, error) {
		return []models.ListSong{}, fmt.Errorf("boom")
	}

	h := &Handler{
		Dao:    d,
		Logger: logrus.New(),
	}

	h.getSongsByAlbumId(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func Test_getSongsByAlbumId_EncodeResponse_error(t *testing.T) {
	EncodeResponse = func(w http.ResponseWriter, value interface{}) error {
		return fmt.Errorf("boom")
	}
	defer func() {
		EncodeResponse = encodeResponse
	}()

	w, r := generateRequestInfo("", "", nil)
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.getSongsByAlbumId(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
