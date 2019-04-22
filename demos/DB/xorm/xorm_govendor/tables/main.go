package main

import (
	"fmt"
	_ "os"
	"strings"
	"xorm_govendor/bean"
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

	// if len(os.Args) < 2 {
	// 	fmt.Println("need db path")
	// 	return
	// }

	orm := CreateEngine()
	err := orm.Ping()
	if err != nil {
		utils.ErrFormat("Ping Engine err=%v", err)
		panic(err.Error())
	} else {
		fmt.Println("数据库连接成功！")
	}
	defer orm.Close()
	orm.ShowSQL(true)

	// IsTableExist() 判断表是否存在
	isExist, err := orm.IsTableExist(&bean.User{})
	if err != nil {
		utils.ErrFormat("IsTableExist err=%v", err)
		panic(err.Error())
	}

	if isExist {
		fmt.Println("table已存在，删除table")
		// DropTables() 删除表
		err = orm.DropTables(&bean.User{})
		if err != nil {
			utils.ErrFormat("DropTables err=%v", err)
			panic(err.Error())
		}
	} else {
		fmt.Println("table不存在")
	}

	// CreateTables 创建表
	orm.CreateTables(&bean.User{})
	if err != nil {
		utils.ErrFormat("CreateTable err=%v", err)
		panic(err.Error())
	}
	// IsTableEmpty() 判断表是否为空
	isTableEmpty, err := orm.IsTableEmpty(&bean.User{})
	if err != nil {
		utils.ErrFormat("IsTableEmpty err=%v", err)
		panic(err.Error())
	}
	if isTableEmpty {
		fmt.Println("table为空")
	}else{
		fmt.Println("table非空")
	}

	// DBMetas 获取到数据库中所有的表，字段，索引的信息。
	tables, err := orm.DBMetas()
	if err != nil {
		utils.ErrFormat("DBMetas err=%v", err)
		panic(err.Error())
	}

	for _, table := range tables {
		fmt.Println(table.Name)
	}

}
