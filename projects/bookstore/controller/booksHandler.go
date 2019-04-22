package controller

import (
	"fmt"
	"goweb_code/bookstore/model"
	"goweb_code/bookstore/model/modelDao"
	"net/http"
	"strconv"
	"text/template"
)

/*查询所有书籍*/
func QueryAllBooks(w http.ResponseWriter, r *http.Request) {
	//调用bookdao中获取所有图书的函数
	books, _ := modelDao.SearchAllBook()
	//解析模板文件
	t := template.Must(template.ParseFiles("../views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, books)
}

/*添加书籍书籍*/
func AddBook(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	priceF, _ := strconv.ParseFloat(price, 64)
	salesI, _ := strconv.ParseInt(sales, 10, 0)
	stockI, _ := strconv.ParseInt(stock, 10, 0)

	book := &model.Books{
		Title:   title,
		Author:  author,
		Price:   priceF,
		Sales:   int(salesI),
		Stock:   int(stockI),
		ImgPath: "/static/img/default.jpg",
	}

	modelDao.AddBook(book)
	QueryAllBooks(w, r)
}

/*修改书籍信息
1、获取ID 书籍
2、数据更新
*/
func EditBook(w http.ResponseWriter, r *http.Request) {
	book_id := r.FormValue("bookId")
	id, _ := strconv.ParseInt(book_id, 10, 0)
	fmt.Println("bookId", int(id))
	book, _ := modelDao.SearchBook(int(id))
	//解析模板文件
	t := template.Must(template.ParseFiles("../views/pages/manager/book_edit.html"))
	//执行
	t.Execute(w, book)
}

// 修改书籍操作
func UpdateBook(w http.ResponseWriter, r *http.Request) {

	id := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	idI, _ := strconv.ParseInt(id, 10, 0)
	priceF, _ := strconv.ParseFloat(price, 64)
	salesI, _ := strconv.ParseInt(sales, 10, 0)
	stockI, _ := strconv.ParseInt(stock, 10, 0)

	book := model.Books{
		ID:      int(idI),
		Title:   title,
		Author:  author,
		Price:   priceF,
		Sales:   int(salesI),
		Stock:   int(stockI),
		ImgPath: "/static/img/default.jpg",
	}
	modelDao.UpdateBook(&book)
	QueryAllBooks(w, r)

}

// 添加书籍和修改书籍整合
func BookEdit(w http.ResponseWriter, r *http.Request) {

	id := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	idI, _ := strconv.ParseInt(id, 10, 0)
	priceF, _ := strconv.ParseFloat(price, 64)
	salesI, _ := strconv.ParseInt(sales, 10, 0)
	stockI, _ := strconv.ParseInt(stock, 10, 0)

	book := model.Books{
		ID:      int(idI),
		Title:   title,
		Author:  author,
		Price:   priceF,
		Sales:   int(salesI),
		Stock:   int(stockI),
		ImgPath: "/static/img/default.jpg",
	}
	//通过判断bookID,来确定操作
	if book.ID > 0 { //修改
		modelDao.UpdateBook(&book)
	} else { //添加
		modelDao.AddBook(&book)
	}
	QueryAllBooks(w, r)

}

/*删除书籍*/
func DeleteBook(w http.ResponseWriter, r *http.Request) {

	book_id := r.FormValue("bookId")
	id, _ := strconv.ParseInt(book_id, 10, 0)
	fmt.Println("bookId", int(id))
	modelDao.DeleteBook(int(id))
	QueryAllBooks(w, r)
}

/*图书管理分页：后台分页*/
func PageBooksByBack(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	pageNoInt, _ := strconv.ParseInt(pageNo, 10, 64)
	//调用bookdao中获取带分页的图书的函数
	page, _ := modelDao.QueryPageBooksByBack(pageNoInt, 4) //当前页数，每页数量
	//解析模板文件
	t := template.Must(template.ParseFiles("../views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, page)
}

/*图书管理分页：前台分页*/
func PageBooksByFront(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	pageNoInt, _ := strconv.ParseInt(pageNo, 10, 64)

	//调用bookdao中获取带分页的图书的函数
	page, _ := modelDao.QueryPageBooksByFront(pageNoInt, 4) //当前页数，每页数量

	//解析模板文件
	t := template.Must(template.ParseFiles("../views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, page)
}
