package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"github.com/Admiral-Piett/musizticle/app/models"
	"io"
	"strconv"

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

var random_chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateRandomString(length int) string {
	if length <= 0 {
		return ""
	}
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	l := len(random_chars)
	for i, v := range(b) {
		// Get the value of the random 8 bit int modulo the max index of the random_chars.  This way we'll use the
		//random int to indicate a value from this list, and cap the possible values at the last index.
		m := int(v) % l
		b[i] = random_chars[m]
	}
	return string(b)
}

func Encrypt(value int) ([]byte, error) {
	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		models.SETTINGS.PublicKey,
		[]byte(strconv.Itoa(value)),
		nil)
	if err != nil {
		return []byte{}, err
	}
	return encryptedBytes, nil
}

func Decrypt(value []byte) (int, error) {
	decryptedBytes, err := models.SETTINGS.PrivateKey.Decrypt(nil, value, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		return 0, err
	}
	s, _ := strconv.Atoi(string(decryptedBytes))
	return s, nil
}

