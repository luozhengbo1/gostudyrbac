package models

import "github.com/astaxie/beego/orm"
//角色查询结构体 y用于搜索的类
type RoleQueryParam struct {
	BaseQueryParam
	NameLike string
}

//角色表
type Role struct {
	Id     				int  `form:"Id"`
	Name				string `form:"Name"`
	Seq					int
	RoleResourceRel		[]*RoleResourceRel  `orm:"reverse(many)" json:"-"` //设置一对多的反向关系
	RoleBackendUserRel 	[]*RoleBackendUserRel `orm:"reverse(many)" json:"-"` //设置一对多的反向关系
}
//RolePageList  获取分页数据
func RolePageList(params *RoleQueryParam) ([]*Role,int64){
	query :=orm.NewOrm().QueryTable(RoleTBName())
	data  := make([]*Role,0)
	//默认排序
	sortorder :="Id"
	switch params.Sort {

	case "Id":
		sortorder="Id"
	case "Seq":
		sortorder="Seq"
	}
	if params.Order =="desc"{
		sortorder = "-" + sortorder
	}
	query = query.Filter("name__istartswith",params.NameLike)
	total , _:= query.Count()
	query.OrderBy(sortorder).Limit(params.Limit,params.Offset).All(&data)
	return data,total
}
