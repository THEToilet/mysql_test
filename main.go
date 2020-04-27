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


	fmt.Printf("%#v\n",db)
	
	userEx := Users{}

	userEx.Name = "Toilet"
	userEx.AcceptedCount = 123
	userEx.AcceptedCountRank = 12
	userEx.RatedPointSum = 34
	userEx.RatedPointSumRank = 12
	userEx.CreatedTime = getDate()
	userEx.UpdatedTime = getDate()

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
	print(users)
	//表示
	for _, user := range users {
		fmt.Printf("%#v\n",user)
		fmt.Printf("%#v\n",user.Id)
		fmt.Printf("%#v\n",user.Name)
		fmt.Printf("%#v\n",user.AcceptedCount)
		fmt.Printf("%#v\n",user.AcceptedCountRank)
		fmt.Printf("%#v\n",user.RatedPointSum)
		fmt.Printf("%#v\n",user.RatedPointSumRank)
		fmt.Printf("%#v\n",user.CreatedTime)
		fmt.Printf("%#v\n",user.UpdatedTime)
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
	Id                   int
	Name                 string    `gorm:"column:name"`
	AcceptedCount       int       `gorm:"column:accepted_count"`
	AcceptedCountRank  int       `gorm:"column:accepted_count_rank"`
	RatedPointSum      int       `gorm:"column:rated_point_sum"`
	RatedPointSumRank int       `gorm:"column:rated_point_sum_rank"`
	CreatedTime         string    `gorm:"column:created_time" sql:"not null;type:date"`
	UpdatedTime         string    `gorm:"column:updated_time" sql:"not null;type:date"`
}
