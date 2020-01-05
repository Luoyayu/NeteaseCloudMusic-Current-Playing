// +build windows

package NeteaseMusicPlaying

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"strings"
)

var db *sql.DB
var HistoryFilePath string
var Playing = &HistoryStruct{}

func init() {
	HistoryFilePath = strings.Join([]string{
		os.Getenv("TMP"),
		"..",
		"Packages",
		"1F8B0F94.122165AE053F_j2p0p5q0044a6",
		"LocalState",
		"_cloudmusic.sqlite",
	}, string(os.PathSeparator))
	log.Println(HistoryFilePath)

	f, err := os.Open(HistoryFilePath)
	if err == nil || os.IsExist(err) {
		log.Println("splite 3 db file is exist")
		defer f.Close()
	} else {
		log.Fatalln("splite 3 db file is Not Exist!")
	}

	db, err = sql.Open("sqlite3", HistoryFilePath)
	if err != nil {
		panic(err)
	}

}

func Update() {
	sqlStmt := `
		SELECT resourcedata FROM playhistory
		ORDER BY updatetime DESC LIMIT 1;`

	var data = ""
	row := db.QueryRow(sqlStmt)
	_ = row.Scan(&data)
	_ = json.Unmarshal([]byte(data), &Playing)

	fmt.Printf("%+v\n", Playing)
}
