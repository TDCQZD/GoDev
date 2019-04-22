package controller

import (
	"fmt"
	"goweb_code/bookstore/model"
	"goweb_code/bookstore/model/modelDao"
	"goweb_code/bookstore/utils"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

//用户支付
func Checkout(w http.ResponseWriter, r *http.Request) {
	if !UserStatue.IsLogin { //用户未登录
		t := template.Must(template.ParseFiles("../views/pages/user/login.html"))
		t.Execute(w, "用户未登录!")
		return
	}
	/*登录操作处理
	1、添加订单
	2、添加订单项
	3、清空购物车
	4、更新当前book的库存和销量
	*/
	carts, _ := modelDao.SearchCartByUserID(UserStatue.UserID)
	// 1、添加订单
	orderID := utils.CreateUUID()
	td := time.Now()
	time := td.Format("2006/01/02 15:04:05")
	order := &model.Oreders{
		ID:          orderID,
		Time:        time,
		TotalAmount: carts.TotalAmount,
		TotalCount:  carts.TotalCount,
		State:       0,
		UserID:      UserStatue.UserID,
	}
	//将订单保存到数据库中
	modelDao.AddOreder(order)
	// 2、添加订单项
	cartItems := carts.CartItems
	//遍历得到每一个购物项
	for _, v := range cartItems {
		//创建OrderItem
		orderItem := &model.OrederItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Books.Title,
			Author:  v.Books.Author,
			Price:   v.Books.Price,
			ImgPath: v.Books.ImgPath,
			OrderID: orderID,
		}
		//将购物项保存到数据库中
		modelDao.AddOrederItem(orderItem)
		// 4、更新当前book的库存和销量
		book := v.Books
		book.Sales = book.Sales + int(v.Count)
		book.Stock = book.Stock - int(v.Count)
		//更新图书的信息
		modelDao.UpdateBook(book)
	}
	// 3、清空购物车
	modelDao.ClearCart(carts.CartID)
	UserStatue.OrderID = orderID
	t := template.Must(template.ParseFiles("../views/pages/cart/checkout.html"))
	t.Execute(w, UserStatue)
}

// 获取当前用户所有订单
func QueryOrders(w http.ResponseWriter, r *http.Request) {
	if !UserStatue.IsLogin { //用户未登录
		// t := template.Must(template.ParseFiles("../views/pages/user/login.html"))
		// t.Execute(w, "用户未登录!")
		QueryAllOrders(w, r)
		return
	}
	//获取所有订单的函数
	orders, _ := modelDao.SearchOrdersByUserID(UserStatue.UserID)
	//解析模板
	t := template.Must(template.ParseFiles("../views/pages/order/order_manager.html"))
	//执行
	UserStatue.Oreders = orders
	t.Execute(w, UserStatue)
}

// 获取所有订单
func QueryAllOrders(w http.ResponseWriter, r *http.Request) {

	//获取所有订单的函数
	orders, _ := modelDao.SearchOrders()
	//解析模板
	t := template.Must(template.ParseFiles("../views/pages/order/order.html"))
	//执行
	UserStatue.Oreders = orders
	t.Execute(w, orders)
}

//查询订单详情
func QueryOrderItems(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderId")
	fmt.Println("orderid", orderid)
	orderItems, _ := modelDao.SearchOrderItems(orderid)
	// orderItems, _ := modelDao.SearchOrderItems("c3f671ff-b18b-42b4-4862-2f9aa2f74448")
	//解析模板
	t := template.Must(template.ParseFiles("../views/pages/order/order_info.html"))
	//执行
	UserStatue.OrederItems = orderItems
	t.Execute(w, UserStatue)
}

//修改订单状态
func UpdateOrderState(w http.ResponseWriter, r *http.Request) {
	orderid := r.FormValue("orderID")
	orderState := r.FormValue("state")
	state, _ := strconv.ParseInt(orderState, 10, 0)
	fmt.Println("orderState", state)
	orderItems, _ := modelDao.SearchOrdersByID(orderid)
	orderItems.State = state
	modelDao.UpdateOreder(orderItems)
	QueryOrders(w, r)
}
