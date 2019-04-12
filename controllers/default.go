package controllers

import (
	_ "fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type TestController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"
}
func (this *MainController) List(){
	this.Ctx.WriteString("ctx.WriteString这个函数与模板渲染冲突")
	//this.Ctx.WriteString("hello")
}

func (this *TestController) Get()  {
	//this.Ctx.Input
	//获取参数
	//this.GetString()
	this.Ctx.WriteContext(this.GetInt("id"))
}