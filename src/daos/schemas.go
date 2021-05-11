package daos

import (
	"fmt"
	"os"
)

var LastModifiedAtUpdateTrigger = `
	CREATE TRIGGER IF NOT EXISTS lastModifiedAtUpdateTrigger
		BEFORE UPDATE ON %s
	BEGIN
		UPDATE %s SET lastModifiedAt = DATETIME("NOW") WHERE id = old.id;
	END;
`

var LastModifiedAtInsertTrigger = `
	CREATE TRIGGER IF NOT EXISTS lastModifiedAtInsertTrigger
		AFTER INSERT ON %s
	BEGIN
		UPDATE %s SET lastModifiedAt = DATETIME("NOW") WHERE id = new.id;
	END;
`

var createdAtInsertTrigger = `
	CREATE TRIGGER IF NOT EXISTS createdAtInsertTrigger
		AFTER INSERT ON %s
	BEGIN
		UPDATE %s SET createdAt = DATETIME("NOW") WHERE id = new.id;
	END;
`

var ArtistsSchema = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		createdAt INTEGER,
		lastModifiedAt INTEGER
	)
`, os.Getenv("ARTISTS_TABLE"))
