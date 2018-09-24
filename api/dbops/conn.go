package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 申明这个dbops包下面的使用到的变量
var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	dbConn.Ping()
}
