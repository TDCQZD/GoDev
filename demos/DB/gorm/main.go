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

type Product struct {
	gorm.Model
	Code  string
	Price uint
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

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 检查模型`Product`表是否存在
	flag := db.HasTable(&Product{})
	if !flag {
		fmt.Println("模型`Product`表不存在")
	} else {
		fmt.Println("模型`Product`表存在")
	}
	// 检查表`product`是否存在
	isflag := db.HasTable("products")
	if !isflag {
		fmt.Println("表products不存在")
		db.CreateTable(&Product{})
		// db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&User{}) 添加引擎
	} else {
		fmt.Println("表products存在")
		// db.DropTable("products") // 删除表
	}

	// 创建
	// db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.Last(&product)
	fmt.Println("product:", product)
	db.First(&product, 1) // 查询id为1的product
	fmt.Println("Code:", product.Code)
	db.First(&product, "code = ?", "L1213") // 查询code为l1212的product
	fmt.Println("Price:", product.Price)

	// 更新 - 更新product的price为2000
	// db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	// db.Delete(&product)

}
