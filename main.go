// +build darwin

package NeteaseMusicPlaying

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"howett.net/plist"
	"log"
	"os"
	"strings"
	"time"
)

var TMPDIR string
var db1DB2 string
var db *sql.DB
var appid string
var CurrentPlayingSong = NotificationPlist{}

type NotificationPlist struct {
	App  string                 `plist:"app"`
	Date float64                `plist:"date"`
	Req  map[string]interface{} `plist:"req"`

	ArtistName string
	AlbumName  string
	SongName   string

	StartPlayTime   int64
	LastingPlayTime float64
	IsPaused        bool
}

func init() {
	log.Println("网易云音乐 -> 偏好设置 -> √ 启用系统歌曲播放通知栏")

	TMPDIR = os.Getenv("TMPDIR")
	db1DB2 = TMPDIR + "../0/com.apple.notificationcenter/db2/db"
	_, err := os.Open(db1DB2)
	if err == nil || os.IsExist(err) {
		log.Println("splite 3 db file is exist")
	} else {
		log.Fatalln("splite 3 db file is Not Exist!")
	}

	db, err = sql.Open("sqlite3", db1DB2)
	if err != nil {
		log.Fatalln(err)
	}

	sqlStmt := `
		SELECT app_id FROM app 
		WHERE app.identifier is 'com.netease.163music';`

	row := db.QueryRow(sqlStmt)
	_ = row.Scan(&appid)
	log.Println("com.netease.163music identifier:", appid)

}

func UpdateNotification() {
	sqlStmt := `
		SELECT data FROM record 
		WHERE app_id is ` + appid + `;`

	row := db.QueryRow(sqlStmt)
	var data []byte
	_ = row.Scan(&data)
	if len(data) == 0 {
		CurrentPlayingSong = NotificationPlist{}
	}
	//log.Println("read hex data from db2/db:", len(data))

	defer func() {
		if r := recover(); r != nil {
			log.Println("Error:", r.(error))
		}
	}()
	_, err := plist.Unmarshal(data, &CurrentPlayingSong)
	if err != nil {
		log.Fatalln(err)
	}

	//fmt.Println("CurrentPlayingSong: ", CurrentPlayingSong)
	body := CurrentPlayingSong.Req["body"].(string)

	bodyS := strings.Split(body, " - ")
	artistName := bodyS[0]
	albumName := bodyS[1]
	if len(bodyS) > 2 {
		albumName = strings.TrimLeft(body, artistName+" - ")
	}

	songName := CurrentPlayingSong.Req["titl"].(string)
	CurrentPlayingSong.StartPlayTime = time.Unix(int64(CurrentPlayingSong.Date)+978307205, 0).Unix()

	if songName == CurrentPlayingSong.SongName {
		CurrentPlayingSong.LastingPlayTime = time.Now().Sub(time.Unix(CurrentPlayingSong.StartPlayTime, 0)).Seconds()
	} else {
		CurrentPlayingSong = NotificationPlist{}
	}

	CurrentPlayingSong.StartPlayTime = time.Unix(int64(CurrentPlayingSong.Date)+978307205, 0).Unix()
	CurrentPlayingSong.ArtistName = strings.TrimSpace(artistName)
	CurrentPlayingSong.AlbumName = strings.TrimSpace(albumName)
	CurrentPlayingSong.SongName = strings.TrimSpace(songName)
}
