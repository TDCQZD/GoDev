package model

/*分页处理
1、前端分页
2、后台分页
*/

//Page 结构
type Page struct {
	Books       []*Books //每页查询出来的图书存放的切片
	PageNo      int64    //当前页
	PageSize    int64    //每页显示的条数
	TotalPageNo int64    //总页数，通过计算得到
	TotalRecord int64    //总记录数，通过查询数据库得到
	MinPrice    float64   // 价格查询 最大价格
	MaxPrice    float64   // 价格查询 最小价格
}

//IsHasPrev 判断是否有上一页
func (p *Page) IsHasPrev() bool {
	return p.PageNo > 1
}

//IsHasNext 判断是否有下一页
func (p *Page) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

//GetPrevPageNo 获取上一页
func (p *Page) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	} else {
		return 1
	}
}

//GetNextPageNo 获取下一页
func (p *Page) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	} else {
		return p.TotalPageNo
	}
}
