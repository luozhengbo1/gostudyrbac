package models


import "github.com/hello/enums"
// 用于返回ajax 请求的基类
type JsonResult struct {
	Code enums.JsonResultCode 	`json:"code"`
	Msg string 					`json:"msg"`
	Obj interface{}				`json:"obj"`
}
// 用于查询的基类
type BaseQueryParam struct {
	Sort string  		`json:"sort"`
	Order string 		`json:"order"`
	Offset int64		`json:"offset"`
	Limit	int			`json:"limit"`
}