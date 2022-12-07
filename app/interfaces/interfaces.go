package interfaces

import (
	"github.com/Admiral-Piett/musizticle/app/models"
)

type AbstractDao interface {
	CloseDao()
	FindOrCreateByName(name, findQuery, insertQuery string) (int64, error)
	FindOrCreateSong(track models.SongMeta, artistId int64, albumId int64, path string, duration int, findQuery string, insertQuery string) (int64, error)
	FindSongById(id int) (models.ListSong, error)
	FindSongsByArtistId(id int) ([]models.ListSong, error)
	FindSongsByAlbumId(id int) ([]models.ListSong, error)
	GetAllAlbums() ([]models.Album, error)
	GetAllArtists() ([]models.Artist, error)
	GetAllSongs() ([]models.ListSong, error)
	GetUser(username, password string) (models.User, error)
}
