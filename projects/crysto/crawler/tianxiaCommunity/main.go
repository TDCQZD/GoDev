package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/axgle/mahonia"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
)

//数据库配置
const (
	userName = "root"
	password = "root"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "crawler"
)

type XQinfo struct {
	name           string
	addr           string
	area           string
	postCode       string
	propertyRight  string
	propertyType   string
	buildTime      string
	developer      string
	buildType      string
	buildArea      string
	buildStruct    string
	floorSpace     string
	manageCompany  string
	greenRatio     string
	plotRatio      string
	propretyFee    string
	AdditionalInfo string
	waterSupply    string
	heatSupply     string
	elecSupply     string
	gas            string
	security       string
	environment    string
	parkingSpace   string
	OtherInfo      string
}

var flagCh = make(chan int)
var count = 1
var fileName = "./flag.txt"
var file *os.File
var err error

func main() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	max := 2
	file = openFile(fileName)

	for i := 0; i < max; i++ {
		for j := 1; j <= 10; j++ {
			DB, _ := sql.Open("mysql", path)
			//验证连接
			if errConn := DB.Ping(); errConn != nil {
				fmt.Println("open database fail")
				return
			}
			fmt.Println("connnect success")
			defer DB.Close()

			link := "http://tj.esf.fang.com/housing/__0_0_0_0_" + strconv.Itoa(i*10+j) + "_0_0_0/"
			go work(link, DB, i*10+j)
		}
	}

	for {
		<-flagCh
		if count < max*10 {
			fmt.Println("<- receive the " + strconv.Itoa(count) + " thread ending flag")
			count++
		} else {
			break
		}

		fmt.Println("All " + strconv.Itoa(count) + " has done")
	}

	defer file.Close()
}

func openFile(fileName string) *os.File {
	if checkFileIsExist(fileName) {
		//如果文件存在
		file, err = os.OpenFile(fileName, os.O_APPEND, 0666)
	} else {
		//创建文件
		file, err = os.Create(fileName)
	}
	check(err)
	return file
}

func writeFile(file *os.File, content string) {
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.WriteString("\r\n")
	writer.Flush()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func work(url string, DB *sql.DB, page int) {
	c := colly.NewCollector()
	detailLink := c.Clone()
	detailController := c.Clone()
	infos := make([]XQinfo, 0)

	c.OnHTML(".plotListwrap > dt > a", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fmt.Printf("link : %s \t", link)
		fmt.Println()

		detailLink.Visit(link)
	})

	detailLink.OnHTML("#kesfxqxq_A01_03_01", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		//content := e.ChildText("a")

		//fmt.Printf("detial link : %s \t", link)
		//fmt.Printf("detial content : %s \t", coverGBKToUTF8(content))
		//fmt.Println()

		detailController.Visit(link)
	})

	detailController.OnHTML("body", func(e *colly.HTMLElement) {

		info := XQinfo{}

		// 小区名称
		name := e.DOM.Find(".ceninfo_sq > h1 > a").Text()
		info.name = coverString(name)

		e.DOM.Find(".inforwrap").Each(func(i int, selection *goquery.Selection) {

			// 模块名称
			modelName := coverString(selection.Prev().Find("h3").Text())
			//fmt.Println("h3  -> ", modelName)

			switch modelName {
			case "基本信息":
				dealInfo(selection, &info)
			case "配套设施":
				dealInfo(selection, &info)
			case "周边信息":
				selection.Find("dl dt").Each(func(_ int, otherSelect *goquery.Selection) {
					tab := coverString(otherSelect.Text())
					del := strings.Index(tab, "本段合作")
					if del == -1 {
						info.OtherInfo = info.OtherInfo + tab + "|"
					}
				})
			}
		})

		infos = append(infos, info)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnScraped(func(response *colly.Response) {

		for _, info := range infos {
			insertDB(DB, info)
		}

		fmt.Println("the " + strconv.Itoa(page) + " thread sending end flag ->")

		flagCh <- 1
	})

	c.Visit(url)
}

// 处理小区基础信息
func dealInfo(selection *goquery.Selection, info *XQinfo) {
	selection.Find("dl dd").Each(func(_ int, selectionbase *goquery.Selection) {
		setXQinfo(selectionbase, info)
	})

	selection.Find("dl dt").Each(func(_ int, selectionbase *goquery.Selection) {
		setXQinfo(selectionbase, info)
	})
}

func setXQinfo(selectionbase *goquery.Selection, info *XQinfo) {

	orgKey := coverString(selectionbase.Find("strong").Text())
	index := strings.Index(orgKey, "：")

	var key string
	if index > 0 {
		key = orgKey[:index]
	} else {
		key = orgKey
	}

	var value string
	var fullValue string
	value, ok := selectionbase.Attr("title")
	if ok {
		value = coverString(value)
	} else {
		fullValue = coverString(selectionbase.Text())
		value = fullValue[strings.Index(fullValue, "：")+3:]
	}

	switch key {
	case "小区地址":
		info.addr = value
	case "所属区域":
		info.area = value
	case "邮编":
		info.postCode = value
	case "产权描述":
		info.propertyRight = value
	case "物业类别":
		info.propertyType = value
	case "建筑年代":
		info.buildTime = value
	case "开 发 商":
		info.developer = value
	case "建筑结构":
		info.buildStruct = value
	case "建筑类型":
		info.buildType = value
	case "建筑面积":
		info.buildArea = value
	case "占地面积":
		info.floorSpace = value
	case "物业公司":
		info.manageCompany = value
	case "绿 化 率":
		info.greenRatio = value
	case "容 积 率":
		info.plotRatio = value
	case "物 业 费":
		info.propretyFee = value
	case "附加信息":
		info.AdditionalInfo = value
	case "供水":
		info.waterSupply = value
	case "供暖":
		info.heatSupply = value
	case "供电":
		info.elecSupply = value
	case "燃气":
		info.gas = value
	case "安全管理":
		info.security = value
	case "卫生服务":
		info.environment = value
	case "停 车 位":
		info.parkingSpace = value
	}
}

//src为要转换的字符串
func coverGBKToUTF8(src string) string {
	return mahonia.NewDecoder("gbk").ConvertString(src)
}

func replaceNullHtml(src string) string {
	return strings.Replace(src, "聽", "", -1)
}

func coverString(src string) string {
	return replaceNullHtml(coverGBKToUTF8(src))
}

func insertDB(DB *sql.DB, info XQinfo) {
	t := reflect.TypeOf(info)
	v := reflect.ValueOf(info)

	sql1 := "insert into rx_xiaoqu("
	sql2 := ") values ("
	sql3 := ")"

	for i := 0; i < t.NumField(); i++ {

		sql1 = sql1 + t.Field(i).Name
		sql2 = sql2 + "'" + v.Field(i).String() + "'"

		if i != t.NumField()-1 {
			sql1 = sql1 + ", "
			sql2 = sql2 + ", "
		}

		//fmt.Printf("key -> %s, value -> %s", t.Field(i).Name, v.Field(i))
		//fmt.Println()
	}
	//fmt.Println(sql1, sql2)

	stmt, err := DB.Prepare(sql1 + sql2 + sql3)
	if err != nil {
		fmt.Println(sql1 + sql2)
		fmt.Println(err)
	}
	res, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.LastInsertId())
}
