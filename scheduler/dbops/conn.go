package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// 此数据库连接定义初始化文件和之前的一样
var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/video_server")
	if err != nil {
		panic(err.Error())
	}
}
