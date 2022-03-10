package daos

import "github.com/Admiral-Piett/musizticle/app/models"

func (d *Dao) GetAllAlbums() ([]models.Album, error) {
	songs := []models.Album{}
	rows, err := d.DBConn.Query(QueryAllAlbums)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	for rows.Next() {
		r := models.Album{}
		err = rows.Scan(&r.Id, &r.Name, &r.CreatedAt, &r.LastModifiedAt)
		if err != nil {
			return songs, err
		}
		songs = append(songs, r)
	}

	return songs, nil
}
