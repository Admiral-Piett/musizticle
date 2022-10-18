package daos

import (
	"database/sql"
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/models"
	"regexp"
	"strconv"
)

func scanSong(rows *sql.Rows, r *models.ListSong, set_names bool) error {
	if !set_names {
		return rows.Scan(&r.Id, &r.Title, &r.ArtistId, &r.AlbumId, &r.TrackNumber, &r.PlayCount, &r.FilePath, &r.Duration, &r.CreatedAt, &r.LastModifiedAt)
	}
	return rows.Scan(&r.Id, &r.Title, &r.ArtistId, &r.AlbumId, &r.TrackNumber, &r.PlayCount, &r.FilePath, &r.Duration, &r.CreatedAt, &r.LastModifiedAt, &r.ArtistName, &r.AlbumName)
}

func (d *Dao) GetAllSongs() ([]models.ListSong, error) {
	songs := []models.ListSong{}
	rows, err := d.DBConn.Query(QueryAllSongs)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	for rows.Next() {
		r := models.ListSong{}
		err = scanSong(rows, &r, true)
		if err != nil {
			return songs, err
		}
		songs = append(songs, r)
	}

	return songs, nil
}

func (d *Dao) FindSongById(id int) (models.ListSong, error) {
	query := fmt.Sprintf(QuerySongById, id)
	rows, err := d.DBConn.Query(query)
	if err != nil {
		return models.ListSong{}, err
	}
	defer rows.Close()
	r := models.ListSong{}
	if rows.Next() {
		err = scanSong(rows, &r, false)
		if err != nil {
			return models.ListSong{}, err
		}
	}
	return r, nil
}

func (d *Dao) FindSongsByArtistId(id int) ([]models.ListSong, error) {
	songs := []models.ListSong{}
	query := fmt.Sprintf(QuerySongsByArtistId, id)
	rows, err := d.DBConn.Query(query)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	r := models.ListSong{}
	for rows.Next() {
		err = scanSong(rows, &r, false)
		if err != nil {
			return songs, err
		}
		songs = append(songs, r)
	}
	return songs, nil
}

// FIXME - Currently unused
func (d *Dao) FindSongsByAlbumId(id int) ([]models.ListSong, error) {
	songs := []models.ListSong{}
	query := fmt.Sprintf(QuerySongsByAlbumId, id)
	rows, err := d.DBConn.Query(query)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	r := models.ListSong{}
	for rows.Next() {
		err = scanSong(rows, &r, false)
		if err != nil {
			return songs, err
		}
		songs = append(songs, r)
	}
	return songs, nil
}

func (d *Dao) FindOrCreateSong(track models.SongMeta, artistId int64, albumId int64, path string, duration int, findQuery string, insertQuery string) (int64, error) {
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
		query = fmt.Sprintf(insertQuery, originalName, artistId, albumId, trackNumber, playCount, path, duration)
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
