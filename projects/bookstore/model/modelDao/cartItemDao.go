package modelDao

import (
	"fmt"
	"goweb_code/bookstore/DB"
	"goweb_code/bookstore/model"
)

var cartItemDB = DB.DB

func AddCartItem(cartItem *model.CartItem) error {

	//1.sql语句
	sqlStr := `INSERT INTO cart_item (BOOKID, SM, ZJE, CARTID) VALUES (?,?,?,?)`
	//2.执行
	res, err := cartItemDB.Exec(sqlStr, cartItem.Books.ID, cartItem.Count, cartItem.GetAmount(), cartItem.CartID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	//3、同步更新Carts

	fmt.Println("执行结果：", res)
	return nil
}

func UpdateCartItem(cartItem *model.CartItem) error {

	//1.sql语句
	sqlStr := `UPDATE  cart_item SET  SM = ? , ZJE = ? WHERE BOOKID = ? AND CARTID = ?`
	//2.执行
	res, err := cartItemDB.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Books.ID, cartItem.CartID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	//3、同步更新Carts

	fmt.Println("执行结果：", res)
	return nil
}

func DeleteCartItem(cartid string) error {

	//1.sql语句
	sqlStr := `DELETE FROM cart_item WHERE CARTID = ?`
	//2.执行
	res, err := cartItemDB.Exec(sqlStr, cartid)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}

	fmt.Println("执行结果：", res)
	return nil
}

func DeleteCartItemByID(ID int64) error {

	//1.sql语句
	sqlStr := `DELETE FROM cart_item WHERE ID = ?`
	//2.执行
	res, err := cartItemDB.Exec(sqlStr, ID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	//3、同步更新Carts

	fmt.Println("执行结果：", res)
	return nil
}

func SearchCartItemByBookIDAndCartID(bookid int, cartID string) (*model.CartItem, error) {
	sqlStr := `SELECT ID,BOOKID, SM, ZJE, CARTID FROM cart_item WHERE BOOKID = ? AND CARTID = ?`

	row := cartItemDB.QueryRow(sqlStr, bookid, cartID)
	book := &model.Books{}
	cartitem := &model.CartItem{
		Books: book,
	}
	err := row.Scan(&cartitem.ID, &cartitem.Books.ID, &cartitem.Count, &cartitem.Amount, &cartitem.CartID)
	if err != nil {
		return nil, err
	}
	return cartitem, nil
}

func SearchCartItemByCartID(cartID string) ([]*model.CartItem, error) {
	sqlStr := `SELECT ID,BOOKID, SM, ZJE, CARTID FROM cart_item WHERE  CARTID = ?`
	rows, _ := cartItemDB.Query(sqlStr, cartID)
	var cartItems []*model.CartItem

	for rows.Next() {
		book := &model.Books{}
		cartitem := &model.CartItem{
			Books: book,
		}
		err := rows.Scan(&cartitem.ID, &cartitem.Books.ID, &cartitem.Count, &cartitem.Amount, &cartitem.CartID)
		if err != nil {
			return nil, err
		}
		// fmt.Println("SearchCartItemByCartID查询结果是：", cartitem.Count)
		cartItems = append(cartItems, cartitem)
	}
	return cartItems, nil
}

func SearchCartItemByID(id int) (*model.CartItem, error) {
	sqlStr := `SELECT ID,BOOKID, SM, ZJE, CARTID FROM cart_item WHERE ID = ? `

	row := cartItemDB.QueryRow(sqlStr, id)
	book := &model.Books{}
	cartitem := &model.CartItem{
		Books: book,
	}
	err := row.Scan(&cartitem.ID, &cartitem.Books.ID, &cartitem.Count, &cartitem.Amount, &cartitem.CartID)
	if err != nil {
		return nil, err
	}
	return cartitem, nil
}