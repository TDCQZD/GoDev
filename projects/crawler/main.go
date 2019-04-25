package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

// 单任务爬虫珍爱网-城市列表页
func main() {

	url := "http://www.zhenai.com/zhenghun"
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:status code:", resp.StatusCode)
		return
	}
	// 通过http协议获取整个页面的数据
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	// writeFile("city.txt", all)
	// fmt.Printf("%s\n", all) // 打印城市信息
	// PrintCityList(all) // 打印城市列表信息
	PrintCityListByJson(all) // 打印城市列表JSON信息
}

//打印城市信息
func PrintCityList(contents []byte) {

	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`)
	all := re.FindAllSubmatch(contents, -1)

	for _, c := range all {
		fmt.Printf("City\t%s\t,URL\t%s", c[2], c[1])
		fmt.Println()
	}
	fmt.Println(len(all))
}

//打印城市信息-通过json获取
func PrintCityListByJson(contents []byte) {

	re := regexp.MustCompile(`<script>window.__INITIAL_STATE__=(.+);\(function`)

	all := re.FindSubmatch(contents)

	fmt.Printf("%s", all[1])

}

func writeFile(filePath string, data []byte) {

	//1、打开文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 066) // For read access.
	if err != nil {
		log.Println(err)
	}
	//2、关闭文件
	defer file.Close()

	// 3、写入数据

	err = ioutil.WriteFile(filePath, data, 066)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("数据写入成功！")
}
