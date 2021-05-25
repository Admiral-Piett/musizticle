package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Admiral-Piett/sound_control/src/daos"
	"github.com/Admiral-Piett/sound_control/src/utils"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type ImportRequest struct {
	ImportDir string `json:"importDirectory"`
}

func (h *Handler) Import() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.songImport(w, r)
	}
}

func (h *Handler) songImport(w http.ResponseWriter, r *http.Request) {
	request := ImportRequest{}
	err := json.NewDecoder(r.Body).Decode(&request);
	if err != nil || request.ImportDir == "" {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.RequestBody: r,
			utils.LogFields.StackContext: "songImport",
			utils.LogFields.ErrorMessage: err,
		}).Error("InvalidSongImportRequest")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request Format"))
		return
	}
	h.Logger.WithField(utils.LogFields.RequestBody, request).Info("ImportRequest")
//	TODO - to filepath.WalkPath here and scope out directory
	err = filepath.Walk(request.ImportDir, h.importSong)
	if err != nil {
		log.Println(err)
	}
}

func (h *Handler) importSong(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	track, err := utils.GetSongMetadata(file)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.ErrorMessage: err,
			utils.LogFields.FilePath: path,
		}).Error("FailureToOpenFile - Skipping")
		return nil
	}
	artistId, err := h.Dao.FindOrCreateByName(track.Artist, daos.QueryArtistIdByName, daos.InsertArtist, true)
	if err != nil || artistId == -1 {
		checkError("FailureToFindOrAddArtist", path, err, h.Logger)
		return nil
	}
	albumId, err := h.Dao.FindOrCreateByName(track.Album, daos.QueryAlbumIdByName, daos.InsertAlbum, true)
	if err != nil || albumId == -1 {
		checkError("FailureToFindOrAddAlbum", path, err, h.Logger)
		return nil
	}
	songId, err := h.Dao.FindOrCreateSong(track.Title, artistId, albumId, path, daos.QuerySongIdByName, daos.InsertSongs, true)
	fmt.Println(artistId, albumId, songId, path, info.Size())
	return nil
}

func checkError(message string, file string, err error, logger *logrus.Logger) {
	logger.WithFields(logrus.Fields{
		utils.LogFields.ErrorMessage: err,
		utils.LogFields.FilePath: file,
	}).Error("%s - Skipping", message)
}

