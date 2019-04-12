package controllers

import (
	"strings"
	"github.com/astaxie/beego"
	"github.com/hello/models"
	"github.com/hello/enums"
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
func (c *BaseController) jsonResult(code enums.JsonResultCode,msg string,obj interface{}){
	r := &models.JsonResult{code,msg,obj}
	c.Data["json"] = r
	c.ServeJSON()
	c.StopRun()
}

// 设置模板
func (c *BaseController) setTpl(template ...string){
	var tplName string
	layout :="shared/layout_page.html"
	switch  {
	case len(template) == 1:
		tplName =template[0]
	case len(template) == 2:
		tplName =template[0]
		layout = template[1]
	default:
		// 不要Controller 这个10字符
		ctrlName := strings.ToLower(c.controllerName[0:len(c.controllerName)-10])
		actionName := strings.ToLower(c.actionName)
		tplName = ctrlName+"/"+actionName +".html"
	}
	c.Layout = layout
	c.TplName = tplName
}

//setBackendUser2Session 获取用户信息 *包括资源(UrlFor) 保存至Session
func (c *BaseController) setBackendUser2Session(userId int) error{
	m,err :=models.BackendUserOne(userId)
	if err!=nil{
		return err
	}
	//获取 这个用户能获取到的所有资源列表
	resourceList :=models.ResourceTreeGridByUserId(userId,1000)
	for _,item :=range resourceList{
		m.ResourceUrlForList = append(m.ResourceUrlForList,strings.TrimSpace(item.UrlFor))
	}
	c.SetSession("backenduser",*m)
	return nil
}