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

func Test_serveSong_success(t *testing.T) {
	serveFile = func(w http.ResponseWriter, r *http.Request, name string) {
		fmt.Println("test")
	}
	defer func() {
		serveFile = http.ServeFile
	}()

	w, r := generateRequestInfo("", "", nil)
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	h.serveSong(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}

func Test_serveSong_invalid_param_error(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)

	d := &mocks.DaoMock{}
	d.FindSongByIdMock = func(id int) (models.ListSong, error) {
		return models.ListSong{}, nil
	}
	h := &Handler{
		Dao:    &mocks.DaoMock{},
		Logger: logrus.New(),
	}

	h.serveSong(w, r)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func Test_serveSong_find_song_by_id_error(t *testing.T) {
	w, r := generateRequestInfo("", "", nil)

	d := &mocks.DaoMock{}
	d.FindSongByIdMock = func(id int) (models.ListSong, error) {
		return models.ListSong{}, fmt.Errorf("boom")
	}
	h := &Handler{
		Dao:    d,
		Logger: logrus.New(),
	}
	r = mux.SetURLVars(r, map[string]string{
		"id": "1",
	})

	h.serveSong(w, r)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
