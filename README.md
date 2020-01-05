NeteaseMusic.app do not provide the API to get what now playing.  

# macOS
1.  access $TMPDIR/../0/com.apple.notificationcenter/db2/db => SQLite 3.x database
2.  access $HOME/Library/Containers/com.netease.163music/Data/Documents/storage/file_storage/webdata/file/history => Apple binary property list
3.  hook NeteaseMusic.app

# windows
1. access $TMP\..\Packages\1F8B0F94.122165AE053F_j2p0p5q0044a6\LocalState\_cloudmusic.sqlite

the latest version use the method osx:<2> and win:<1>

# NeteaseMusic API
see more in https://github.com/Luoyayu/goutils/tree/master/netease

# License
MIT
