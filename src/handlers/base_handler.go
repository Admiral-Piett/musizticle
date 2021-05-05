package handlers

import (
	"github.com/Admiral-Piett/sound_control/src/daos"
)

type Handlers struct {
	IndexHandler   *IndexHandler
	ArtistsHandler *ArtistHandler
}

func InitializeHandlers(daos *daos.Daos) *Handlers {
	//FIXME - Should maybe let the app method do this?
	handlers := &Handlers{
		IndexHandler: &IndexHandler{},
		ArtistsHandler: &ArtistHandler{
			ArtistsDao: daos.ArtistsDao,
		},
	}
	return handlers
}
