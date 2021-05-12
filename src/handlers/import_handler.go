package handlers

import (
	"encoding/json"
	"github.com/Admiral-Piett/sound_control/src/utils"
	"github.com/sirupsen/logrus"
	"net/http"
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
}
