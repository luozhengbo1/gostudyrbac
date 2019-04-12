package models
import (
	"github.com/astaxie/beego/orm"
)
//BackendUser 实体类
type BackendUser struct {
	Id 			int
	RealName	string `orm:"size(32)"`
	UserName	string `orm:"size(32)"`
	UserPwd		string `json:"-"`
	IsSuper		bool
	Status		int
	Mobile		string		`orm:"size(16)"`
	Email		string		`orm:"size(256)"`
	Avatar		string		`orm:"size(256)"`
	RoleIds		[]int 		`orm:"-" form:"RolsIds"`
	RoleBackendUserRel []*RoleBackendUserRel 	`orm:"reverse(many)"` //设置一对多的反向关系
	ResourceUrlForList []string 	`orm:"-"`
}
//tablename 设置backenduser 表名
func (a *BackendUser) TableName() string  {
	return BackendUserTBName()
}
// BackendUserQueryParam 用于查询的类
type BackendUserQueryParam struct {
		BaseQueryParam
		UserNameLike 	string //模糊查询
		RealNameLike	string // 模糊查询
		Mobile			string // 模糊查询
		SearchStatus	string //为空不查询 ，有值精准查询
}
//获取分页数据
func BackendUserPageList(params *BackendUserQueryParam) ([]*BackendUser,int64){

	query :=orm.NewOrm().QueryTable(BackendUserTBName())

	return data ,total
}

