NeteaseMusic.app do not provide the API to get what now playing.  

below are some method to access:  

# macOS
1.  `$TMPDIR/../0/com.apple.notificationcenter/db2/db` => SQLite 3.x database  
2.  `$HOME/Library/Containers/com.netease.163music/Data/Documents/storage/file_storage/webdata/file/history` => Apple binary property list  
3.  hook NeteaseMusic.app  

# windows
1. `%TMP%\..\Packages\1F8B0F94.122165AE053F_j2p0p5q0044a6\LocalState\_cloudmusic.sqlite`   

the latest version use the method osx:<2> and win:<1>  

!!! when first boot NeteaseMusic, you need change song to active !!!   

# build
windows version need enable cgo and corresponding toolchain   

# Release
release exe provides api on `127.0.0.1:9163/playing`  

# NeteaseMusic API
see more in https://github.com/Luoyayu/goutils/tree/master/netease  

# License
MIT   
