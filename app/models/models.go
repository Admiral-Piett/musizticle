package models

import (
	"crypto/rsa"
	"github.com/dhowden/tag"
)

var Settings struct {
	Port string
	SqliteDB string
	PublicKey *rsa.PublicKey
	PrivateKey *rsa.PrivateKey
	TokenExpiration int
}

var SETTINGS = Settings

//TODO - Use me
var InvalidFileTypes = []string{
	".DS_Store",
	".7z",
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// FIXME - there has to be a more elegant object for this
type TablesStruct struct {
	Albums  string
	Artists string
	Songs   string
	Users   string
}

var Tables = TablesStruct{
	Albums:  "albums",
	Artists: "artists",
	Songs:   "songs",
	Users:   "users",
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

// ----- User

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id             int
	Username       string
	Password       string
	CreatedAt      string
	LastModifiedAt string
}

type JwtToken struct {
	UserId []byte `json:"userId"`
	CreatedAt []byte `json:"createdAt"`
	ExpiresAt []byte `json:"expiresAt"`
}

// ----- Media

type Artist struct {
	Id             int
	Name           string
	CreatedAt      string
	LastModifiedAt string
}

type Album struct {
	Id             int
	Name           string
	CreatedAt      string
	LastModifiedAt string
}

type ListSong struct {
	Id          int
	Title       string
	ArtistId    int
	ArtistName  string
	AlbumId     int
	AlbumName   string
	TrackNumber int
	PlayCount   int
	FilePath    string
	Duration    int
	//FIXME - wtf, these are strings??
	CreatedAt      string
	LastModifiedAt string
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
