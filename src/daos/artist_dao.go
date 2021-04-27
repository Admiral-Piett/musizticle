package daos

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

//TODO - environmentalize
var (
	dbUsername         = "secret"
	dbPassword         = "secret"
	dbHost             = "secret"
	dbTable            = "secret"
	dbPort             = "secret"
	dbConnectionString = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbTable, dbPassword)
)

type ArtistDao struct {
	db *sqlx.DB
}

var schema = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s 
	(
		id SERIAL PRIMARY KEY,
		name TEXT,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)
`, dbTable)

func (d *ArtistDao) Open() error {
	pg, err := sqlx.Open(dbTable, dbConnectionString)
	if err != nil {
		return err
	}
	log.Printf("Connected to database - %s", dbTable)

	pg.MustExec(schema)

	d.db = pg
	return nil
}

func (d *ArtistDao) Close() error {
	return d.db.Close()
}
