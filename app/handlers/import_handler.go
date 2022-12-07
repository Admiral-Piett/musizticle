package handlers

import (
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/Admiral-Piett/musizticle/app/utils"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// TODO - Use me
var InvalidFileTypes = []string{
	".DS_Store",
	".7z",
}

func (h *Handler) songImport(w http.ResponseWriter, r *http.Request) {
	h.Logger.Info("SongImportStart")

	go func() {
		err := filepath.Walk(models.SETTINGS.LibraryPath, h.importSong)
		if err != nil {
			h.Logger.Error(err)
		}
		h.Logger.Info("SongImportComplete")
	}()
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) importSong(path string, info os.FileInfo, err error) error {
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

	p := strings.Split(path, models.SETTINGS.LibraryPath)
	localPath := p[len(p)-1]
	songId, err := h.Dao.FindOrCreateSong(track, artistId, albumId, localPath, track.Duration, daos.QuerySongIdByName, daos.InsertSongs)
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
