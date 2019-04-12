package controllers

import "github.com/astaxie/beego"
type LoginController struct {
	//继承框架核心控制器
	beego.Controller
}
// 重写方法get
func (login *LoginController) Get(){
	login.TplName = "login.tpl"
}
