package engine

// ParseResult 解析后返回的结果
type ParseResult struct {
	Requests []Request
	// Items    []interface{}
	Items []Item
}

// Request 返回的结果结构体
type Request struct {
	Url        string                   //解析出来的URL
	ParserFunc func([]byte) ParseResult //处理这个URL所需要的函数
}

// NilParser 空解析器
func NilParser([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url     string //URL
	Type    string //存储到ElasticSearch时的type
	Id      string //用户Id
	Payload interface{}
}
