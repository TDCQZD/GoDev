package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/godb")
	if err != nil {
		panic(err.Error())
	}
}