package daos

type Artist struct {
	Id             int
	Name           string
	CreatedAt      string
	LastModifiedAt string
}

func (d *Dao) FetchAllArtists() ([]Artist, error) {
	songs := []Artist{}
	rows, err := d.DBConn.Query(QueryAllArtists)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	for rows.Next() {
		r := Artist{}
		err = rows.Scan(&r.Id, &r.Name, &r.CreatedAt, &r.LastModifiedAt)
		if err != nil {
			return songs, err
		}
		songs = append(songs, r)
	}

	return songs, nil
}
