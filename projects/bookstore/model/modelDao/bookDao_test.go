package modelDao

import (
	"fmt"
	"goweb_code/bookstore/model"
	"testing"
)

func TestBook(t *testing.T) {
	fmt.Println("测试bookdao中的函数")
	// t.Run("测试查询所有Books：", testSearchBooks)
	// t.Run("测试添加Books：", testAddBook)
	// t.Run("测试删除Books：", testDeleteBook)
	t.Run("测试修改Books：", testUpdateBook)

}

func testSearchBooks(t *testing.T) {
	books, _ := SearchAllBook()
	fmt.Println("获取书籍信息是：", books)
}

func testAddBook(t *testing.T) {
	book := model.Books{
		Title:   "《三国演义》",
		Author:  "罗贯中",
		Price:   100,
		Sales:   100,
		Stock:   100,
		ImgPath: "static/img/default.jpg",
	}
	err := AddBook(&book)
	fmt.Println("获取书籍信息是：", err)
}
func testUpdateBook(t *testing.T) {
	book := model.Books{
		ID:      32,
		Title:   "三国演义",
		Author:  "罗贯中",
		Price:   100,
		Sales:   10000,
		Stock:   100,
		ImgPath: "static/img/default.jpg",
	}
	err := UpdateBook(&book)
	fmt.Println("获取书籍信息是：", err)
}
func testDeleteBook(t *testing.T) {

	err := DeleteBook(33)
	fmt.Println("获取书籍信息是：", err)
}
