package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/Admiral-Piett/musizticle/app/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
)

//TODO - Use me
var InvalidFileTypes = []string{
	".DS_Store",
	".7z",
}

type ImportRequest struct {
	ImportDir string `json:"importDirectory"`
}

func (h *Handler) songImport(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("SongImportStart")
	request := ImportRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.ImportDir == "" {
		h.Logger.WithFields(logrus.Fields{
			LogFields.RequestBody:  r,
			LogFields.StackContext: "songImport",
			LogFields.ErrorMessage: err,
		}).Error("InvalidSongImportRequest")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request Format"))
		return
	}

	go func() {
		h.Logger.WithField(LogFields.RequestBody, request).Info("SongImportStart")
		err = filepath.Walk(request.ImportDir, h.importSong)
		if err != nil {
			h.Logger.Error(err)
		}
		h.Logger.WithField(LogFields.RequestBody, request).Info("SongImportComplete")
	}()
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) importSong(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fi, _ := file.Stat()
	if err != nil {
		return err
	}
	//Return here if this is a directory
	if fi.IsDir() {
		return nil
	}
	if !(fi.Mode().IsRegular()) {
		return nil
	}
	track, err := utils.GetSongMetadata(file)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			LogFields.ErrorMessage: err,
			LogFields.FilePath:     path,
		}).Error("FailureToOpenFile - Skipping")
		return nil
	}
	artistId, err := h.Dao.FindOrCreateByName(track.Artist, daos.QueryArtistIdByName, daos.InsertArtist)
	if err != nil || artistId == -1 {
		checkError("FailureToFindOrAddArtist", path, err, h.Logger)
		return nil
	}
	albumId, err := h.Dao.FindOrCreateByName(track.Album, daos.QueryAlbumIdByName, daos.InsertAlbum)
	if err != nil || albumId == -1 {
		checkError("FailureToFindOrAddAlbum", path, err, h.Logger)
		return nil
	}
	songId, err := h.Dao.FindOrCreateSong(track, artistId, albumId, path, track.Duration, daos.QuerySongIdByName, daos.InsertSongs)
	if err != nil || songId == -1 {
		checkError("FailureToFindOrAddSong", path, err, h.Logger)
		return nil
	}
	h.Logger.WithFields(logrus.Fields{
		LogFields.FilePath: path,
		LogFields.SongID:   songId,
		LogFields.AlbumId:  albumId,
		LogFields.ArtistId: artistId,
		LogFields.Size:     info.Size(),
	}).Debug("SongAdded")
	return nil
}

func checkError(message string, file string, err error, logger *logrus.Logger) {
	logger.WithFields(logrus.Fields{
		LogFields.ErrorMessage: err,
		LogFields.FilePath:     file,
	}).Error(fmt.Sprintf("%s - Skipping", message))
}
