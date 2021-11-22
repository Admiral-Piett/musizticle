package utils

import (
	"fmt"
	"io"

	"github.com/dhowden/tag"
	"github.com/tcolgate/mp3"
)

//TODO - environmentalize
var PORT string = "9000"
var SQLITE_DB string = "musizticle.db"

//TODO - Use me
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

func getTime(file io.ReadSeeker) (int, error) {
	// FIXME - should I not add in extra?  I could get more exact I suppose?
	t := 1.0

	d := mp3.NewDecoder(file)
	var f mp3.Frame
	skipped := 0

	for {

		if err := d.Decode(&f, &skipped); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			return int(t), err
		}

		t = t + f.Duration().Seconds()
	}
	return int(t), nil
}

func GetSongMetadata(file io.ReadSeeker) (SongMeta, error) {
	track, err := tag.ReadFrom(file)
	if err != nil {
		return SongMeta{}, err
	}
	trackNumber, totalTracks := track.Track()
	discNumber, totalDiscs := track.Disc()
	duration, err := getTime(file)
	if err != nil {
		return SongMeta{}, err
	}
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
		Duration:     duration,
	}
	return song_meta, nil
}
