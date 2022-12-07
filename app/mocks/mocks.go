package mocks

import (
	"github.com/Admiral-Piett/musizticle/app/models"
)

type DaoMock struct {
	CloseDaoCalled            bool
	FindOrCreateByNameCalled  bool
	FindOrCreateSongCalled    bool
	FindSongByIdCalled        bool
	FindSongsByArtistIdCalled bool
	FindSongsByAlbumIdCalled  bool
	GetAllAlbumsCalled        bool
	GetAllArtistsCalled       bool
	GetAllSongsCalled         bool
	GetUserCalled             bool
	CloseDaoMock              func()
	FindOrCreateByNameMock    func(name, findQuery, insertQuery string) (int64, error)
	FindOrCreateSongMock      func(track models.SongMeta, artistId int64, albumId int64, path string, duration int, findQuery string, insertQuery string) (int64, error)
	FindSongByIdMock          func(id int) (models.ListSong, error)
	FindSongsByArtistIdMock   func(id int) ([]models.ListSong, error)
	FindSongsByAlbumIdMock    func(id int) ([]models.ListSong, error)
	GetAllAlbumsMock          func() ([]models.Album, error)
	GetAllArtistsMock         func() ([]models.Artist, error)
	GetAllSongsMock           func() ([]models.ListSong, error)
	GetUserMock               func(username, password string) (models.User, error)
}

func (m *DaoMock) CloseDao() {
	m.CloseDaoCalled = true
	if m.CloseDaoMock != nil {
		m.CloseDaoMock()
	}
}

func (m *DaoMock) FindOrCreateByName(name string, findQuery string, insertQuery string) (int64, error) {
	m.FindOrCreateByNameCalled = true
	if m.FindOrCreateByNameMock != nil {
		return m.FindOrCreateByNameMock(name, findQuery, insertQuery)
	}
	return int64(0), nil
}

func (m *DaoMock) FindOrCreateSong(track models.SongMeta, artistId int64, albumId int64, path string, duration int, findQuery string, insertQuery string) (int64, error) {
	m.FindOrCreateSongCalled = true
	if m.FindOrCreateSongMock != nil {
		return m.FindOrCreateSongMock(track, artistId, albumId, path, duration, findQuery, insertQuery)
	}
	return int64(0), nil
}

func (m *DaoMock) FindSongById(id int) (models.ListSong, error) {
	m.FindSongByIdCalled = true
	if m.FindSongByIdMock != nil {
		return m.FindSongByIdMock(id)
	}
	return models.ListSong{}, nil
}

func (m *DaoMock) FindSongsByArtistId(id int) ([]models.ListSong, error) {
	m.FindSongsByArtistIdCalled = true
	if m.FindSongsByArtistIdMock != nil {
		return m.FindSongsByArtistIdMock(id)
	}
	return []models.ListSong{}, nil
}

func (m *DaoMock) FindSongsByAlbumId(id int) ([]models.ListSong, error) {
	m.FindSongsByAlbumIdCalled = true
	if m.FindSongsByAlbumIdMock != nil {
		return m.FindSongsByAlbumIdMock(id)
	}
	return []models.ListSong{}, nil
}

func (m *DaoMock) GetAllAlbums() ([]models.Album, error) {
	m.GetAllAlbumsCalled = true
	if m.GetAllAlbumsMock != nil {
		return m.GetAllAlbumsMock()
	}
	return []models.Album{}, nil
}

func (m *DaoMock) GetAllArtists() ([]models.Artist, error) {
	m.GetAllArtistsCalled = true
	if m.GetAllArtistsMock != nil {
		return m.GetAllArtistsMock()
	}
	return []models.Artist{}, nil
}

func (m *DaoMock) GetAllSongs() ([]models.ListSong, error) {
	m.GetAllSongsCalled = true
	if m.GetAllSongsMock != nil {
		return m.GetAllSongsMock()
	}
	return []models.ListSong{}, nil
}

func (m *DaoMock) GetUser(username, password string) (models.User, error) {
	m.GetUserCalled = true
	if m.GetUserMock != nil {
		return m.GetUserMock(username, password)
	}
	return models.User{}, nil
}
