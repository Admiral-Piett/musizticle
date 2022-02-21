package models

import "github.com/dhowden/tag"

//TODO - environmentalize
var PORT string = "9000"
var SQLITE_DB string = "musizticle.db"

//TODO - Use me
var InvalidFileTypes = []string{
	".DS_Store",
	".7z",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// FIXME - there has to be a more elegant object for this
type TablesStruct struct {
	Albums  string
	Artists string
	Songs   string
}

var Tables = TablesStruct{
	Albums:  "albums",
	Artists: "artists",
	Songs:   "songs",
}

type LogFieldStruct = struct {
	AlbumId      string
	ArtistId     string
	ErrorMessage string
	FilePath     string
	SongID       string
	RequestBody  string
	Size         string
	StackContext string
}

var LogFields = LogFieldStruct{
	AlbumId:      "album_id",
	ArtistId:     "artist_id",
	ErrorMessage: "error_message",
	FilePath:     "file_path",
	SongID:       "song_id",
	RequestBody:  "request_body",
	Size:         "size",
	StackContext: "stack_context",
}

type SongMeta struct {
	Title        string
	Album        string
	Artist       string
	AlbumnArtist string
	Composer     string
	Year         int
	Genre        string
	TrackNumber  int
	TotalTracks  int
	Disc         int
	TotalDiscs   int
	Picture      *tag.Picture
	Lyrics       string
	Comment      string
	Format       tag.Format
	Duration     int
}
