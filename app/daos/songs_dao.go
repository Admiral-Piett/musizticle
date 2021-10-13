package daos

import (
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/utils"
	"regexp"
	"strconv"
)

type ListSong struct {
	Id          int
	Name        string
	ArtistId    int
	ArtistName  string
	AlbumId     int
	AlbumName   string
	TrackNumber int
	PlayCount   int
	FilePath    string
	// FIXME - wtf, these are strings??
	CreatedAt      string
	LastModifiedAt string
}

type Song struct {
	Id          int
	Name        string
	ArtistId    int
	AlbumId     int
	TrackNumber int
	PlayCount   int
	FilePath    string
	// FIXME - wtf, these are strings??
	CreatedAt      string
	LastModifiedAt string
}

func (d *Dao) FetchAllSongs() ([]ListSong, error) {
	songs := []ListSong{}
	rows, err := d.DBConn.Query(QueryAllSongs)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	for rows.Next() {
		r := ListSong{}
		err = rows.Scan(&r.Id, &r.Name, &r.ArtistId, &r.AlbumId, &r.TrackNumber, &r.PlayCount, &r.FilePath, &r.CreatedAt, &r.LastModifiedAt, &r.ArtistName, &r.AlbumName)
		if err != nil {
			return songs, err
		}
		songs = append(songs, r)
	}

	return songs, nil
}

func (d *Dao) FindSongById(id int) (Song, error) {
	query := fmt.Sprintf(QuerySongById, id)
	rows, err := d.DBConn.Query(query)
	if err != nil {
		return Song{}, err
	}
	defer rows.Close()
	r := Song{}
	if rows.Next() {
		err = rows.Scan(&r.Id, &r.Name, &r.ArtistId, &r.AlbumId, &r.TrackNumber, &r.PlayCount, &r.FilePath, &r.CreatedAt, &r.LastModifiedAt)
		if err != nil {
			return Song{}, err
		}
	}
	return r, nil
}

func (d *Dao) FindSongsByArtistId(id int) ([]Song, error) {
	songs := []Song{}
	query := fmt.Sprintf(QuerySongsByArtistId, id)
	rows, err := d.DBConn.Query(query)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	r := Song{}
	for rows.Next() {
		err = rows.Scan(&r.Id, &r.Name, &r.ArtistId, &r.AlbumId, &r.TrackNumber, &r.PlayCount, &r.FilePath, &r.CreatedAt, &r.LastModifiedAt)
		if err != nil {
			return songs, err
		}
		songs = append(songs, r)
	}
	return songs, nil
}

func (d *Dao) FindSongsByAlbumId(id int) ([]Song, error) {
	songs := []Song{}
	query := fmt.Sprintf(QuerySongsByAlbumId, id)
	rows, err := d.DBConn.Query(query)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	r := Song{}
	for rows.Next() {
		err = rows.Scan(&r.Id, &r.Name, &r.ArtistId, &r.AlbumId, &r.TrackNumber, &r.PlayCount, &r.FilePath, &r.CreatedAt, &r.LastModifiedAt)
		if err != nil {
			return songs, err
		}
		songs = append(songs, r)
	}
	return songs, nil
}

func (d *Dao) FindOrCreateSong(track utils.SongMeta, artistId int64, albumId int64, path string, findQuery string, insertQuery string) (int64, error) {
	originalName, cleanedName := santizeString(track.Title)

	id := int64(-1)
	query := fmt.Sprintf(findQuery, cleanedName, artistId, albumId)
	rows, err := d.DBConn.Query(query)
	if err != nil {
		return id, err
	}
	defer rows.Close()
	if rows.Next() {
		//TODO - need to deal with prioritizing matches
		err = rows.Scan(&id)
		if err != nil {
			return id, err
		}
	} else {
		playCount := 0
		trackNumber := track.TrackNumber
		if trackNumber <= 0 {
			re := regexp.MustCompile("^([0-9]*)")
			a := re.Find([]byte(cleanedName))
			trackNumber, _ = strconv.Atoi(string(a))
		}
		query = fmt.Sprintf(insertQuery, originalName, artistId, albumId, trackNumber, playCount, path)
		stmt, err := d.DBConn.Prepare(query)
		if err != nil {
			return id, err
		}
		r, err := stmt.Exec()
		if err != nil {
			return id, err
		}
		id, err = r.LastInsertId()
	}
	return id, nil
}
