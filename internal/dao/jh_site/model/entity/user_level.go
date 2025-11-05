// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLevel is the golang structure for table user_level.
type UserLevel struct {
	Id                 uint        `json:"id"                   orm:"id"                   description:""`                //
	SiteId             int         `json:"site_id"              orm:"site_id"              description:"站点ID"`            // 站点ID
	Name               string      `json:"name"                 orm:"name"                 description:"层级名称"`            // 层级名称
	IsRebate           int         `json:"is_rebate"            orm:"is_rebate"            description:"是否返水.1=返水；0=不返水"` // 是否返水.1=返水；0=不返水
	RebateRuleId       int         `json:"rebate_rule_id"       orm:"rebate_rule_id"       description:"返水规则ID"`          // 返水规则ID
	DailyWithdrawTimes int         `json:"daily_withdraw_times" orm:"daily_withdraw_times" description:"单日提款次数上限"`        // 单日提款次数上限
	LoginUrl           string      `json:"login_url"            orm:"login_url"            description:"专用登录网址"`          // 专用登录网址
	Status             int         `json:"status"               orm:"status"               description:"状态.1=可用；0=禁用"`    // 状态.1=可用；0=禁用
	CreatedAt          *gtime.Time `json:"created_at"           orm:"created_at"           description:""`                //
	UpdatedAt          *gtime.Time `json:"updated_at"           orm:"updated_at"           description:""`                //
}
