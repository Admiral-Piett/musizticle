package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/musizticle/app/daos"
	"github.com/Admiral-Piett/musizticle/app/utils"
	"github.com/sirupsen/logrus"
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
	h.Logger.Info("SongImportStart")
	request := ImportRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil || request.ImportDir == "" {
		h.Logger.WithFields(logrus.Fields{
			utils.LogFields.RequestBody:  r,
			utils.LogFields.StackContext: "songImport",
			utils.LogFields.ErrorMessage: err,
		}).Error("InvalidSongImportRequest")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request Format"))
		return
	}
	h.Logger.WithField(utils.LogFields.RequestBody, request).Info("SongImportStart")
	//	TODO - to filepath.WalkPath here and scope out directory
	err = filepath.Walk(request.ImportDir, h.importSong)
	if err != nil {
		h.Logger.Error(err)
	}
	h.Logger.WithField(utils.LogFields.RequestBody, request).Info("SongImportComplete")
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
			utils.LogFields.ErrorMessage: err,
			utils.LogFields.FilePath:     path,
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
	songId, err := h.Dao.FindOrCreateSong(track, artistId, albumId, path, daos.QuerySongIdByName, daos.InsertSongs)
	if err != nil || songId == -1 {
		checkError("FailureToFindOrAddSong", path, err, h.Logger)
		return nil
	}
	h.Logger.WithFields(logrus.Fields{
		utils.LogFields.FilePath: path,
		utils.LogFields.SongID:   songId,
		utils.LogFields.AlbumId:  albumId,
		utils.LogFields.ArtistId: artistId,
		utils.LogFields.Size:     info.Size(),
	}).Debug("SongAdded")
	return nil
}

func checkError(message string, file string, err error, logger *logrus.Logger) {
	logger.WithFields(logrus.Fields{
		utils.LogFields.ErrorMessage: err,
		utils.LogFields.FilePath:     file,
	}).Error("%s - Skipping", message)
}
