package daos

import (
	"fmt"
	"github.com/Admiral-Piett/sound_control/src/utils"
	"regexp"
	"strconv"
)

type ListSong struct {
	Id             int
	Name           string
}

type Song struct {
	Id             int
	Name           string
	ArtistId       int
	AlbumId        int
	TrackNumber    int
	PlayCount      int
	FilePath       string
	// FIXME - wtf, these are strings??
	CreatedAt      string
	LastModifiedAt string
}

func (d *Dao) FetchAllSongs() ([]Song, error) {
	songs := []Song{}
	rows, err := d.DBConn.Query(QueryAllSongs)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	for rows.Next() {
		r := Song{}
		err = rows.Scan(&r.Id, &r.Name, &r.ArtistId, &r.AlbumId, &r.TrackNumber, &r.PlayCount, &r.FilePath, &r.CreatedAt, &r.LastModifiedAt)
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

func (d *Dao) FindOrCreateSong(track utils.SongMeta, artistId int64, albumId int64, path string, findQuery string, insertQuery string, sanitize bool) (int64, error) {
	name := track.Title
	if sanitize {
		name = santizeString(name)
	}
	id := int64(-1)
	query := fmt.Sprintf(findQuery, name, artistId, albumId)
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
		trackNumber := track.TrackNumber
		if trackNumber <= 0 {
			re := regexp.MustCompile("^([0-9]*)")
			a := re.Find([]byte(name))
			trackNumber, _ = strconv.Atoi(string(a))
		}
		query = fmt.Sprintf(insertQuery, track.Title, artistId, albumId, trackNumber, 0, path)
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
