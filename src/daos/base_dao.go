package daos

import (
	"database/sql"
	"fmt"
	"github.com/Admiral-Piett/sound_control/src/utils"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Dao struct {
	DBConn     *sql.DB
}

var schemas = map[string]string{
	utils.Tables.Albums:  AlbumnSchema,
	utils.Tables.Artists: ArtistsSchema,
	utils.Tables.Songs:   SongsSchema,
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


func (d *Dao) FindOrCreateByName(name string, findQuery string, insertQuery string, sanitize bool) (int64, error) {
	if sanitize {
		name = santizeString(name)
	}
	id := int64(-1)
	query := fmt.Sprintf(findQuery, name)
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
		query = fmt.Sprintf(insertQuery, name)
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

var nonSearchableStrings = map[string]bool{
	"the": true,
	"a": true,
	"ost": true,
	"soundtrack": true,
	"score": true,
}

func santizeString(value string) string {
	s := strings.Split(value, " ")
	if nonSearchableStrings[strings.ToLower(s[0])] {
		s = s[1:]
	}
	if nonSearchableStrings[strings.ToLower(s[len(s)-1])] {
		s = s[:len(s)-1]
	}
	cleaned := []string{}
	for _, v := range s {
		v = strings.Replace(v,"\"", "`", -1)
		cleaned = append(cleaned, v)
	}
	cleanedStr := strings.Join(cleaned, " ")
	if cleanedStr == "" {
		cleanedStr = "UNKNOWN"
	}
	return cleanedStr
}
