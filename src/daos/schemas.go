package daos

import "fmt"

var ArtistsSchema = fmt.Sprintf(`
	CREATE TABLE IF NOT EXISTS %s 
	(
		id SERIAL PRIMARY KEY,
		name TEXT,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	)
`, dbTable)
