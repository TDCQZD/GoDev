package parser

import (
	"GoDev/projects/crawler/crawlerConcurrent/concurrent/engine"
	"log"
	"regexp"

	"github.com/bitly/go-simplejson"
)

// 城市列表正则
// const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`
const cityListReg = `<script>window.__INITIAL_STATE__=(.+);\(function`

// ParseCityList 城市列表解析器 获取城市url和城市名字
func ParseCityList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityListReg)
	// datas := reg.FindAllSubmatch(contents, -1)
	json := reg.FindSubmatch(contents)
	datas := parseJsonCityList(json[1])

	result := engine.ParseResult{}

	for _, c := range datas {
		// result.Items = append(result.Items, string(c[2])) //城市名字
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			// ParserFunc: engine.NilParser, // 空解析器
			ParserFunc: ParseCity, // 城市解析器

		})

	}

	return result
}

//解析json数据
func parseJsonCityList(json []byte) [][]string {
	res, err := simplejson.NewJson(json)
	if err != nil {
		log.Println("解析json失败。。")
	}
	infos, _ := res.Get("cityListData").Get("cityData").Array()
	//infos是一个切片，里面的类型是interface{}

	var dataList [][]string
	//所以我们遍历这个切片，里面使用断言来判断类型
	for _, v := range infos {
		if each_map, ok := v.(map[string]interface{}); ok {
			//fmt.Println(each_map)
			map2 := each_map["cityList"]
			for _, v2 := range map2.([]interface{}) {
				if data, ok := v2.(map[string]interface{}); ok {
					var datas []string
					cityName := data["linkContent"].(string)
					cityUrl := data["linkURL"].(string)
					datas = append(datas, cityName, cityUrl)
					dataList = append(dataList, datas)
				}
			}
		}
	}
	return dataList

}
