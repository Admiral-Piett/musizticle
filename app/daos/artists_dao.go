package daos

import "github.com/Admiral-Piett/musizticle/app/models"

func (d *Dao) GetAllArtists() ([]models.Artist, error) {
	songs := []models.Artist{}
	rows, err := d.DBConn.Query(QueryAllArtists)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	for rows.Next() {
		r := models.Artist{}
		err = rows.Scan(&r.Id, &r.Name, &r.CreatedAt, &r.LastModifiedAt)
		if err != nil {
			return songs, err
		}
		songs = append(songs, r)
	}

	return songs, nil
}
