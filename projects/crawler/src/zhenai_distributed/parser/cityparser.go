package parser

import (
	"GoDev/projects/crawler/src/zhenai_Single/engine"
	"fmt"
	"regexp"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//解析信息 获取用户的url和用户名
func ParseCity(contents []byte) engine.ParseResult {

	re := regexp.MustCompile(cityRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, c := range all {
		//fmt.Println("用户url：", string(c[1]))
		result.Items = append(result.Items, "User:"+string(c[2])) //用户名字

		name := string(c[2])
		fmt.Println(name)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(c[1]),
			// ParserFunc: engine.NilParser,
			// ParserFunc: ParseProfile,
			ParserFunc: func(c []byte) engine.ParseResult {
				return ParseProfile(c, name)

			},
		})
	}

	return result
}
