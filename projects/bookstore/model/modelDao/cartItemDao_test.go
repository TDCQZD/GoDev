package modelDao

import (
	"fmt"
	"goweb_code/bookstore/model"

	"testing"
)

func TestCartItem(t *testing.T) {
	fmt.Println("测试CartItem中的函数")
	// t.Run("测试添加CartItem：", testAddCartItem)
	// t.Run("测试修改CartItem：", testUpdateCartItem)
	t.Run("测试删除CartItem：", testDeleteCartItem)
	// t.Run("测试查询CartItemByBookIDAndCartID：", testSearchCartItemByBookIDAndCartID)
	// t.Run("测试查询CartItemByCartID：", testSearchCartItemByCartID)
}

func testAddCartItem(t *testing.T) {
	book := &model.Books{ID: 52, Price: 30}
	cartItem := &model.CartItem{
		Amount: 1000,
		CartID:"e8090e53-68f0-49d8-76c4-6e7fccd150ab",
		Count:  10,
		Books:  book,
	}
	err := AddCartItem(cartItem)
	//同步更新carts
	fmt.Println("添加结果是：", err)
}

func testUpdateCartItem(t *testing.T) {
	book := &model.Books{ID: 52, Price: 30}
	cartItem := &model.CartItem{
		Amount: 1000,
		CartID: "6",
		Count:  30,
		Books:  book,
	}
	err := UpdateCartItem(cartItem)
	//同步更新carts
	fmt.Println("结果是：", err)
}

func testDeleteCartItem(t *testing.T) {

	err := DeleteCartItem("6")
	fmt.Println("结果是：", err)
}

func testSearchCartItemByBookIDAndCartID(t *testing.T) {

	cartitem, _ := SearchCartItemByBookIDAndCartID(52, "6")
	fmt.Println("查询结果是：", cartitem.Count)
}

func testSearchCartItemByCartID(t *testing.T) {
	cartitems, _ := SearchCartItemByCartID("e8090e53-68f0-49d8-76c4-6e7fccd150ab")
	for _, v := range cartitems {
		fmt.Println("查询结果是：", v.Count)
	}

}
