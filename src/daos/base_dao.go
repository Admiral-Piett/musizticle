package daos

import "os"

var (
	dbUsername = os.Getenv("POSTGRES_USER")
	dbPassword = os.Getenv("POSTGRES_PASSWORD")
	dbHost = os.Getenv("POSTGRES_HOST")
	dbPort = os.Getenv("POSTGRES_PORT")
	dbName = os.Getenv("POSTGRES_DB")
)

type PostgresDao interface {
	Open() error
	Close() error
}

type Daos struct {
	ArtistsDao *ArtistsDao
}

func InitializeDaos() *Daos {
	daos := &Daos{
		ArtistsDao: &ArtistsDao{},
	}
	daos.openAllDaos()
	return daos
}

func (d *Daos) openAllDaos() {
	// TODO - do this for every table
	// TODO - there HAS to be a more elegant way to do this?
	if err := d.ArtistsDao.Open(); err != nil {
		panic(err)
	}
}


func (d *Daos) CloseAllDaos() {
	if err := d.ArtistsDao.Close(); err != nil {
		panic(err)
	}
}
