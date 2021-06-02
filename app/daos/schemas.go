package daos

// ---- Triggers
var createdAtInsertTrigger = `
	CREATE TRIGGER IF NOT EXISTS created_at_insert_trigger_%s
		AFTER INSERT ON %s
	BEGIN
		UPDATE %s SET createdAt = DATETIME("NOW") WHERE id = new.id;
	END;
`

var LastModifiedAtInsertTrigger = `
	CREATE TRIGGER IF NOT EXISTS last_modified_at_insert_trigger_%s
		AFTER INSERT ON %s
	BEGIN
		UPDATE %s SET lastModifiedAt = DATETIME("NOW") WHERE id = new.id;
	END;
`

var LastModifiedAtUpdateTrigger = `
	CREATE TRIGGER IF NOT EXISTS last_modified_at_update_trigger_%s
		BEFORE UPDATE ON %s
	BEGIN
		UPDATE %s SET lastModifiedAt = DATETIME("NOW") WHERE id = old.id;
	END;
`


// ---- Schemas
var AlbumnSchema = `
	CREATE TABLE IF NOT EXISTS albums (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		createdAt INTEGER,
		lastModifiedAt INTEGER
	)
`

var ArtistsSchema = `
	CREATE TABLE IF NOT EXISTS artists (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		createdAt INTEGER,
		lastModifiedAt INTEGER
	)
`

var SongsSchema = `
	CREATE TABLE IF NOT EXISTS songs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		artistId INTEGER,
		albumId INTEGER,
		trackNumber INTEGER,
		playCount INTEGER,
		filePath TEXT UNIQUE,
		createdAt INTEGER,
		lastModifiedAt INTEGER,
		FOREIGN KEY(albumId) REFERENCES albums(id),
		FOREIGN KEY(artistId) REFERENCES artists(id)
	)
`
