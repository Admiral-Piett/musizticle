package utils

import (
	"github.com/dhowden/tag"
	"io"
)

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
	ErrorMessage string
	FilePath string
	RequestBody string
	StackContext string
}

var LogFields = LogFieldStruct{
	ErrorMessage: "error_message",
	FilePath: "file_path",
	RequestBody: "request_body",
	StackContext: "stack_context",
}

type SongMeta struct {
	Title string
	Album string
	Artist string
	AlbumnArtist string
	Composer string
	Year int
	Genre string
	Track int
	TotalTracks int
	Disc int
	TotalDiscs int
	Picture *tag.Picture
	Lyrics string
	Comment string
	Format tag.Format
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
		Track:        trackNumber,
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