package routers

import (
	"GoDev/projects/chat_beego/controllers"


	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.AppController{})
	beego.Router("/join", &controllers.AppController{}, "post:Join")
	// Long polling.
	beego.Router("/lp", &controllers.LongPollingController{}, "get:Join")
	beego.Router("/lp/post", &controllers.LongPollingController{})
	beego.Router("/lp/fetch", &controllers.LongPollingController{}, "get:Fetch")

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")

}
