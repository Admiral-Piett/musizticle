package daos

import (
	"fmt"
)

func (d *Dao) FindOrCreateSong(name string, artistId int64, albumId int64, path string, findQuery string, insertQuery string, sanitize bool) (int64, error) {
	if sanitize {
		name = santizeString(name)
	}
	id := int64(-1)
	query := fmt.Sprintf(findQuery, name, artistId, albumId)
	rows, err := d.DBConn.Query(query)
	defer rows.Close()
	if err != nil {
		return id, err
	}
	if rows.Next() {
		//TODO - need to deal with prioritizing matches
		err = rows.Scan(&id)
		if err != nil {
			return id, err
		}
	} else {
		query = fmt.Sprintf(insertQuery, name, artistId, albumId, 0, path)
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
