package modelDao

import (
	"fmt"
	"goweb_code/bookstore/DB"
	"goweb_code/bookstore/model"
)

var cartDB = DB.DB

func AddCart(cart *model.Cart) error {

	//1.sql语句
	sqlStr := `INSERT INTO cart (ID,ZSM,ZJE,YHID) VALUES (?,?,?,?)`
	//2.执行
	_, err := cartDB.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

func UpdateCart(cart *model.Cart) error {

	//1.sql语句
	sqlStr := `UPDATE cart set ZSM = ? ,ZJE = ? WHERE ID = ?`
	//2.执行
	_, err := cartDB.Exec(sqlStr, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

func DeleteCart(cartID string) error {

	//1.sql语句
	sqlStr := `DELETE FROM cart WHERE ID = ?`
	//2.执行
	_, err := cartDB.Exec(sqlStr, cartID)
	if err != nil {
		fmt.Println("执行出现异常：", err)
		return err
	}
	return nil
}

func SearchCartByUserID(userID string) (*model.Cart, error) {
	sqlStr := `SELECT ID, ZSM, ZJE, YHID FROM cart WHERE YHID =  ?`

	row := cartDB.QueryRow(sqlStr, userID)
	carts := &model.Cart{}

	err := row.Scan(&carts.CartID, &carts.TotalCount, &carts.TotalAmount, &carts.UserID)
	if err != nil {
		return nil, err
	}
	//查询cartItems
	cartItems, _ := SearchCartItemByCartID(carts.CartID)
	for _, v := range cartItems {
		book, _ := SearchBook(v.Books.ID)
		v.Books = book
	}
	carts.CartItems = cartItems

	return carts, nil
}

/*
清空购物车
1、删除购物项
2、删除购物车
*/
func ClearCart(cartID string) error {
	err := DeleteCartItem(cartID)

	if err != nil {
		return err
	}

	err = DeleteCart(cartID)
	if err != nil {
		return err
	}
	return nil
}
