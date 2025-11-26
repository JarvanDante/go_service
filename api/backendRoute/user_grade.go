package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type UserGradesReq struct {
	g.Meta `path:"/user-grades" method:"get" summary:"获取会员等级列表"`
}

type SaveUserGradesReq struct {
	g.Meta        `path:"/save-user-grades" method:"post" summary:"保存会员等级"`
	Data          string `json:"data" v:"required#数据不能为空"`
	FieldsDisable string `json:"fieldsDisable"`
	AutoProviding string `json:"autoProviding"`
}

type DeleteUserGradeReq struct {
	g.Meta `path:"/delete-user-grades" method:"post" summary:"删除会员等级"`
	Id     int `json:"id" v:"required|min:1#ID不能为空|ID必须大于0"`
}
