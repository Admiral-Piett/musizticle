package utils

import (
	"github.com/dhowden/tag"
	"io"
)

//TODO - environmentalize
var PORT string = "9000"
var SQLITE_DB string = "sound_control.db"

var InvalidFileTypes = []string{
	".DS_Store",
	".7z",
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
	AlbumId string
	ArtistId string
	ErrorMessage string
	FilePath     string
	SongID       string
	RequestBody  string
	Size string
	StackContext string
}

var LogFields = LogFieldStruct{
	AlbumId: "album_id",
	ArtistId: "artist_id",
	ErrorMessage: "error_message",
	FilePath:     "file_path",
	SongID:       "song_id",
	RequestBody:  "request_body",
	Size: "size",
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
}

func GetSongMetadata(file io.ReadSeeker) (SongMeta, error) {
	track, err := tag.ReadFrom(file)
	if err != nil {
		return SongMeta{}, err
	}
	trackNumber, totalTracks := track.Track()
	discNumber, totalDiscs := track.Disc()
	song_meta := SongMeta{
		Title:        track.Title(),
		Album:        track.Album(),
		Artist:       track.Artist(),
		AlbumnArtist: track.AlbumArtist(),
		Composer:     track.Composer(),
		Year:         track.Year(),
		Genre:        track.Genre(),
		TrackNumber:  trackNumber,
		TotalTracks:  totalTracks,
		Disc:         discNumber,
		TotalDiscs:   totalDiscs,
		Picture:      track.Picture(),
		Lyrics:       track.Lyrics(),
		Comment:      track.Comment(),
		Format:       track.Format(),
	}
	return song_meta, nil
}
