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
	TokenKey []byte
}

var SETTINGS = Settings

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type AuthResponse struct {
	AuthToken string `json:"authToken"`
	ReauthToken string `json:"reauthToken"`
	ExpirationTime string `json:"expirationTime"`
}

type TablesStruct struct {
	Albums  string
	Artists string
	Songs   string
	Users   string
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

// ----- Canned Responses
var UnauthorizedResponse = ErrorResponse{Code: "UNAUTHORIZED", Message: "Unauthorized"}



