package daos

import (
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

type DatabaseTableInfo struct {
	Cid       int
	Name      string
	Type      string
	NotNull   int
	DfltValue interface{}
	Pk        int
}

type DatabaseTriggerInfo struct {
	Type    string
	Name    string
	TblName string
}

func resetDao() *Dao {
	dao := InitializeDao()
	commands := []string{
		"DELETE FROM artists;",
		"DELETE FROM albums;",
		"DELETE FROM songs;",
		"DELETE FROM users;",
		"VACUUM;",
	}
	for _, c := range commands {
		stmt, err := dao.DBConn.Prepare(c)
		if err != nil {
			panic(err)
		}
		_, err = stmt.Exec()
		if err != nil {
			panic(err)
		}
	}
	return dao
}

func seedArtistOrAlbum(table, value string, dao *Dao) int64 {
	query := ""
	if "artists" == strings.ToLower(table) {
		query = InsertArtist
	} else if "albums" == strings.ToLower(table) {
		query = InsertAlbum
	} else {
		panic(fmt.Errorf("Invalid table: %s", table))
	}

	stmt, _ := dao.DBConn.Prepare(fmt.Sprintf(query, value))
	r, err := stmt.Exec()
	if err != nil {
		panic(err)
	}
	id, _ := r.LastInsertId()
	return id
}

func seedUser(username, password string, dao *Dao) int64 {
	query := fmt.Sprintf(InsertUsers, username, password)
	stmt, _ := dao.DBConn.Prepare(query)
	r, err := stmt.Exec()
	if err != nil {
		panic(err)
	}
	id, _ := r.LastInsertId()
	return id
}

func seedSong(name string, dao *Dao, artistId, albumId int64, filePath string) int64 {
	if artistId == 0 || albumId == 0 {
		//Seed these every time to make sure that the test ids have valid foreign keys
		artistId = seedArtistOrAlbum("artists", "test-artist", dao)
		albumId = seedArtistOrAlbum("albums", "test-album", dao)
	}
	if filePath == "" {
		filePath = "/file/path"
	}

	query := fmt.Sprintf(InsertSongs, name, artistId, albumId, 1, 0, filePath, 300)
	stmt, _ := dao.DBConn.Prepare(query)
	r, err := stmt.Exec()
	if err != nil {
		panic(err)
	}
	id, _ := r.LastInsertId()
	return id
}

func TestMain(m *testing.M) {
	models.SETTINGS.SqliteDB = "dao_tests.db"
	models.SETTINGS.SqliteDriver = "sqlite3"

	_, file, _, _ := runtime.Caller(0)
	projectDirectory := filepath.Join(filepath.Dir(file), "../..")
	dbPath := fmt.Sprintf("%s/data/%s", projectDirectory, models.SETTINGS.SqliteDB)
	defer os.Remove(dbPath)

	code := m.Run()
	os.Exit(code)
}

func Test_InitializeDao_success(t *testing.T) {
	dao := InitializeDao()

	tests := []struct {
		table        string
		expectedInfo []DatabaseTableInfo
	}{
		{"albums", []DatabaseTableInfo{
			{Cid: 0, Name: "id", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 1},
			{Cid: 1, Name: "name", Type: "TEXT", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 2, Name: "createdAt", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 3, Name: "lastModifiedAt", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
		}},
		{"artists", []DatabaseTableInfo{
			{Cid: 0, Name: "id", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 1},
			{Cid: 1, Name: "name", Type: "TEXT", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 2, Name: "createdAt", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 3, Name: "lastModifiedAt", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
		}},
		{"songs", []DatabaseTableInfo{
			{Cid: 0, Name: "id", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 1},
			{Cid: 1, Name: "name", Type: "TEXT", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 2, Name: "artistId", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 3, Name: "albumId", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 4, Name: "trackNumber", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 5, Name: "playCount", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 6, Name: "filePath", Type: "TEXT", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 7, Name: "duration", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 8, Name: "createdAt", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 9, Name: "lastModifiedAt", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
		}},
		{"users", []DatabaseTableInfo{
			{Cid: 0, Name: "id", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 1},
			{Cid: 1, Name: "username", Type: "TEXT", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 2, Name: "password", Type: "TEXT", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 3, Name: "createdAt", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
			{Cid: 4, Name: "lastModifiedAt", Type: "INTEGER", NotNull: 0, DfltValue: nil, Pk: 0},
		}},
	}
	for _, test := range tests {
		results, err := dao.DBConn.Query(fmt.Sprintf("PRAGMA table_info(%s);", test.table))

		index := 0
		for results.Next() {
			r := DatabaseTableInfo{}
			err = results.Scan(&r.Cid, &r.Name, &r.Type, &r.NotNull, &r.DfltValue, &r.Pk)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedInfo[index], r)
			index++
		}
	}
}

func Test_InitializeDao_success_triggers(t *testing.T) {
	dao := InitializeDao()

	tests := []struct {
		expectedInfo []DatabaseTriggerInfo
	}{
		{[]DatabaseTriggerInfo{
			{Type: "trigger", Name: "last_modified_at_update_trigger_albums", TblName: "albums"},
			{Type: "trigger", Name: "last_modified_at_insert_trigger_albums", TblName: "albums"},
			{Type: "trigger", Name: "created_at_insert_trigger_albums", TblName: "albums"},
			{Type: "trigger", Name: "last_modified_at_update_trigger_artists", TblName: "artists"},
			{Type: "trigger", Name: "last_modified_at_insert_trigger_artists", TblName: "artists"},
			{Type: "trigger", Name: "created_at_insert_trigger_artists", TblName: "artists"},
			{Type: "trigger", Name: "last_modified_at_update_trigger_songs", TblName: "songs"},
			{Type: "trigger", Name: "last_modified_at_insert_trigger_songs", TblName: "songs"},
			{Type: "trigger", Name: "created_at_insert_trigger_songs", TblName: "songs"},
			{Type: "trigger", Name: "last_modified_at_update_trigger_users", TblName: "users"},
			{Type: "trigger", Name: "last_modified_at_insert_trigger_users", TblName: "users"},
			{Type: "trigger", Name: "created_at_insert_trigger_users", TblName: "users"},
		}},
	}
	for _, test := range tests {
		// I don't care about all the values in here, this table stores the raw sql and everything else.  I don't
		//  need to assert that I can write SQL and it works - that's a smoke test's job.
		results, err := dao.DBConn.Query("select type, name, tbl_name from sqlite_master where type = 'trigger';")

		index := 0
		for results.Next() {
			r := DatabaseTriggerInfo{}
			err = results.Scan(&r.Type, &r.Name, &r.TblName)

			assert.Nil(t, err)
			assert.Equal(t, test.expectedInfo[index], r)
			index++
		}
	}
}

func Test_InitializeDao_can_not_open_db_panic(t *testing.T) {
	models.SETTINGS.SqliteDB = "/garbage/path"
	models.SETTINGS.SqliteDriver = "crap-driver"
	_, file, _, _ := runtime.Caller(0)
	projectDirectory := filepath.Join(filepath.Dir(file), "../..")
	dbPath := fmt.Sprintf("%s/data/%s", projectDirectory, models.SETTINGS.SqliteDB)
	defer os.Remove(dbPath)

	assert.Panics(t, func() {
		InitializeDao()
	})

	// Clean up
	models.SETTINGS.SqliteDB = "dao_tests.db"
	models.SETTINGS.SqliteDriver = "sqlite3"
}

func Test_CloseDao_success(t *testing.T) {
	dao := resetDao()

	err := dao.DBConn.Ping()
	assert.Nil(t, err)
	dao.CloseDao()
	err = dao.DBConn.Ping()
	assert.Error(t, err)
}

func Test_FindOrCreateByName_finds_and_returns_data(t *testing.T) {
	dao := resetDao()
	seedArtistOrAlbum("artists", "Hansy Zimmerino", dao)

	artistId, err := dao.FindOrCreateByName("Hansy Zimmerino", QueryArtistIdByName, InsertArtist)

	// Since we're not recreating the db (and stupid library didn't make any interfaces, grr...) we can't be sure
	//  what id each entry is going to have.  Since we don't want to repurpose the ids.  If we hadn't done anything
	//  this would have been -1 so this is enough for now.
	assert.Greater(t, artistId, int64(0))
	assert.Nil(t, err)
}

func Test_FindOrCreateByName_creates_data(t *testing.T) {
	dao := resetDao()
	artistId, err := dao.FindOrCreateByName("Hansy Zimmerino", QueryArtistIdByName, InsertArtist)

	assert.Greater(t, artistId, int64(0))
	assert.Nil(t, err)
}

func Test_FindOrCreateByName_query_error_returns_error(t *testing.T) {
	dao := resetDao()
	artistId, err := dao.FindOrCreateByName("Hansy Zimmerino", "not-a-query", InsertArtist)

	assert.Equal(t, int64(-1), artistId)
	assert.Error(t, err)
}

func Test_FindOrCreateByName_scan_error_returns_error(t *testing.T) {
	dao := resetDao()
	seedArtistOrAlbum("artists", "Hansy Zimmerino", dao)
	// This query would find something, but it shouldn't match what we are trying to retrieve and give us an error.
	query := `select * from artists WHERE name LIKE "%%%s%%";`
	artistId, err := dao.FindOrCreateByName("Hansy Zimmerino", query, InsertArtist)

	assert.Equal(t, int64(-1), artistId)
	assert.Error(t, err)
}

func Test_FindOrCreateByName_prepare_error_returns_error(t *testing.T) {
	dao := resetDao()
	// This query would find something, but it shouldn't match what we are trying to retrieve and give us an error.
	query := `INSERT INTO artists(name)	values("";`
	artistId, err := dao.FindOrCreateByName("Hansy Zimmerino", QueryArtistIdByName, query)

	assert.Equal(t, int64(-1), artistId)
	assert.Error(t, err)
}

func Test_santizeString_returns_already_clean_value(t *testing.T) {
	original, cleaned := santizeString("my test")

	assert.Equal(t, "my test", original)
	assert.Equal(t, "my test", cleaned)
}

func Test_santizeString_strips_first_and_last_extra_words(t *testing.T) {
	original, cleaned := santizeString("the test ost")

	assert.Equal(t, "the test ost", original)
	assert.Equal(t, "test", cleaned)
}

func Test_santizeString_returns_unknowns(t *testing.T) {
	original, cleaned := santizeString("")

	assert.Equal(t, "UNKNOWN", original)
	assert.Equal(t, "UNKNOWN", cleaned)
}
