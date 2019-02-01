package mygo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MyDB struct {
	db *sql.DB
}

func NewMyDB() *MyDB {
	return &MyDB{}
}

func (myDb *MyDB) Open() (err error) {
	dsn := myDb.parseConnConfig()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	myDb.db = db
	return
}

func (myDb *MyDB) parseConnConfig() (dsn string) {
	conf := G_config.Mysql
	dsn = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Character,
	)

	return
}

func (myDb *MyDB) Close() {
	myDb.db.Close()
}
