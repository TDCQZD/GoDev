package modelDao

import (
	"fmt"
	"goweb_code/bookstore/model"
	"testing"
)

func TestCart(t *testing.T) {
	fmt.Println("测试CartDao 中的函数")
	// t.Run("测试Cart 添加：", testAddCart)
	// t.Run("测试Cart 修改：", testUpdateCart)
	t.Run("测试Cart 删除：", testDeleteCart)
	// t.Run("测试Cart 查询：", testSeachCarts)

}

func testAddCart(t *testing.T) {
	book := &model.Books{ID: 50, Price: 10}
	cartItem := &model.CartItem{
		Books: book,
		Count: 10,
	}
	cartItems := []*model.CartItem{}
	cartItems = append(cartItems, cartItem)
	cart := &model.Cart{
		CartItems:   cartItems,
		TotalCount:  1000,
		TotalAmount: 30000,
		UserID:      "2",
	}
	err := AddCart(cart)
	fmt.Println("添加结果是：", err)
}

func testUpdateCart(t *testing.T) {
	book := &model.Books{ID: 49, Price: 10}
	cartItem := &model.CartItem{
		Books: book,
		Count: 100,
	}
	cartItems := []*model.CartItem{}
	cartItems = append(cartItems, cartItem)
	cart := &model.Cart{
		CartItems:   cartItems,
		TotalCount:  1000,
		TotalAmount: 30000,
		CartID:      "102c6e3b-6ce8-4d0a-75c4-4e118a4dd360",
	}
	err := UpdateCart(cart)
	fmt.Println("结果是：", err)

}
func testDeleteCart(t *testing.T) {

	err := DeleteCart("11")
	fmt.Println("结果是：", err)

}

func testSeachCarts(t *testing.T) {
	cart, _ := SearchCartByUserID("2")
	fmt.Println("查询结果是：", cart)

}
