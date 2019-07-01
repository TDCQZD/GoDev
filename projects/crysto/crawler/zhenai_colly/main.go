package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
)

//珍爱网用户对象模型
type UserInfo struct {
	ID         string // ID
	NickName   string //昵称
	Marriage   string //婚况
	Age        string //年龄
	Gender     string //性别
	Height     string //身高
	Weight     string //体重
	Income     string //收入
	Education  string //教育
	Occupation string //职业
	Hukou      string //籍贯户口
	Xingzuo    string //星座
	House      string //房子
	Car        string //车
}

func main() {
	c := colly.NewCollector(
	// colly.URLFilters(
	// 	regexp.MustCompile(`http://www.zhenai.com/zhenghun/+`),
	// 	regexp.MustCompile(`http://www.zhenai.com/zhenghun/[0-9a-z]+`),
	// ),
	)
	cityCollector := colly.NewCollector()
	userCollector := colly.NewCollector()

	// 在请求之前调用
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	// 如果在请求期间发生错误，则调用
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	// 收到回复后调用
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML(`.g-container a[href]`, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// e.Request.Visit(link)
		reg := regexp.MustCompile(`http://www.zhenai.com/zhenghun/[0-9a-z]+`)
		if reg.MatchString(link) {
			cityCollector.Visit(link)
			fmt.Println("link", link)
		}

	})
	cityCollector.OnHTML(`.g-list a[href]`, func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// e.Request.Visit(link)
		reg := regexp.MustCompile(`http://album.zhenai.com/u/[0-9]+`)
		if reg.MatchString(link) {
			fmt.Println("cityCollector", link)
			userCollector.Visit(link)
		}
	})

	userCollector.OnHTML("body", func(e *colly.HTMLElement) {
		user := UserInfo{}
		nickName := e.DOM.Find(".m-userInfo .top .right .info .name .nickName").Text()

		id := e.DOM.Find(".id").Text()
		rs := []rune(id)
		id = string(rs[2:])
		// fmt.Printf("id=%s---nickName=%s \n", id, nickName)

		marrige := e.DOM.Find(".m-userInfoDetail .m-content-box .purple-btns > div:first-child").Text()
		high := e.DOM.Find(".m-userInfoDetail .m-content-box .purple-btns > div:nth-child(4)").Text()

		fmt.Printf("id=%s---nickName=%s---marrige=%s---high=%s \n", id, nickName, marrige, high)

		user.ID = id
		user.NickName = nickName
		user.Marriage = marrige
		user.Height = high
		insertDB(user)
	})

	c.Visit("http://www.zhenai.com/zhenghun/")

}

func insertDB(user UserInfo) {
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
		db.CreateTable(&UserInfo{})
	}

	// 启用Logger，显示详细日志
	db.LogMode(true)

	// 自动迁移模式
	db.AutoMigrate(&UserInfo{})

	db.Create(&user)
}

//数据库配置常量
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "crawler"
)
