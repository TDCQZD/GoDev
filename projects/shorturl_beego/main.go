package main

import (
	"GoDev/projects/shorturl_beego/controllers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/shorten", &controllers.ShortController{})
	beego.Router("/v1/expand", &controllers.ExpandController{})

	beego.Run()
}
