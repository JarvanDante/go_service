package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type RebateSettingReq struct {
	g.Meta `path:"/rebate-setting" method:"get" summary:"获取返水设置"`
}

type CreateRebateRuleReq struct {
	g.Meta `path:"/create-rebate-rule" method:"post" summary:"添加返水规则"`
	Name   string `json:"name" v:"required|length:2,50#请输入规则名称|规则名称长度为2-50个字符"`
}

type UpdateRebateRuleReq struct {
	g.Meta `path:"/update-rebate-rule" method:"post" summary:"编辑返水规则"`
	Id     int    `json:"id" v:"required|min:1#ID不能为空|ID必须大于0"`
	Name   string `json:"name" v:"required|length:2,50#请输入规则名称|规则名称长度为2-50个字符"`
}

type DeleteRebateRuleReq struct {
	g.Meta `path:"/delete-rebate-rule" method:"post" summary:"删除返水规则"`
	Id     int `json:"id" v:"required|min:1#ID不能为空|ID必须大于0"`
}

type CreateRebateOptionReq struct {
	g.Meta `path:"/create-rebate-option" method:"post" summary:"添加返水规则条件"`
	Id     int `json:"id" v:"required|min:1#ID不能为空|ID必须大于0"`
}

type UpdateRebateOptionReq struct {
	g.Meta  `path:"/update-rebate-option" method:"post" summary:"编辑返水规则条件"`
	Id      int     `json:"id" v:"required|min:1#ID不能为空|ID必须大于0"`
	Min     float64 `json:"min" v:"required#请输入单日有效投注范围最小金额"`
	Max     float64 `json:"max" v:"required#请输入单日有效投注范围最大金额"`
	Content string  `json:"content"`
}

type DeleteRebateOptionReq struct {
	g.Meta `path:"/delete-rebate-option" method:"post" summary:"删除返水规则条件"`
	Id     int `json:"id" v:"required|min:1#ID不能为空|ID必须大于0"`
}
