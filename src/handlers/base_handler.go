package handlers

import (
	"github.com/Admiral-Piett/sound_control/src/daos"
)

type Handlers struct {
	IndexHandler   *IndexHandler
	ArtistsHandler *ArtistHandler
}

func InitializeHandlers(dao *daos.Dao) *Handlers {
	//FIXME - Should maybe let the app method do this?
	handlers := &Handlers{

		IndexHandler: &IndexHandler{},
		ArtistsHandler: &ArtistHandler{
			Dao: dao,
		},
	}
	return handlers
}


