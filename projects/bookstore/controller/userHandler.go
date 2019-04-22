package controller

import (
	"fmt"
	"goweb_code/bookstore/model"
	"goweb_code/bookstore/model/modelDao"
	"goweb_code/bookstore/utils"
	"net/http"
	"text/template"
)

var UserStatue = &model.UserStatue{
	IsLogin: false,
}

/*登录处理*/
func Login(w http.ResponseWriter, r *http.Request) {
	userName := r.PostFormValue("username")
	password := r.PostFormValue("password")
	flag, user, _ := modelDao.LoginAccount(userName, password)
	if flag { //登录成功

		uuid := getUUID(user, r)
		/*设置Cookie*/
		cookie := http.Cookie{
			Name:     "usersession",
			Value:    uuid,
			HttpOnly: true,
		}

		http.SetCookie(w, &cookie)
		UserStatue.IsLogin = true
		UserStatue.UserID = user.ID
		UserStatue.UserName = user.Username
		t := template.Must(template.ParseFiles("../views/pages/user/login_success.html"))
		t.Execute(w, userName)
	} else { //登录失败
		t := template.Must(template.ParseFiles("../views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}

func getUUID(user *model.Users, r *http.Request) string {
	/*组装Session*/

	uuid := utils.CreateUUID()
	session := model.Seesion{
		UUID:     uuid,
		UserName: user.Username,
		UserID:   user.ID,
	}

	//获取Cookie
	userCookie, _ := r.Cookie("usersession")

	if UserStatue.IsLogin { //已登录状态
		uuid = userCookie.Value
	} else {
		modelDao.AddSession(&session)
	}
	return uuid
}

/*用户注销*/
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	userCookie, _ := r.Cookie("usersession")
	// fmt.Println("uuid", userCookie.Value)
	if userCookie != nil { //已登录
		cookieValue := userCookie.Value
		modelDao.DeleteSession(cookieValue) //删除session
		//设置Cookie失效
		userCookie.MaxAge = -1
		//告诉浏览器Cookie失效
		http.SetCookie(w, userCookie)
	}
	UserStatue = &model.UserStatue{
		IsLogin: false,
	}
	//跳转到首页
	MainBookPageHandler(w, r)
	// fmt.Println("user", UserStatue.IsLogin)
}

/*注册处理*/
func Register(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	fmt.Println("RawQuery:", r.URL.RawQuery)
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user := model.Users{
		Username: username,
		Password: password,
		Email:    email,
	}
	// fmt.Println("Register user:", user)
	flag, err := modelDao.RegisterAccount(&user)

	if flag { //注册成功
		t := template.Must(template.ParseFiles("../views/pages/user/login.html"))
		t.Execute(w, "注册成功")
	} else { //注册失败
		t := template.Must(template.ParseFiles("../views/pages/user/regist.html"))
		t.Execute(w, err)
	}
}

/*验证用户存在*/
func VerifyName(w http.ResponseWriter, r *http.Request) {

	userName := r.PostFormValue("username")

	flag, _ := modelDao.VerifyName(userName)
	if flag { //用户已存在
		w.Write([]byte("该用户已注册！"))
	} else { //用户不存在
		w.Write([]byte("<font style='color:green'>该用户未注册！</font>"))
	}
}
