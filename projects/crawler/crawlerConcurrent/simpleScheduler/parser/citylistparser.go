package parser

import (
	"GoDev/projects/crawler/crawlerConcurrent/simpleScheduler/engine"
	"regexp"
)

// 城市列表正则
const cityListReg = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

// ParseCityList 城市列表解析器 获取城市url和城市名字
func ParseCityList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityListReg)
	datas := reg.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	i := 0
	for _, c := range datas {
		result.Items = append(result.Items, string(c[2])) //城市名字
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			// ParserFunc: engine.NilParser, // 空解析器
			ParserFunc: ParseCity, // 城市解析器

		})
		// 现在仅爬取10个城市
		i++
		if i == 1 {
			break
		}
	}
	return result
}
