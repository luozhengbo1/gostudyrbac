package routers

import (
	"github.com/hello/controllers"
	"github.com/astaxie/beego"
)

func init() {


	//后台 用户中心
	beego.Router("/home/index",&controllers.HomeController{},"*:Index")
	// 访问方式 控制器  * 方式：方法
	beego.Router("/home/login",&controllers.HomeController{},"*:Login")
	beego.Router("/home/dologin", &controllers.HomeController{}, "Post:DoLogin")

	beego.Router("/", &controllers.HomeController{}, "*:Index")
}
