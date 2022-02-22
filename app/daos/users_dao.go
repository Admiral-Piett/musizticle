package daos

import (
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/models"
)

func (d *Dao) GetUser(username, password string) (models.User, error) {
	user := models.User{}
	query := fmt.Sprintf(QueryUserByUsername, username, password)

	row, err := d.DBConn.Query(query)
	if err != nil {
		return user, err
	}
	defer row.Close()

	// We should only ever have one here, so we'll just grab next.
	row.Next()
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.CreatedAt, &user.LastModifiedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
