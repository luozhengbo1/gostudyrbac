package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hello/models"
	)


type BaseController struct {
	beego.Controller
	controllerName  string //当前控制器名
	actionName 		string // 当前action 名称
	curUser			models.BackendUser //当前用户信息
}

//检查用户是否登录
func (c *BaseController) checkLogin()  {
	if c.curUser.Id == 0 {
		//登录页面地址
		urlstr :=c.URLFor("HomeController.Login")+"?url="

		//登录成功之后返回的地址为当前
		returnURL :=c.Ctx.Request.URL.Path

		//如果ajax请求则返回相应的错误码和跳转的地址
		if c.Ctx.Input.IsAjax() {
			//由于是ajax 请求， 因此地址是header 里的Referer
			returnURL := c.Ctx.Input.Refer()
			c.jsonResult(enums.JRCode302,"请登录",urlstr+returnURL)
		}
		c.Redirect(urlstr+returnURL,302)
		c.StopRun()

	}
}