package main

import (
	"goweb_code/bookstore/controller"
	"net/http"
)

func main() {
	//静态文件处理
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../views/static"))))
	//html文件处理
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("../views/pages"))))

	/*首页*/
	http.HandleFunc("/main", controller.MainHandler)                 //首页
	http.HandleFunc("/mainbookpage", controller.MainBookPageHandler) //首页图书分页显示

	/*用户操作*/
	http.HandleFunc("/login", controller.Login)              //登录
	http.HandleFunc("/logout", controller.Logout)            //登录
	http.HandleFunc("/regist", controller.Register)          //注册
	http.HandleFunc("/verifyName", controller.VerifyName)    //验证用户是否已存在
	http.HandleFunc("/querybooks", controller.QueryAllBooks) //验证用户是否已存在

	/*书籍管理操作*/
	http.HandleFunc("/addbook", controller.AddBook)       //添加书籍操作
	http.HandleFunc("/editbook", controller.EditBook)     //编辑书籍界面
	http.HandleFunc("/updatebook", controller.UpdateBook) //修改书籍操作
	//添加和修改整合
	http.HandleFunc("/bookedit", controller.BookEdit) //修改书籍操作

	http.HandleFunc("/deletebook", controller.DeleteBook) //删除书籍

	/*图书管理分页显示*/
	http.HandleFunc("/bookpagebyback", controller.PageBooksByBack) //图书管理分页：后台分页
	// http.HandleFunc("/bookpagebyfront", controller.PageBooksByFront) //图书管理分页：前台分页

	/*购物车管理*/
	http.HandleFunc("/cartadd", controller.CartManagaement_add)      //购物车添加管理
	http.HandleFunc("/cartquery", controller.CartManagaement_query)  //购物车查询管理
	http.HandleFunc("/clearcart", controller.CartManagaement_delete) //清空购物车
	http.HandleFunc("/deletecartitem", controller.DeleteCartItem)    //删除购物项
	http.HandleFunc("/updateBookCount", controller.UpdateCartItem)   //修改购物项
	/*订单管理界面*/
	http.HandleFunc("/apply", controller.Checkout)            //用户支付
	http.HandleFunc("/queryOrders", controller.QueryOrders)   //查询订单
	http.HandleFunc("/orderInfo", controller.QueryOrderItems) //订单详情
	http.HandleFunc("/updateOrderState", controller.UpdateOrderState) //订单详情

	http.ListenAndServe(":8080", nil)
}
