package utils

import (
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/models"
	"io"

	"github.com/dhowden/tag"
	"github.com/tcolgate/mp3"
)


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

func GetSongMetadata(file io.ReadSeeker) (models.SongMeta, error) {
	track, err := tag.ReadFrom(file)
	if err != nil {
		return models.SongMeta{}, err
	}
	trackNumber, totalTracks := track.Track()
	discNumber, totalDiscs := track.Disc()
	duration, err := getTime(file)
	if err != nil {
		return models.SongMeta{}, err
	}
	song_meta := models.SongMeta{
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
