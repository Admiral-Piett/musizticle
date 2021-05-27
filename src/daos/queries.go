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
	songs(name, artistId, albumId, trackNumber, playCount, filePath)
	values("%s", %d, %d, %d, %d, "%s")
`

var QueryAllSongs = `
SELECT s.*,
       a.name,
       al.name
FROM songs s
JOIN artists a ON s.artistId = a.id
JOIN albums al ON s.albumId = al.id
`

var QuerySongById = `
SELECT * FROM songs WHERE id = %d
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
