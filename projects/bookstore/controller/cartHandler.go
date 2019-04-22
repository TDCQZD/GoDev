package controller

import (
	"encoding/json"
	"fmt"
	"goweb_code/bookstore/model"
	"goweb_code/bookstore/model/modelDao"
	"goweb_code/bookstore/utils"
	"net/http"
	"strconv"
	"text/template"
)

/*购物车管理-添加
1、用户登录后才能进行购物车操作
2、判断当前用户是否已经创建购物车
2、判断当前用户购物车中是否已经存在同bookid的购物项
*/
func CartManagaement_add(w http.ResponseWriter, r *http.Request) {
	if !UserStatue.IsLogin { //用户未登录
		w.Write([]byte("用户未登录！"))
		return
	}
	/*登录操作处理*/
	//获取要添加的图书的id
	bookID := r.FormValue("bookId")
	bookid, _ := strconv.ParseInt(bookID, 10, 0)
	//根据图书的id获取图书信息
	book, _ := modelDao.SearchBook(int(bookid))

	//通过用户id获取购物车信息

	cart, _ := modelDao.SearchCartByUserID(UserStatue.UserID)
	if cart != nil { //该用户已创建购物车
		updateUserCart(book, cart)
	} else { //该用户没有创建购物车
		createUserCart(book)
	}

	w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车！"))
}

//用户创建购物车和购物项
func createUserCart(book *model.Books) {

	//1.创建购物车
	//创建cartid
	uuid := utils.CreateUUID()
	cart := &model.Cart{
		CartID: uuid,
		UserID: UserStatue.UserID,
	}
	//2.创建购物车中的购物项
	//声明一个CartItem类型的切片
	var cartItems []*model.CartItem
	cartItem := &model.CartItem{
		Books:  book,
		Count:  1,
		CartID: uuid,
		Amount: book.Price,
	}

	//将购物项添加到切片中
	cartItems = append(cartItems, cartItem)

	//3将切片设置到cart中
	cart.CartItems = cartItems
	//4.将购物车cart保存到数据库中
	modelDao.AddCart(cart)
	modelDao.AddCartItem(cartItem)
}

// 用户更新购物车

func updateUserCart(book *model.Books, cart *model.Cart) {

	cartItem, _ := modelDao.SearchCartItemByBookIDAndCartID(book.ID, cart.CartID)
	// 1、判断购物项是否存在
	if cartItem != nil { //该购物项已存在
		//1)该购物项数目加1
		cartItem.Count = cartItem.Count + 1
		cartItem.Books = book
		//2)更新该购物项
		modelDao.UpdateCartItem(cartItem)
	} else { //该购物项不存在，创建购物项并更新购物车
		createCartItem(book, cart)
	}
	// 2、更新购物车，此时CartItem表中数据已经最新
	carts, _ := modelDao.SearchCartItemByCartID(cart.CartID)
	cart.CartItems = carts
	modelDao.UpdateCart(cart)
}

//创建新的购物项并更新购物车
func createCartItem(book *model.Books, cart *model.Cart) {
	//创建购物车中的购物项
	cartItem := &model.CartItem{
		Books:  book,
		Count:  1,
		CartID: cart.CartID,
	}
	//将购物项添加到当前cart的切片中
	cart.CartItems = append(cart.CartItems, cartItem)
	//将新创建的购物项添加到数据库中
	modelDao.AddCartItem(cartItem)
}

//购物车信息查询
func CartManagaement_query(w http.ResponseWriter, r *http.Request) {
	if !UserStatue.IsLogin { //用户未登录
		t := template.Must(template.ParseFiles("../views/pages/user/login.html"))
		t.Execute(w, "用户未登录!")
		return
	}
	//通过用户id获取购物车信息
	cart, _ := modelDao.SearchCartByUserID(UserStatue.UserID)
	if cart != nil && cart.TotalCount > 0 { //该用户已创建购物车且物品总数量不为0
		UserStatue.CartInfor = cart
		t := template.Must(template.ParseFiles("../views/pages/cart/cart.html"))
		t.Execute(w, UserStatue)
	} else { //该用户没有购物车
		//用户状态
		userStatue := &model.UserStatue{
			IsLogin:  UserStatue.IsLogin,
			UserID:   UserStatue.UserID,
			UserName: UserStatue.UserName,
		}
		t := template.Must(template.ParseFiles("../views/pages/cart/cart.html"))
		t.Execute(w, userStatue)
	}

}

