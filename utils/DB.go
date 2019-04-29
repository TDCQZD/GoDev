package utils

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	ip   = "127.0.0.1"
	port = "3306"
)

func DBMySql(userName, password, dbName string) (*sql.DB, error) {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	DB, err := sql.Open("mysql", path)
	//验证连接
	if errConn := DB.Ping(); errConn != nil {
		fmt.Println("open database fail")
		return nil, err
	}
	fmt.Println("connnect success")
	defer DB.Close()
	return DB, nil
}
