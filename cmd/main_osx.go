//+build darwin

package main

import NeteaseMusicPlaying "163musicPlaying"

func main() {
	NeteaseMusicPlaying.Watch(NeteaseMusicPlaying.HistoryFilePath)
}
