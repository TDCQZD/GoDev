package main

import (
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	// _ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

// SyncUser2 describes a user
type SyncUser2 struct {
	Id      int64
	Name    string `xorm:"unique"`
	Age     int    `xorm:"index"`
	Title   string
	Address string
	Genre   string
	Area    string
	Date    int
}

// SyncLoginInfo2 describes a login information
type SyncLoginInfo2 struct {
	Id       int64
	IP       string `xorm:"index"`
	UserId   int64
	AddedCol int
	// timestamp should be updated by database, so only allow get from db
	TimeStamp string
	// assume
	Nonuse int    `xorm:"unique"`
	Newa   string `xorm:"index"`
}

//  sync 同步能够部分智能的根据结构体的变动检测表结构的变动，并自动同步
func sync(engine *xorm.Engine) error {
	return engine.Sync(&SyncLoginInfo2{}, &SyncUser2{})
}

/*
func sqliteEngine() (*xorm.Engine, error) {
	f := "sync.db"
	//os.Remove(f)

	return xorm.NewEngine("sqlite3", f)
}
func postgresEngine() (*xorm.Engine, error) {
	return xorm.NewEngine("postgres", "dbname=xorm_test sslmode=disable")
}
*/
//mysql数据库配置常量
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "go_db_xorm"
)

func mysqlEngine() (*xorm.Engine, error) {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	engine, err := xorm.NewEngine("mysql", path)
	if err != nil {
		log.Fatal(err)
	}
	return engine, nil
}

type engineFunc func() (*xorm.Engine, error)

func main() {
	//engines := []engineFunc{sqliteEngine, mysqlEngine, postgresEngine}
	//engines := []engineFunc{sqliteEngine}
	//engines := []engineFunc{mysqlEngine}
	engines := []engineFunc{mysqlEngine}
	for _, enginefunc := range engines {
		Orm, err := enginefunc()
		fmt.Println("--------", Orm.DriverName(), "----------")
		if err != nil {
			fmt.Println(err)
			return
		}
		Orm.ShowSQL(true)

		
		// sync 同步
		err = sync(Orm)
		if err != nil {
			fmt.Println(err)
		}

		_, err = Orm.Where("id > 0").Delete(&SyncUser2{})
		if err != nil {
			fmt.Println(err)
		}

		user := &SyncUser2{
			Name:    "testsdf",
			Age:     15,
			Title:   "newsfds",
			Address: "fasfdsafdsaf",
			Genre:   "fsafd",
			Area:    "fafdsafd",
			Date:    1000,
		}
		_, err = Orm.Insert(user)
		if err != nil {
			fmt.Println(err)
			return
		}

		isexist, err := Orm.IsTableExist("sync_user2")
		if err != nil {
			fmt.Println(err)
			return
		}
		if !isexist {
			fmt.Println("sync_user2 is not exist")
			return
		}
	}
}
