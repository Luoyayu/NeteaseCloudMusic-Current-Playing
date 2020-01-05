// +build windows

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

	duration, _ := time.ParseDuration(fmt.Sprint(playingSong.Track.Dt/1000) + "s")

	c.JSONP(http.StatusOK, map[string]interface{}{
		"songId":         playingSong.Track.Id,
		"songName":       playingSong.Track.Name,
		"songPopularity": playingSong.Track.Pop,
		"songIsPayed":    playingSong.Track.Privilege.Payed,
		"sonDuration":    duration,

		"artistName": strings.Join(artists, " / "),
		"AlbumName":  playingSong.Track.Album.Name,
		"albumPic":   playingSong.Track.Album.Pic,

		"from": playingSong.Track.Source.SourceName,
	})

}

func main() {

	go func() {
		NetEaseMusicPlaying.Watch(NetEaseMusicPlaying.HistoryFilePath)
	}()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Handle(http.MethodGet, "/playing", Playing)

	log.Fatal(r.Run("127.0.0.1:9163"))
}
