package main

import (
	//"database/sql"
	//"fmt"
	//"github.com/alphazero/Go-Redis"
	"github.com/astaxie/beego"
	// _ "github.com/go-sql-driver/mysql"
	// "log"
	//"strconv"
)

func main() {
	// var a int = 65
	// b := strconv.Itoa(a)
	// a, _ = strconv.Atoi(b)
	// fmt.Println(a)
	beego.Run()
}

// func redisTest() {
// 	spec := redis.DefaultSpec().Db(0).Password("")
// 	client, err := redis.NewSynchClientWithSpec(spec)
// }

// func mysqlTest() {
// 	db, err := sql.Open("mysql", "proj_appserver:appserver21696@tcp(192.168.1.10:3306)/hsmjia")
// 	if err != nil {
// 		log.Fatal("Open database error: %s\n", err)
// 	}
// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	rows, err := db.Query("select factoryid,name from factories where factoryid = ?", 2)
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer rows.Close()
// 	var factoryid int
// 	var name string
// 	for rows.Next() {
// 		err := rows.Scan(&factoryid, &name)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		log.Println(factoryid, name)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
