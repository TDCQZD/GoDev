package main

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//数据库配置常量
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "go_db_gorm"
)

// Use pointer value
type User struct {
	gorm.Model
	Name string
	Age  int `gorm:"default:18"`
}

func main() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	db, err := gorm.Open("mysql", path)
	if err != nil {
		panic("连接数据库失败")
	} else {
		fmt.Println("连接数据库成功！")
	}
	defer db.Close()

	isflag := db.HasTable("users")
	if !isflag {
		db.CreateTable(&User{})
	}

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// 自动迁移模式
	db.AutoMigrate(&User{})
	/*
		user := User{Name: "Jinzhu", Age: 18}

		db.NewRecord(user) // => returns `true` as primary key is blank

		db.Create(&user)

		db.NewRecord(user) // => return `false` after `user` created
	*/

	var user User
	db.First(&user) // 没有数据？？
	fmt.Println(user.Age)

	row := db.Table("users").Where("name = ?", "jinzhu").Select("name, age").Row() // (*sql.Row)
	row.Scan(&user.Name, &user.Age)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
}
