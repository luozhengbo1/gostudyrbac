package controllers

import 	(
	"strings"
	"github.com/hello/utils"
	"github.com/hello/models"
	"github.com/hello/enums"
)


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
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["headcssjs"] = "home/login_headcssjs.html"
	c.LayoutSections["footerjs"]  = "home/login_footerjs.html"
	c.setTpl("home/login.html","shared/layout_base.html")
}


func (c *HomeController) Error() {
	c.Data["error"] = c.GetString(":error")
	c.setTpl("home/error.html", "shared/layout_pullbox.html")
}
//登录判断
func (c *HomeController) DoLogin()  {
 username := strings.TrimSpace(c.GetString("UserName"))
 userpwd  := strings.TrimSpace(c.GetString("UserPwd"))
 if len(username)==0 || len(userpwd)==0{
 	c.jsonResult(enums.JRCodeFailed,"用户名或密码不正确","")
 }
 userpwd = utils.String2md5(userpwd)
 user,err :=models.BackendUserOneByUserName(username,userpwd)
 if user!=nil && err==nil{
 	if user.Status == enums.Disabled{
 		c.jsonResult(enums.JRCodeFailed,"用户被禁用","")
 	}
 	// 保存用户到session
 	c.setBackendUser2Session(user.Id)
 	//获取用户信息
 	c.jsonResult(enums.JRCodeSucc,"登录成功","")
 }else{
 	c.jsonResult(enums.JRCodeFailed,"用户民或密码错误","")
 }
}