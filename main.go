package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

func main() {
	// db接続
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	userEx := Users{}

	userEx.name = "Toilet"
	userEx.accepted_count = 123
	userEx.accepted_count_rank = 12
	userEx.rated_point_sum = 34
	userEx.rated_point_sum_rank = 12
	userEx.created_time = getDate()
	userEx.updated_time = getDate()

	// INSERTを実行
  db.Create(&userEx)	
/*	
	error := db.Create(&Users{
								name: "Toilet",
								accepted_count: 123,
								accepted_count_rank: 12,
								rated_point_sum: 34,
								rated_point_sum_rank: 12,
								created_time: getDate(),
								updated_time: getDate(),
	}).Error
	
  if error != nil {
        fmt.Println(error)
    } else {
        fmt.Println("データ追加成功")
   }
*/

	//データを格納する変数を定義
	users := []Users{}

	//全取得
	db.Find(&users)

	//表示
	for _, user := range users {
		fmt.Println(user.id)
		fmt.Println(user.name)
		fmt.Println(user.accepted_count)
		fmt.Println(user.accepted_count_rank)
		fmt.Println(user.rated_point_sum)
		fmt.Println(user.rated_point_sum_rank)
		fmt.Println(user.created_time)
		fmt.Println(user.updated_time)
	}
}


func getDate() string {
    const layout = "2006-01-02 15:04:05"
    now := time.Now()
    return now.Format(layout)
}

// SQLConnect DB接続
func sqlConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := "root"
	PASS := "pass"
	PROTOCOL := "tcp(127.0.0.1:3306)"
	DBNAME := "user_db"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	return gorm.Open(DBMS, CONNECT)
}

// Users ユーザー情報のテーブル情報
type Users struct {
	id                   int
	name                 string    `json:"name"`
	accepted_count       int       `json:"accepted_count"`
	accepted_count_rank  int       `json:"accepted_count_rank"`
	rated_point_sum      int       `json:"rated_point_sum"`
	rated_point_sum_rank int       `json:"rated_point_sum_rank"`
	created_time         string    `json:"created_time" sql:"not null;type:date"`
	updated_time         string    `json:"updated_time" sql:"not null;type:date"`
}
