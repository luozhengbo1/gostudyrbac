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

//  获取角色与用户多对多关系表
func RoleBackendUserTBName() string  {
	return TableName("role_backenduser_rel")
}

//获取角色 表
func RoleTBName() string {
	return TableName("role")
}

// 角色与资源多对多关系表
func RoleResourceRelTBName()  string {
	return TableName("role_resource_rel")
}

// ResourceTBName  获取 Resource 对应的表名称
func ResourceTBName() string  {
	return TableName("resource")
}

// RoleBackendUserRelTBName 角色与用户多对多关系表
func RoleBackendUserRelTBName() string {
	return TableName("role_backenduser_rel")
}
