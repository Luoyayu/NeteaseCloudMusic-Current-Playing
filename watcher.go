package NeteaseMusicPlaying

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"sync"
)

var PlayingMutex sync.Mutex

func Watch(filePath string) {
	//Update()
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	defer watch.Close()

	done := make(chan struct{})

	go func() {
		for {
			select {
			case event, ok := <-watch.Events:
				if !ok {
					return
				}
				log.Println("event:", event.Op)
				if event.Op == fsnotify.Create {
					log.Println("watcher: log file create")
					Update()
				} else if event.Op == fsnotify.Write {
					log.Println("watcher: log file write")
					//time.Sleep(time.Second * 1)
					Update()
				}

			case err, ok := <-watch.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	if err = watch.Add(filePath); err != nil {
		log.Fatalln(err)
	}
	<-done
}
