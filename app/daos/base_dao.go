package daos

import (
	"database/sql"
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/models"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var Tables = models.TablesStruct{
	Albums:  "albums",
	Artists: "artists",
	Songs:   "songs",
	Users:   "users",
}

var nonSearchableStrings = map[string]bool{
	"the":        true,
	"a":          true,
	"ost":        true,
	"soundtrack": true,
	"score":      true,
}

type Dao struct {
	DBConn *sql.DB
}

var schemas = map[string]string{
	Tables.Albums:  AlbumnSchema,
	Tables.Artists: ArtistsSchema,
	Tables.Songs:   SongsSchema,
	Tables.Users:   UsersSchema,
}

func InitializeDao() *Dao {
	_, file, _, _ := runtime.Caller(0)
	projectDirectory := filepath.Join(filepath.Dir(file), "../..")
	os.Mkdir(fmt.Sprintf("%s/data", projectDirectory), 0755)
	db, err := sql.Open(models.SETTINGS.SqliteDriver, fmt.Sprintf("%s/data/%s", projectDirectory, models.SETTINGS.SqliteDB))
	if err != nil {
		panic(err)
	}
	dao := &Dao{
		DBConn: db,
	}
	dao.setupTables()
	return dao
}

func (d *Dao) setupTables() {
	for table, schema := range schemas {
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
	trigger := fmt.Sprintf(LastModifiedAtUpdateTrigger, table, table, table)
	stmt, err := d.DBConn.Prepare(trigger)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
	// Add auto-lastModifiedAt on row INSERT
	trigger = fmt.Sprintf(LastModifiedAtInsertTrigger, table, table, table)
	stmt, err = d.DBConn.Prepare(trigger)
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
	// Add auto-createdAt on row INSERT
	trigger = fmt.Sprintf(createdAtInsertTrigger, table, table, table)
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

func (d *Dao) FindOrCreateByName(name, findQuery, insertQuery string) (int64, error) {
	originalName, cleanedName := santizeString(name)

	id := int64(-1)
	query := fmt.Sprintf(findQuery, cleanedName)
	rows, err := d.DBConn.Query(query)
	if err != nil {
		return id, err
	}
	defer rows.Close()
	if rows.Next() {
		//TODO - need to deal with prioritizing matches
		err = rows.Scan(&id)
		if err != nil {
			return id, err
		}
	} else {
		query = fmt.Sprintf(insertQuery, originalName)
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

func santizeString(originalValue string) (string, string) {
	if originalValue == "" {
		return "UNKNOWN", "UNKNOWN"
	}
	value := originalValue
	s := strings.Split(value, " ")
	if nonSearchableStrings[strings.ToLower(s[0])] {
		s = s[1:]
	}
	if nonSearchableStrings[strings.ToLower(s[len(s)-1])] {
		s = s[:len(s)-1]
	}
	//TODO - make this more strict on invalid characters
	//Strip out invalid values from both the originalValue and the cleaned one, as both will interact with the database.
	cleaned := []string{}
	for _, v := range s {
		v = strings.Replace(v, "\"", "`", -1)
		cleaned = append(cleaned, v)
	}
	originalValue = strings.Replace(originalValue, "\"", "`", -1)

	cleanedStr := strings.Join(cleaned, " ")
	return originalValue, cleanedStr
}
