package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"
)

// User describes a user
type User struct {
	Id   int64
	Name string
}

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
		log.Fatalf("CreateEngine |NewEngine err=%v", err)
		panic(err.Error())
	}
	return engine
}
func main() {
	Orm := CreateEngine()
	err := Orm.Ping()
	if err != nil {
		log.Fatalf("Ping Engine err=%v", err)
		panic(err.Error())
	} else {
		fmt.Println("数据库连接成功！")
	}
	Orm.ShowSQL(true)

	// 缓存设置
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Orm.SetDefaultCacher(cacher)

	err = Orm.CreateTables(&User{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = Orm.Insert(&User{Name: "xlw"})
	if err != nil {
		fmt.Println(err)
		return
	}

	var users []User
	err = Orm.Find(&users)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("users:", users)

	var users2 []User
	err = Orm.Find(&users2)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("users2:", users2)

	var users3 []User
	err = Orm.Find(&users3)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("users3:", users3)

	user4 := new(User)
	has, err := Orm.ID(1).Get(user4)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("user4:", has, user4)

	user4.Name = "xiaolunwen"
	_, err = Orm.ID(1).Update(user4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user4:", user4)

	user5 := new(User)
	has, err = Orm.ID(1).Get(user5)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("user5:", has, user5)

	_, err = Orm.ID(1).Delete(new(User))
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		user6 := new(User)
		has, err = Orm.ID(1).Get(user6)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("user6:", has, user6)
	}
}
