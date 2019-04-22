package modelDao

import (
	"fmt"
	"goweb_code/bookstore/DB"
	"goweb_code/bookstore/model"
)

var orderDB = DB.DB

func AddOreder(order *model.Oreders) error {

	//1.sql语句
	sqlStr := `INSERT INTO orders (ID,STATE,SJ,COUNT,AMOUNT,YHID) VALUES (?,?,?,?,?,?)`
	//2.执行
	_, err := orderDB.Exec(sqlStr, order.ID, order.State, order.Time, order.TotalCount, order.TotalAmount, order.UserID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

func UpdateOreder(order *model.Oreders) error {

	//1.sql语句
	sqlStr := `UPDATE  orders SET STATE = ?,SJ = ? ,COUNT = ? ,AMOUNT = ? WHERE ID = ?`
	//2.执行
	_, err := orderDB.Exec(sqlStr, order.State, order.Time, order.TotalCount, order.TotalAmount, order.ID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

func DeleteOreder(orderID string) error {

	//1.sql语句
	sqlStr := `DELETE FROM orders WHERE ID = ?`
	//2.执行
	_, err := orderDB.Exec(sqlStr, orderID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

//查询用户所有订单
func SearchOrdersByUserID(userID string) ([]*model.Oreders, error) {
	sqlStr := `SELECT ID,STATE,SJ,COUNT,AMOUNT,YHID FROM orders WHERE YHID = ?`

	//执行
	rows, err := orderDB.Query(sqlStr, userID)
	if err != nil {
		return nil, err
	}
	//创建orders切片
	var orders []*model.Oreders
	for rows.Next() {
		order := &model.Oreders{}
		err := rows.Scan(&order.ID, &order.State, &order.Time, &order.TotalCount, &order.TotalAmount, &order.UserID)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	return orders, nil
}
//查询所有订单
func SearchOrders() ([]*model.Oreders, error) {
	sqlStr := `SELECT ID,STATE,SJ,COUNT,AMOUNT,YHID FROM orders `

	//执行
	rows, err := orderDB.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建orders切片
	var orders []*model.Oreders
	for rows.Next() {
		order := &model.Oreders{}
		err := rows.Scan(&order.ID, &order.State, &order.Time, &order.TotalCount, &order.TotalAmount, &order.UserID)
		if err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}
	return orders, nil
}
func SearchOrdersByID(ID string) (*model.Oreders, error) {
	sqlStr := `SELECT ID,STATE,SJ,COUNT,AMOUNT,YHID FROM orders WHERE ID = ?`
	row := orderDB.QueryRow(sqlStr, ID)
	order := &model.Oreders{}
	err := row.Scan(&order.ID, &order.State, &order.Time, &order.TotalCount, &order.TotalAmount, &order.UserID)
	if err != nil {
		return nil, err
	}
	return order, nil
}
