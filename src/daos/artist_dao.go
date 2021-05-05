package daos

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

//TODO - environmentalize
var (
	dbTable            = os.Getenv("ARTISTS_TABLE")
)

type ArtistsDao struct {
	artists *sqlx.DB
}

func (d *ArtistsDao) Open() error {
	dbConnectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUsername, dbName, dbPassword)
	pg, err := sqlx.Connect("postgres", dbConnectionString)
	if err != nil {
		return err
	}
	log.Printf("Connected to database - %s", dbTable)

	pg.MustExec(ArtistsSchema)

	d.artists = pg
	return nil
}

func (d *ArtistsDao) Close() error {
	return d.artists.Close()
}
