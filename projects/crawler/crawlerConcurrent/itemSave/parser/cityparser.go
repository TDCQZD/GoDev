package parser

import (
	"GoDev/projects/crawler/crawlerConcurrent/itemSave/engine"
	"regexp"
)

var (
	// 城市
	cityRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	// 下一页
	cityUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

// ParseCity 城市解析器 获取用户的url和用户名
func ParseCity(contents []byte) engine.ParseResult {

	all := cityRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all {
		result.Items = append(result.Items, "User:"+string(c[2])) //用户名字

		name := string(c[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			// ParserFunc: engine.NilParser, //爬取城市列表页时，为了让程序不报错能够正常执行，设置ParserFunc为engine.NilParser函数
			// ParserFunc: ParseUser, //城市解析器
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)
			}, // 用户解析器
		})

	}

	//爬取下一页
	next := cityUrlRe.FindAllSubmatch(contents, -1)
	for _, c := range next {
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(c[1]),
			ParserFunc: ParseCity,
		})
	}
	return result
}
