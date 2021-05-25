package daos

// --- Albumn

var InsertAlbum = `
INSERT INTO 
	albums(name)
	values("%s")
`

var QueryAlbumIdByName = `
SELECT
       id
FROM
	 albums
WHERE
	name LIKE "%%%s%%"
`

// --- Artist

var InsertArtist = `
INSERT INTO 
	artists(name)
	values("%s")
`

var QueryArtistIdByName = `
SELECT
       id
FROM
	 artists
WHERE
	name LIKE "%%%s%%"
`

// --- Songs
var InsertSongs = `
INSERT INTO 
	songs(name, artistId, albumId, playCount, filePath)
	values("%s", %d, %d, %d, "%s")
`

var QuerySongIdByName = `
SELECT
       id
FROM
	 songs
WHERE name LIKE "%%%s%%"
AND artistId = %d
AND albumId = %d
`
