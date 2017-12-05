package Database

import(
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"strconv"
)

//func GetGameData(uid int) {
//
//	var pId = strconv.Itoa(uid);
//
//	db, err := sql.Open("sqlite3",
//		"./PathFynder.db")
//
//	rows, err := db.Query("SELECT * FROM ABSOLUTE ")
//
//}

func GetHighScore(uid int) {

	var userId = strconv.Itoa(uid);

	db, err := sql.Open("sqlite3",
		"./PathFynder.db")

	rows, err := db.Query("SELECT playerid, highscore, date FROM Highscores " +
		"WHERE playerid =" + userId + ";")
	if err != nil {
		log.Fatal(err)
	}

	var id int
	var score int
	var date string

	for rows.Next() {
		err = rows.Scan(&id, &score, &date)
		fmt.Println(id)
		fmt.Println(score)
		fmt.Println(date)
	}
}