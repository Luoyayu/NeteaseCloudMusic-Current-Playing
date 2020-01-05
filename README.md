NeteaseMusic.app do not provide the API to get what now playing.  

1.  access $TMPDIR/../0/com.apple.notificationcenter/db2/db => SQLite 3.x database
2.  access $HOME/Library/Containers/com.netease.163music/Data/Documents/storage/file_storage/webdata/file/history => Apple binary property list
3.  hook NeteaseMusic.app

the latest version use the method <2>
the history path refers to [obs-netease-music-now-playing](https://github.com/TinkoLiu/obs-netease-music-now-playing/blob/a42c008294e08438702168bf58de098acb0e19d7/neteasemusic.lua#L410)
 

!!! only +build for Darwin !!!
test in macOS Catalina 10.15.2 (Darwin 19.2.0)   

NeteaseMusic version: 2.3.2(build: 832)

# NeteaseMusic API
see more in https://github.com/Luoyayu/goutils/tree/master/netease

# License
MIT
