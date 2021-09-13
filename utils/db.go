package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"  //使用mysql，一定要导入这个包，但要用_隐起来
)

var (
	Db *sql.DB
	err error
)

func init() {
	Db,err = sql.Open("mysql","root:root@tcp(localhost:3306)/bookstore0612")
	if err != nil{
		//抛异常
		panic(err.Error())
	}
}
