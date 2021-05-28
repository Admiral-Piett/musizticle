package daos

type Album struct {
	Id             int
	Name           string
	CreatedAt      string
	LastModifiedAt string
}

func (d *Dao) FetchAllAlbums() ([]Album, error) {
	songs := []Album{}
	rows, err := d.DBConn.Query(QueryAllAlbums)
	if err != nil {
		return songs, err
	}
	defer rows.Close()
	for rows.Next() {
		r := Album{}
		err = rows.Scan(&r.Id, &r.Name, &r.CreatedAt, &r.LastModifiedAt)
		if err != nil {
			return songs, err
		}
		songs = append(songs, r)
	}

	return songs, nil
}
