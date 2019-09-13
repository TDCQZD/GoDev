package main

import (
	"github.com/beego/i18n"
	_ "GoDev/projects/chat_beego/routers"
	"github.com/astaxie/beego"
)
const (
	APP_VER = "1.0"
)
func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)

	beego.AddFuncMap("i18n",i18n.Tr)
	beego.Run()
}

