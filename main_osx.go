// +build darwin

package NeteaseMusicPlaying

import (
	"encoding/json"
	"fmt"
	"howett.net/plist"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type historyStruct struct {
	Objects []string `plist:"$objects"`
}

var HistoryFilePath string

func init() {
	HistoryFilePath = os.Getenv("HOME") +
		"/Library/Containers/com.netease.163music/Data/Documents/storage/file_storage/webdata/file/history"
}


var Playing = &Song{}

func Update() {
	f, err := os.Open(HistoryFilePath)
	historyData, _ := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var history historyStruct

	_, err = plist.Unmarshal(historyData, &history)
	if err != nil {
		log.Fatalln(err)
	}

	var tracks []Song
	_ = json.Unmarshal([]byte(history.Objects[1]), &tracks)

	var artists []string
	for _, a := range tracks[0].Track.Artists {
		artists = append(artists, a.Name)
	}

	PlayingMutex.Lock()
	Playing = &tracks[0]
	PlayingMutex.Unlock()

	duration, _ := time.ParseDuration(fmt.Sprint(Playing.Track.Duration/1000) + "s")

	fmt.Println("song id:", Playing.Track.Id)
	fmt.Println("song name:", Playing.Track.Name)
	fmt.Println("song alias:", Playing.Track.Alias)
	fmt.Println("song popularity:", Playing.Track.Popularity)
	fmt.Println("song isPayed:", Playing.Track.Privilege.Payed)
	fmt.Println("song duration:", duration)

	fmt.Println("artists name:", strings.Join(artists, " / "))

	fmt.Println("album name:", Playing.Track.Album.Name)
	fmt.Println("album alias", Playing.Track.Album.Alias)
	fmt.Println("album pic:", Playing.Track.Album.PicUrl)

	fmt.Println("from:", Playing.Text)
	fmt.Println("////////////////////////////////////////////////////////////////////////////")
}
