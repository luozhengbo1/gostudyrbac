package routers

import (
	"github.com/hello/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/api/list", &controllers.MainController{}, "*:List")
    beego.Router("/", &controllers.MainController{})

	//后台 用户中心
	beego.Router("/home/index",&controllers.HomeController{},"*:Index")


}
