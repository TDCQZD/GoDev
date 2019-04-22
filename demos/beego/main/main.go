package main

import (
	"github.com/astaxie/beego" //1、首先我们导入了包 github.com/astaxie/beego，beego 包中会初始化一个 BeeAPP 的应用和一些参数。
)

/*
beego 快速入门示例
*/

// 2、定义 Controller，这里我们定义了一个 struct 为 MainController，充分利用了 Go 语言的组合的概念，匿名包含了 beego.Controller，这样我们的 MainController 就拥有了 beego.Controller 的所有方法。
type MainController struct {
	beego.Controller
}

// 3、定义 RESTful 方法，通过匿名组合之后，其实目前的 MainController 已经拥有了 Get、Post、Delete、Put 等方法，这些方法是分别用来对应用户请求的 Method 函数，如果用户发起的是 POST 请求，那么就执行 Post 函数。所以这里我们定义了 MainController 的 Get 方法用来重写继承的 Get 函数，这样当用户发起 GET 请求的时候就会执行该函数。
func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

// 4、定义 main 函数，所有的 Go 应用程序和 C 语言一样都是 main 函数作为入口，所以我们这里定义了我们应用的入口。
func main() {
	// 5、Router 注册路由，路由就是告诉 beego，当用户来请求的时候，该如何去调用相应的 Controller，这里我们注册了请求 / 的时候，请求到 MainController。这里我们需要知道，Router 函数的两个参数函数，第一个是路径，第二个是 Controller 的指针。
	beego.Router("/", &MainController{})
	// 6、Run 应用，最后一步就是把在步骤 1 中初始化的 BeeApp 开启起来，其实就是内部监听了 8080 端口：Go 默认情况会监听你本机所有的 IP 上面的 8080 端口。
	beego.Run()
}

// 停止服务的话，请按 Ctrl+c。
