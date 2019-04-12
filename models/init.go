package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)
//init 初始化
func init()  {
	orm.RegisterModel(new(BackendUser))
}
// 下面是统一的表名管理
func TableName(name string) string  {
	//获取表前缀
	prefix :=beego.AppConfig.String("db_dt_prefix")
	return prefix+name
}

//BackendUserTBName 获取BackendUser 对应的表名称
func BackendUserTBName() string  {
	return TableName("backend_user")
}
