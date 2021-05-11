package daos

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"runtime"
)

type Dao struct {
	DBConn     *sql.DB
	ArtistsTable string
}

var schemas = map[string]string{
	os.Getenv("ARTISTS_TABLE"): ArtistsSchema,
}

func InitializeDao() *Dao {
	_, file, _, _ := runtime.Caller(0)
	projectDirectory := filepath.Join(filepath.Dir(file), "../..")
	os.Mkdir(fmt.Sprintf("%s/data", projectDirectory), 0755)
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/data/%s", projectDirectory, os.Getenv("SQL_LITE_FILE")))
	if err != nil {
		panic(err)
	}
	dao := &Dao{
		DBConn: db,
		ArtistsTable: os.Getenv("ARTISTS_TABLE"),
	}
	dao.setupTables()
	return dao
}

func (d *Dao) setupTables() {
	for table, schema := range(schemas) {
		stmt, err := d.DBConn.Prepare(schema)
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec()
		if err != nil {
			panic(err)
		}
		d.setDatetimeTriggers(table)
	}

}

func (d *Dao) setDatetimeTriggers(table string) {
	// Add auto-lastModifiedAt on row UPDATE
	trigger := fmt.Sprintf(LastModifiedAtUpdateTrigger, table, table)
	stmt, err := d.DBConn.Prepare(trigger)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
	// Add auto-lastModifiedAt on row INSERT
	trigger = fmt.Sprintf(LastModifiedAtInsertTrigger, table, table)
	stmt, err = d.DBConn.Prepare(trigger)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
	// Add auto-createdAt on row INSERT
	trigger = fmt.Sprintf(createdAtInsertTrigger, table, table)
	stmt, err = d.DBConn.Prepare(trigger)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

func (d *Dao) CloseDao() {
	if err := d.DBConn.Close(); err != nil {
		panic(err)
	}
}
