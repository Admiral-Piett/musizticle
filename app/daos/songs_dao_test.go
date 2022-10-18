package daos

import (
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetAllSongs_success(t *testing.T) {
	dao := resetDao()
	seedSong("Wonderboy", dao, 0, 0, "")

	songs, err := dao.GetAllSongs()

	assert.Nil(t, err)
	assert.Len(t, songs, 1)
	assert.Equal(t, "Wonderboy", songs[0].Title)
	assert.Equal(t, "test-artist", songs[0].ArtistName)
	assert.Equal(t, "test-album", songs[0].AlbumName)
}

func Test_GetAllSongs_success_returns_no_rows(t *testing.T) {
	dao := resetDao()
	songs, err := dao.GetAllSongs()

	assert.Nil(t, err)
	assert.Len(t, songs, 0)
	assert.Equal(t, []models.ListSong{}, songs)
}

func Test_FindSongById_success(t *testing.T) {
	dao := resetDao()
	id := seedSong("Wonderboy", dao, 0, 0, "")

	song, err := dao.FindSongById(int(id))

	assert.Nil(t, err)
	assert.Equal(t, "Wonderboy", song.Title)
}

func Test_FindSongById_success_returns_no_rows(t *testing.T) {
	dao := resetDao()
	songs, err := dao.FindSongById(0)

	assert.Nil(t, err)
	assert.Equal(t, models.ListSong{}, songs)
}

func Test_FindSongsByArtistId_success(t *testing.T) {
	dao := resetDao()
	artistId := seedArtistOrAlbum("artists", "Tenacious D", dao)
	albumId := seedArtistOrAlbum("albums", "Tenacious D (album)", dao)
	id1 := seedSong("Tribute", dao, artistId, albumId, "/path/1")
	id2 := seedSong("Wonderboy", dao, artistId, albumId, "/path/2")

	songs, err := dao.FindSongsByArtistId(int(artistId))

	assert.Nil(t, err)
	assert.Equal(t, int(id1), songs[0].Id)
	assert.Equal(t, "Tribute", songs[0].Title)
	assert.Equal(t, int(id2), songs[1].Id)
	assert.Equal(t, "Wonderboy", songs[1].Title)
}

func Test_FindSongsByArtistId_success_returns_no_rows(t *testing.T) {
	dao := resetDao()

	songs, err := dao.FindSongsByArtistId(0)

	assert.Nil(t, err)
	assert.Equal(t, []models.ListSong{}, songs)
}

func Test_FindSongsByAlbumId_success(t *testing.T) {
	dao := resetDao()
	artistId := seedArtistOrAlbum("artists", "Tenacious D", dao)
	albumId := seedArtistOrAlbum("albums", "Tenacious D (album)", dao)
	id1 := seedSong("Tribute", dao, artistId, albumId, "/path/1")
	id2 := seedSong("Wonderboy", dao, artistId, albumId, "/path/2")

	songs, err := dao.FindSongsByAlbumId(int(albumId))

	assert.Nil(t, err)
	assert.Equal(t, int(id1), songs[0].Id)
	assert.Equal(t, "Tribute", songs[0].Title)
	assert.Equal(t, int(id2), songs[1].Id)
	assert.Equal(t, "Wonderboy", songs[1].Title)
}

func Test_FindSongsByAlbumId_success_returns_no_rows(t *testing.T) {
	dao := resetDao()

	songs, err := dao.FindSongsByAlbumId(0)

	assert.Nil(t, err)
	assert.Equal(t, []models.ListSong{}, songs)
}

func Test_FindOrCreateSong_finds_song_and_returns_data(t *testing.T) {
	dao := resetDao()
	artistId := seedArtistOrAlbum("artists", "Tenacious D", dao)
	albumId := seedArtistOrAlbum("albums", "Tenacious D (album)", dao)
	seededId := seedSong("Tribute", dao, artistId, albumId, "/path/1")

	id, err := dao.FindOrCreateSong(models.SongMeta{Title: "Tribute"}, artistId, albumId, "/path/1", 30, QuerySongIdByName, InsertSongs)

	assert.Nil(t, err)
	assert.Equal(t, seededId, id)
}

func Test_FindOrCreateSong_creates_song_and_returns_data(t *testing.T) {
	dao := resetDao()

	id, err := dao.FindOrCreateSong(models.SongMeta{Title: "Tribute"}, 0, 0, "/path/1", 30, QuerySongIdByName, InsertSongs)

	assert.Nil(t, err)
	assert.Greater(t, id, int64(0))
}

func Test_FindOrCreateSong_query_error_returns_error(t *testing.T) {
	dao := resetDao()

	id, err := dao.FindOrCreateSong(models.SongMeta{Title: "Tribute"}, 0, 0, "/path/1", 30, "SELECT id FROM song;", InsertSongs)

	assert.Error(t, err)
	assert.Equal(t, id, int64(-1))
}

func Test_FindOrCreateSong_scan_error_returns_error(t *testing.T) {
	dao := resetDao()
	artistId := seedArtistOrAlbum("artists", "Tenacious D", dao)
	albumId := seedArtistOrAlbum("albums", "Tenacious D (album)", dao)
	seedSong("Tribute", dao, artistId, albumId, "")
	query := `
		SELECT
			   *
		FROM
			 songs
		WHERE name LIKE "%%%s%%"
		AND artistId = %d
		AND albumId = %d
	`

	id, err := dao.FindOrCreateSong(models.SongMeta{Title: "Tribute"}, artistId, albumId, "", 30, query, InsertSongs)

	assert.Error(t, err)
	assert.Equal(t, id, int64(-1))
}

func Test_FindOrCreateSong_prepare_error_returns_error(t *testing.T) {
	dao := resetDao()

	id, err := dao.FindOrCreateSong(models.SongMeta{Title: "Tribute"}, 0, 0, "", 30, QuerySongIdByName, "garbage")

	assert.Error(t, err)
	assert.Equal(t, id, int64(-1))
}
