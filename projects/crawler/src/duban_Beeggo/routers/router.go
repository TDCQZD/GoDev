package routers

import (
	"GoDev/projects/crawler/src/duban_Beeggo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/crawlmovie", &controllers.CrawlMovieController{}, "*:CrawlMovie")

}
