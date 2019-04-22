package modelDao

import (
	"goweb_code/bookstore/DB"
	"goweb_code/bookstore/model"
	_ "strconv"
)

var pageBookDB = DB.DB

/*Books 分页处理*/
//分页处理——后台分页
/* 参数说明：
pageNo:当前页
pageSize：每页数量
*/
func QueryPageBooksByBack(pageNo int64, pageSize int64) (pageStruct *model.Page, err error) {

	//将页码转换为int64类型
	// pageNo, _ := strconv.ParseInt(pageNos, 10, 64)

	/*统计books数量*/
	sqlStr := `SELECT count(*) FROM books`
	var count int64
	row := pageBookDB.QueryRow(sqlStr)
	row.Scan(&count)

	/* */
	//分页显示共需要多少页
	var totalPage int64
	if count%pageSize == 0 {
		totalPage = count / pageSize
	} else {
		totalPage = count/pageSize + 1
	}
	/*
		1、if totalRecord%pageSize == 0 {
			totalPageNo = totalRecord / pageSize
		} else {
			totalPageNo = totalRecord/pageSize + 1
		}
		2、math.Ceil( totalRecord/PageSize )
		3、(totalRecord - 1)/pageSize + 1
	*/
	// totalPage := (count-1)/pageSize + 1

	//获取当前页中的图书
	sqlStr = "select id,title,author,price,sales,stock,img_path from books limit ?,?"

	rows, err := pageBookDB.Query(sqlStr, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Books
	for rows.Next() {
		book := &model.Books{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}

	pageStruct = &model.Page{
		Books:       books,
		PageNo:      pageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPage,
		TotalRecord: count,
	}
	return

}

//分页处理——前端分页
func QueryPageBooksByFront(pageNo int64, pageSize int64) (pageStruct *model.Page, err error) {

	var count int64
	var rowNo int64
	//分页显示共需要多少页
	var totalPage int64

	//获取所有的图书，和图书数目
	sqlStr := `SELECT * FROM(
		SELECT Y.*,@rownum:=@rownum+1 as ROWNO FROM(
		select b.id,b.title,b.author,b.price,b.sales,b.stock,b.img_path ,C.COUNT from (SELECT COUNT(1) as count FROM books )  C,books b
		)Y,(select @rownum:=0) z ) X `

	rows, err := pageBookDB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Books
	// var mapBooks map[string][*model.Books];
	// mapBooks = make(map[string][*model.Books], 100)
	for rows.Next() {
		book := &model.Books{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath, &count, &rowNo)
		totalPage = (count-1)/pageSize + 1
		//将book添加到books中
		books = append(books, book)
	}

	pageStruct = &model.Page{
		Books:       books,
		PageNo:      pageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPage,
		TotalRecord: count,
	}
	return
}

//价格搜索
func QueryBooksByPrices(pageNo int64, pageSize int64, minPrice float64, maxPrice float64) (pageStruct *model.Page, err error) {
	//将页码转换为int64类型
	// pageNo, _ := strconv.ParseInt(pageNos, 10, 64)

	/*统计books数量*/
	sqlStr := "select count(*) from books where price between ? and ?"
	var count int64
	row := pageBookDB.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&count)

	var totalPage int64
	if count%pageSize == 0 {
		totalPage = count / pageSize
	} else {
		totalPage = count/pageSize + 1
	}

	//获取当前页中的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books where price between ? and ? limit ?,?"
	//执行
	rows, err := pageBookDB.Query(sqlStr2, minPrice, maxPrice, (pageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	var books []*model.Books
	for rows.Next() {
		book := &model.Books{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}

	pageStruct = &model.Page{
		Books:       books,
		PageNo:      pageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPage,
		TotalRecord: count,
	}
	return
}
