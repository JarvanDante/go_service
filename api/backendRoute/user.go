package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type UsersReq struct {
	g.Meta    `path:"/users" method:"get" summary:"获取会员列表"`
	Grade     *int    `json:"grade" dc:"会员等级"`
	Level     *int    `json:"level" dc:"会员层级"`
	Status    *int    `json:"status" dc:"状态"`
	Username  *string `json:"username" dc:"用户名"`
	Realname  *string `json:"realname" dc:"真实姓名"`
	Agent     *string `json:"agent" dc:"代理用户名"`
	Mobile    *string `json:"mobile" dc:"手机号"`
	CardNo    *string `json:"card_no" dc:"银行卡号"`
	Domain    *string `json:"domain" dc:"域名"`
	StartDate *string `json:"start_date" dc:"开始日期"`
	EndDate   *string `json:"end_date" dc:"结束日期"`
	Charge    *int    `json:"charge" dc:"是否首存"`
	SortField *string `json:"sort_field" dc:"排序字段"`
	SortRule  *int    `json:"sort_rule" d:"1" dc:"排序规则 1=升序 2=降序"`
	Page      *int    `json:"page" d:"1" dc:"页数"`
	Num       *int    `json:"num" d:"20" dc:"每页数量"`
}
