package controllers

import 	"github.com/astaxie/beego"


type HomeController struct {
	BaseController
}

func (c *HomeController) Index()  {
	//判断是否登录
	c.checkLogin()
	c.setTpl()
}
//登录
func (c *HomeController) Login()  {
	
}

//判断是否登录
func (c *HomeController) checkLogin()  {

}
