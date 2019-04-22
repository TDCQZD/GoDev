package main

import (
	"fmt"
	"strings"
	"xorm_govendor/utils"

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

func CreateEngine() *xorm.Engine {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	engine, err := xorm.NewEngine("mysql", path)
	if err != nil {
		utils.ErrFormat("ConnectDb |NewEngine err=%v", err)
		panic(err.Error())
	}
	return engine
}

func main() {
	engine := CreateEngine()
	err := engine.Ping()
	if err != nil {
		utils.ErrFormat("Ping Engine err=%v", err)
		panic(err.Error())
	} else {
		fmt.Println("数据库的连接成功！")
	}
}
