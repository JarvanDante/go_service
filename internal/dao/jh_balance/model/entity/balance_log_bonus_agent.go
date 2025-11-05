// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogBonusAgent is the golang structure for table balance_log_bonus_agent.
type BalanceLogBonusAgent struct {
	Id           int         `json:"id"             orm:"id"             description:""`    //
	SiteId       int         `json:"site_id"        orm:"site_id"        description:""`    //
	UserId       int         `json:"user_id"        orm:"user_id"        description:""`    //
	AgentId      int         `json:"agent_id"       orm:"agent_id"       description:""`    //
	AgentLevelId int         `json:"agent_level_id" orm:"agent_level_id" description:""`    //
	TradeNo      string      `json:"trade_no"       orm:"trade_no"       description:"订单号"` // 订单号
	Money        float64     `json:"money"          orm:"money"          description:""`    //
	CreatedAt    *gtime.Time `json:"created_at"     orm:"created_at"     description:""`    //
	UpdatedAt    *gtime.Time `json:"updated_at"     orm:"updated_at"     description:""`    //
}
