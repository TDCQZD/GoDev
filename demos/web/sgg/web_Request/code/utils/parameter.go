package utils

import (
	"fmt"
	"net/http"
)

func handlerParameter(w http.ResponseWriter, r *http.Request) {
	/*1、Form字段*/
	//解析表单，在调用r.Form之前必须执行该操作
	r.ParseForm()
	//获取请求参数
	//如果form表单的action属性的URL地址中也有与form表单参数名相同的请求参数，
	//那么参数值都可以得到，并且form表单中的参数值在ULR的参数值的前面
	fmt.Fprintln(w, "请求参数有：", r.Form)

	/*2、PostForm字段*/
	fmt.Fprintln(w, "POST请求的form表单中的请求参数有：", r.PostForm)

	/*3、FormValue方法和PostFormValue方法*/
	//通过直接调用FormValue方法和PostFormValue方法直接获取请求参数的值
	fmt.Fprintln(w, "URL中的user请求参数的值是：", r.FormValue("name"))
	fmt.Fprintln(w, "Form表单中的username请求参数的值是：", r.PostFormValue("username"))

}

/*MultipartForm字段*/
func handlerMultipartForm(w http.ResponseWriter, r *http.Request) {
	//解析表单
	r.ParseMultipartForm(1024)
	//打印表单数据
	fmt.Fprintln(w, r.MultipartForm)

}

func ParameterDemo() {
	http.HandleFunc("/paramter", handlerParameter)
	http.HandleFunc("/multipart", handlerMultipartForm)
	http.ListenAndServe(":8080", nil)
}
