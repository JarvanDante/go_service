package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type UserLevelsReq struct {
	g.Meta `path:"/user-levels" method:"get" summary:"获取会员层级列表"`
}

type CreateUserLevelReq struct {
	g.Meta              `path:"/create-user-level" method:"post" summary:"新增会员层级"`
	Name                string `json:"name" v:"required|length:2,50#层级名称必填|层级名称长度为2-50个字符"`
	IsRebate            int    `json:"is_rebate" v:"required|in:0,1#是否返水必填|是否返水参数错误"`
	RebateRuleId        int    `json:"rebate_rule_id"`
	DailyWithdrawTimes  int    `json:"daily_withdraw_times" v:"required|min:0#单日提款次数上限必填|单日提款次数上限不能小于0"`
	TransferAccountList string `json:"transfer_account_list"`
	PaymentAccountList  string `json:"payment_account_list"`
}

type GetUpdateUserLevelReq struct {
	g.Meta `path:"/update-user-level" method:"get" summary:"获取会员层级信息用于编辑"`
	Id     int `json:"id" v:"required|min:1#ID不能为空|ID必须大于0"`
}

type UpdateUserLevelReq struct {
	g.Meta              `path:"/update-user-level" method:"post" summary:"编辑会员层级"`
	Id                  int    `json:"id" v:"required|min:1#ID不能为空|ID必须大于0"`
	Name                string `json:"name" v:"required|length:2,50#层级名称必填|层级名称长度为2-50个字符"`
	IsRebate            int    `json:"is_rebate" v:"required|in:0,1#是否返水必填|是否返水参数错误"`
	RebateRuleId        int    `json:"rebate_rule_id"`
	DailyWithdrawTimes  int    `json:"daily_withdraw_times" v:"required|min:0#单日提款次数上限必填|单日提款次数上限不能小于0"`
	TransferAccountList string `json:"transfer_account_list"`
	PaymentAccountList  string `json:"payment_account_list"`
}

type DeleteUserLevelReq struct {
	g.Meta `path:"/delete-user-level" method:"post" summary:"删除会员层级"`
	Id     int `json:"id" v:"required|min:1#ID不能为空|ID必须大于0"`
}
