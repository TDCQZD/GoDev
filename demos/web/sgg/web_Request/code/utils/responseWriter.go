package utils

import (
	"encoding/json"
	"goweb_code/web_Request/code/model"
	"net/http"
)

/*客户端响应*/
func handlerHTML(w http.ResponseWriter, r *http.Request) {
	html := `<html>
				<head>
					<title>测试响应内容为网页</title>
					<meta charset="utf-8"/>
				</head>
				<body>
					我是以网页的形式响应过来的！
				</body>
			</html>`
	w.Write([]byte(html))

}

/*客户端响应——JSON*/
func handlerJSON(w http.ResponseWriter, r *http.Request) {
	//设置响应内容的类型
	w.Header().Set("Content-Type", "application/json")
	//创建User
	user := model.User{
		ID:       1,
		Username: "admin",
		Password: "123456",
		Email:    "admin@atguigiu.com",
	}
	//将User转换为Json个数
	json, _ := json.Marshal(user)
	//将json格式的数据响应给客户端
	w.Write(json)
}

/*客户端响应——重定向*/
func handlerLocation(w http.ResponseWriter, r *http.Request) {
	//设置响应头中的Location属性
	w.Header().Set("Location", "https://www.baidu.com")
	//设置响应的状态码
	w.WriteHeader(302)
}

func ResponeDemo() {
	http.HandleFunc("/html", handlerHTML)
	http.HandleFunc("/json", handlerJSON)
	http.HandleFunc("/location", handlerLocation)
	http.ListenAndServe(":8080", nil)
}
