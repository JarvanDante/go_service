// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogGame is the golang structure for table balance_log_game.
type BalanceLogGame struct {
	Id                uint        `json:"id"                   orm:"id"                   description:""`                    //
	SiteId            int         `json:"site_id"              orm:"site_id"              description:"站点ID"`                // 站点ID
	UserId            int         `json:"user_id"              orm:"user_id"              description:"会员ID"`                // 会员ID
	GameId            int         `json:"game_id"              orm:"game_id"              description:"游戏ID"`                // 游戏ID
	InOrOut           int         `json:"in_or_out"            orm:"in_or_out"            description:"转入还是转出；0=转出；1=转入"`    // 转入还是转出；0=转出；1=转入
	TradeNo           string      `json:"trade_no"             orm:"trade_no"             description:""`                    //
	Money             float64     `json:"money"                orm:"money"                description:"操作金额"`                // 操作金额
	Status            int         `json:"status"               orm:"status"               description:"状态. 1=待确认，2=成功，3=失败"` // 状态. 1=待确认，2=成功，3=失败
	Remark            string      `json:"remark"               orm:"remark"               description:"备注"`                  // 备注
	ThirdPartyTradeNo string      `json:"third_party_trade_no" orm:"third_party_trade_no" description:"第三方订单号"`              // 第三方订单号
	CreatedAt         *gtime.Time `json:"created_at"           orm:"created_at"           description:""`                    //
	UpdatedAt         *gtime.Time `json:"updated_at"           orm:"updated_at"           description:""`                    //
}
