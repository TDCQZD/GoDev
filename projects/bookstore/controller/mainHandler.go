package controller

import (
	"goweb_code/bookstore/model"
	"goweb_code/bookstore/model/modelDao"
	"net/http"
	"strconv"

	"text/template"
)

//处理首页
func MainHandler(w http.ResponseWriter, r *http.Request) {
	books, _ := modelDao.SearchAllBook()
	t := template.Must(template.ParseFiles("../views/index.html"))
	t.Execute(w, books)
}

//首页图书分页显示
func MainBookPageHandler(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	//获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}

	pageNoInt, _ := strconv.ParseInt(pageNo, 10, 64)
	minPriceFloat, _ := strconv.ParseFloat(minPrice, 64)
	maxPriceFloat, _ := strconv.ParseFloat(maxPrice, 64)

	var page *model.Page
	if minPriceFloat == 0 && maxPriceFloat == 0 {
		//调用bookdao中获取带分页的图书的函数
		page, _ = modelDao.QueryPageBooksByBack(pageNoInt, 4) //当前页数，每页数量
	} else {

		//调用bookdao中获取带分页和价格范围的图书的函数
		page, _ = modelDao.QueryBooksByPrices(pageNoInt, 4, minPriceFloat, maxPriceFloat)
		//将价格范围设置到page中
		page.MinPrice = minPriceFloat
		page.MaxPrice = maxPriceFloat
	}
	UserStatue.PageInfor = page

	if UserStatue.IsLogin { //已登录
		//获取Cookie
		userCookie, _ := r.Cookie("usersession")
		cookieValue := userCookie.Value
		seeion, err := modelDao.SearchSessionByUUID(cookieValue)
		if err == nil && seeion != nil { //查询到登录用户信息
			UserStatue.IsLogin = true
			UserStatue.UserName = seeion.UserName
			
		}
	}

	//解析模板文件
	t := template.Must(template.ParseFiles("../views/main.html"))
	//执行
	// t.Execute(w, page)
	t.Execute(w, UserStatue)
}