//清空购物车
func CartManagaement_delete(w http.ResponseWriter, r *http.Request) {
	//获取要删除的购物车的id
	cartID := r.FormValue("cartId")
	//清空购物车
	modelDao.ClearCart(cartID)
	CartManagaement_query(w, r)
	// //跳转到cart界面
	// t := template.Must(template.ParseFiles("../views/pages/cart/cart.html"))
	// //清空购物车,重构用户状态
	// userStatue := &model.UserStatue{
	// 	IsLogin:  UserStatue.IsLogin,
	// 	UserID:   UserStatue.UserID,
	// 	UserName: UserStatue.UserName,
	// }
	// t.Execute(w, userStatue)
}

//删除购物项
func DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	//获取要要删除的购物项ID
	cartItemId := r.FormValue("id")
	id, _ := strconv.ParseInt(cartItemId, 10, 0)
	fmt.Println("id:", id)
	modelDao.DeleteCartItemByID(id)
	//更新Cart
	carts, _ := modelDao.SearchCartItemByCartID(UserStatue.CartInfor.CartID)
	cart := &model.Cart{
		CartItems: carts,
		CartID:    UserStatue.CartInfor.CartID,
	}
	modelDao.UpdateCart(cart)
	CartManagaement_query(w, r)
}

//修改购物项-book数量
func UpdateCartItem(w http.ResponseWriter, r *http.Request) {

	// bookID := r.FormValue("bookID")
	// ibookID, _ := strconv.ParseInt(bookID, 10, 64)
	// fmt.Println("bookID:", bookID)
	//获取要更新的购物项的id
	cartItemID := r.FormValue("cartItemId")
	//将购物项的id转换为int64
	iCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	fmt.Println("iCartItemID:", iCartItemID)
	//获取用户输入的图书的数量
	bookCount := r.FormValue("bookCount")
	iBookCount, _ := strconv.ParseInt(bookCount, 10, 64)
	fmt.Println("bookCount:", bookCount)
	cartItem,_ := modelDao.SearchCartItemByID(int(iCartItemID))
	book, _ := modelDao.SearchBook(cartItem.Books.ID)
	//更新购物项
	// cartItem, _ = modelDao.SearchCartItemByBookIDAndCartID(int(ibookID), UserStatue.CartInfor.CartID)
	cartItem.Count = iBookCount
	cartItem.Books = book
	modelDao.UpdateCartItem(cartItem)

	//更新购物车中的图书的总数量和总金额
	carts, _ := modelDao.SearchCartItemByCartID(UserStatue.CartInfor.CartID)
	cart := &model.Cart{
		CartItems: carts,
		CartID:    UserStatue.CartInfor.CartID,
	}
	modelDao.UpdateCart(cart)

	/*修改购物项后全局修改*/
	//调用获取购物项信息的函数再次查询购物车信息
	// CartManagaement_query(w, r)

	/*修改购物项后局部修改*/
	//调用获取购物项信息的函数再次查询购物车信息
	cart, _ = modelDao.SearchCartByUserID(UserStatue.UserID)
	//获取购物车中图书的总数量
	totalCount := cart.TotalCount
	//获取购物车中图书的总金额
	totalAmount := cart.TotalAmount

	var amount float64
	//获取购物车中更新的购物项中的金额小计
	cIs := cart.CartItems
	for _, v := range cIs {
		if iCartItemID == v.ID {
			//这个就是我们寻找的购物项，此时获取当前购物项中的金额小计
			amount = v.Amount
		}
	}
	//创建Data结构
	data := model.Data{
		Amount:      amount,
		TotalAmount: totalAmount,
		TotalCount:  totalCount,
	}
	//将data转换为json字符串
	json, _ := json.Marshal(data)
	//响应到浏览器
	w.Write(json)
}
