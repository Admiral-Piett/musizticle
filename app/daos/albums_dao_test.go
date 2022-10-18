package daos

import (
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetAllAlbums_success(t *testing.T) {
	dao := resetDao()
	seedArtistOrAlbum("albums", "The greatest and best album in the world - Tribute.", dao)

	albums, err := dao.GetAllAlbums()

	assert.Nil(t, err)
	assert.Len(t, albums, 1)
	assert.Equal(t, "The greatest and best album in the world - Tribute.", albums[0].Name)
}

func Test_GetAllAlbums_success_returns_no_rows(t *testing.T) {
	dao := resetDao()
	albums, err := dao.GetAllAlbums()

	assert.Nil(t, err)
	assert.Len(t, albums, 0)
	assert.Equal(t, []models.Album{}, albums)
}
