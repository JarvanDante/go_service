package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type BasicSettingReq struct {
	g.Meta `path:"/basic-setting" method:"get" summary:"获取站点基本信息"`
}

type UpdateBasicSettingReq struct {
	g.Meta               `path:"/update-basic-setting" method:"post" summary:"修改站点基本信息"`
	AgentUrl             string  `json:"agent_url" v:"required#代理链接地址必填" dc:"代理链接地址"`
	MobileUrl            string  `json:"mobile_url" v:"required#手机域名地址必填" dc:"手机域名地址"`
	AgentRegisterUrl     string  `json:"agent_register_url" v:"required#代理推广地址必填" dc:"代理推广地址"`
	MinWithdraw          float64 `json:"min_withdraw" v:"required#单笔最低提现金额必填" dc:"单笔最低提现金额"`
	MaxWithdraw          float64 `json:"max_withdraw" v:"required#单笔最高提现金额必填" dc:"单笔最高提现金额"`
	RegisterTimeInterval string  `json:"register_time_interval" dc:"同一IP重复注册时间"`
	ServiceUrl           string  `json:"service_url" dc:"在线客服链接地址"`
	MobileLogo           string  `json:"mobile_logo" dc:"手机端LOGO"`
	IsClose              int     `json:"is_close" dc:"关闭网站"`
	SwitchRegister       int     `json:"switch_register" dc:"是否开放注册"`
}
