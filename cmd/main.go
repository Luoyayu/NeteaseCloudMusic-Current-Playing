// +build darwin

package main

import (
	NetEaseMusicPlaying "163musicPlaying"
	"fmt"
	"time"
)

func main() {
	for {
		NetEaseMusicPlaying.UpdateNotification()
		if NetEaseMusicPlaying.CurrentPlayingSong.LastingPlayTime > 40*60 {
			NetEaseMusicPlaying.CurrentPlayingSong.IsPaused = true
			fmt.Println("歌曲已暂停播放, 或单曲循环 40 min 以上")
		} else if NetEaseMusicPlaying.CurrentPlayingSong.SongName != "" {
			fmt.Println("歌曲:", NetEaseMusicPlaying.CurrentPlayingSong.SongName)
			fmt.Println("专辑:", NetEaseMusicPlaying.CurrentPlayingSong.AlbumName)
			fmt.Println("歌手:", NetEaseMusicPlaying.CurrentPlayingSong.ArtistName)
			fmt.Println("持续播放: ", NetEaseMusicPlaying.CurrentPlayingSong.LastingPlayTime)
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		} else {
			fmt.Println("未使用网易云播放")
		}

		time.Sleep(time.Second * 2)
	}
}
