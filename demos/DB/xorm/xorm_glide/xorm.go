package main

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//mysql数据库配置常量
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "go_db_xorm"
)

func ConnectDB() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	engine, err := xorm.NewEngine("mysql", path)
	if err != nil {
		ErrFormat("ConnectDb |NewEngine err=%v", err)
		panic(err.Error())
	}
}

// func createTable(DB *sql.DB) error {}
// func InsertData(DB *sql.DB, user bean.User) bool {}
// func SelectAllData(DB *sql.DB) []bean.User {}
// func SelectDataById(DB *sql.DB, id int) bean.User {}
// func UpdateData(DB *sql.DB, user bean.User) bool {
// func DeleteData(DB *sql.DB, user bean.User) bool {
func main() {

}
func ErrFormat(format string, args ...interface{}) {
	logs.Error(fmt.Sprintf(format, args...))
}

func ErrLog(comment string, uid int64, err error) string {
	return fmt.Sprintf("%v uid = %v | err = %v", comment, uid, err)
}
