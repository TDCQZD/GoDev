package modelDao

import (
	"fmt"
	"goweb_code/bookstore/DB"
	"goweb_code/bookstore/model"
)

var bookDB = DB.DB

//添加数据
func AddBook(book *model.Books) (err error) {

	//1.sql语句
	sqlStr := `insert into books (title, author, price,sales,stock,img_path) values (?, ?, ?, ?, ?, ?)`
	//2.执行
	res, err := bookDB.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return
	}
	fmt.Println("执行结果：", res)
	return
}

//修改数据
func UpdateBook(book *model.Books) (err error) {
	//1.sql语句
	sqlStr := `UPDATE books SET  title = ?, author  = ?, price  = ?,sales  = ?,stock = ?,img_path = ? where id = ?`
	//2.执行
	res, err := bookDB.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath, book.ID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return
	}
	fmt.Println("执行结果：", res)
	return nil
}



//查询单条数据
func SearchBook(id int) (*model.Books, error) {
	
	sqlStr := `select id, title, author, price,sales,stock,img_path from books where id = ?`

	row := bookDB.QueryRow(sqlStr, id)
	book := &model.Books{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if err != nil {
		return nil, err
	}
	return book, nil
}

//查询所有数据
func SearchAllBook() ([]*model.Books, error) {
	sqlStr := `select id, title, author, price,sales,stock,img_path from books `

	//执行
	rows, err := bookDB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建Books切片
	var books []*model.Books
	for rows.Next() {
		book := &model.Books{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}
	return books, nil
}

//删除数据
func DeleteBook(id int) (err error) {
	//1.sql语句
	sqlStr := `DELETE FROM books WHERE id = ?`
	//2.执行
	res, err := bookDB.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return
	}
	fmt.Println("执行结果：", res)
	return nil
}


