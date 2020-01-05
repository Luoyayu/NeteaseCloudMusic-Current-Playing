package main

import (
	NetEaseMusicPlaying "163musicPlaying"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

func Playing(c *gin.Context) {
	NetEaseMusicPlaying.PlayingMutex.Lock()
	var playingSong = NetEaseMusicPlaying.Playing
	NetEaseMusicPlaying.PlayingMutex.Unlock()

	var artists []string
	for _, a := range playingSong.Track.Artists {
		artists = append(artists, a.Name)
	}

	duration, _ := time.ParseDuration(fmt.Sprint(playingSong.Track.Duration/1000) + "s")

	c.JSONP(http.StatusOK, map[string]interface{}{
		"songId":         playingSong.Track.Id,
		"songName":       playingSong.Track.Name,
		"songAlias":      playingSong.Track.Alias,
		"songPopularity": playingSong.Track.Popularity,
		"songIsPayed":    playingSong.Track.Privilege.Payed,
		"sonDuration":    duration,

		"artistName": strings.Join(artists, " / "),
		"AlbumName":  playingSong.Track.Album.Name,
		"albumAlias": playingSong.Track.Album.Alias,
		"albumPic":   playingSong.Track.Album.PicUrl,

		"from": playingSong.Text,
	})

}

func main() {

	go func() {
		NetEaseMusicPlaying.Watch()
	}()
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Handle(http.MethodGet, "/CC981E5D-9E4E-4DBB-95B1-FA6855E5F3B5/playing", Playing)

	log.Fatal(r.Run("127.0.0.1:9163"))
}
