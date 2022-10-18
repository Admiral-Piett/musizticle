package daos

import (
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetAllArtists_success(t *testing.T) {
	dao := resetDao()
	seedArtistOrAlbum("artists", "Tenacious D", dao)

	artists, err := dao.GetAllArtists()

	assert.Nil(t, err)
	assert.Len(t, artists, 1)
	assert.Equal(t, "Tenacious D", artists[0].Name)
}

func Test_GetAllArtists_success_returns_no_rows(t *testing.T) {
	dao := resetDao()
	albums, err := dao.GetAllArtists()

	assert.Nil(t, err)
	assert.Len(t, albums, 0)
	assert.Equal(t, []models.Artist{}, albums)
}
