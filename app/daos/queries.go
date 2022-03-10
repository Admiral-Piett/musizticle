package daos

// --- Albumn

var InsertAlbum = `
INSERT INTO 
	albums(name)
	values("%s")
`

var QueryAllAlbums = `
SELECT *
FROM albums
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

var QueryAllArtists = `
SELECT *
FROM artists
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
	songs(name, artistId, albumId, trackNumber, playCount, filePath, duration)
	values("%s", %d, %d, %d, %d, "%s", %d)
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

var QuerySongsByAlbumId = `
SELECT * FROM songs WHERE artistId = %d
`

var QuerySongsByArtistId = `
SELECT * FROM songs WHERE artistId = %d
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

var QueryUserByUsername = `
SELECT
       *
FROM
	 users
WHERE username = "%s" 
AND password = "%s"
`
