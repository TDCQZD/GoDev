package utils

import (
	"fmt"
	"net/http"
)

// Set和Add方法 创建Cookie
func handlerCookie(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{
		Name:     "lskd",
		Value:    "nshlz",
		HttpOnly: true,
	}

	cookie2 := http.Cookie{
		Name:     "yexiaochai",
		Value:    "HeartSword",
		HttpOnly: true,
	}

	//将Cookie发送给浏览器,即添加第一个Cookie
	w.Header().Set("Set-Cookie", cookie1.String())
	//再添加一个Cookie
	w.Header().Add("Set-Cookie", cookie2.String())

}

//SetCookie方法创建Cookie
func handlerSetCookie(w http.ResponseWriter, r *http.Request) {
	cookie1 := http.Cookie{
		Name:     "rpm",
		Value:    "pmjf",
		HttpOnly: true,
	}

	cookie2 := http.Cookie{
		Name:     "ngh",
		Value:    "wjsg",
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie1)
	http.SetCookie(w, &cookie2)

}

// 读取Cookie
func handlerReadCookie(w http.ResponseWriter, r *http.Request) {
	//获取请求头中的Cookie
	cookies := r.Header["Cookie"]
	//获取一个具体Cookie
	rpmCookie, _ := r.Cookie("rpm")
	fmt.Fprintln(w, cookies)
	fmt.Fprintln(w, rpmCookie)

}

// 设置Cookie的有效时间
func handlerMaxAge(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "yiyeshu",
		Value:    "bjwj",
		HttpOnly: true,
		MaxAge:   10,
	}

	http.SetCookie(w, &cookie)

}

// Session
func handlerSession(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    "bjwj",
		HttpOnly: true,
		MaxAge:   10,
	}

	http.SetCookie(w, &cookie)

}

func CookieDemo() {
	http.HandleFunc("/cookie", handlerCookie)
	http.HandleFunc("/setcookie", handlerSetCookie)
	http.HandleFunc("/readcookie", handlerReadCookie)
	http.HandleFunc("/maxage", handlerMaxAge)

	http.HandleFunc("/session", handlerSession)

	http.ListenAndServe(":8080", nil)
}
