package modelDao

import (
	"fmt"
	"goweb_code/bookstore/DB"
	"goweb_code/bookstore/model"
)

var orderItemDB = DB.DB

func AddOrederItem(orderItem *model.OrederItem) error {

	//1.sql语句
	sqlStr := `INSERT INTO order_item (TITLE,AUTHOR,PRICE,IMG,COUNT,AMOUNT,ORDERID) VALUES(?,?,?,?,?,?,?)`
	//2.执行
	_, err := orderItemDB.Exec(sqlStr, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.Count, orderItem.Amount, orderItem.OrderID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

func DeleteOrederItem(orderID string) error {

	//1.sql语句
	sqlStr := `DELETE FROM order_item WHERE orderID = ?`
	//2.执行
	_, err := orderItemDB.Exec(sqlStr, orderID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//查询订单号下所有订单
func SearchOrderItems(orderID string) ([]*model.OrederItem, error) {
	sqlStr := `SELECT ID,TITLE,AUTHOR,PRICE,IMG,COUNT,AMOUNT,ORDERID FROM order_item WHERE ORDERID = ?`

	//执行
	rows, err := orderItemDB.Query(sqlStr, orderID)
	if err != nil {
		return nil, err
	}
	//创建orederItems切片
	var orderItems []*model.OrederItem
	for rows.Next() {
		orderItem := &model.OrederItem{}
		err := rows.Scan(&orderItem.ID, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.Count, &orderItem.Amount, &orderItem.OrderID)
		if err != nil {
			return nil, err
		}

		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}

