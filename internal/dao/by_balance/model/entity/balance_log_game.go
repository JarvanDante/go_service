// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogGame is the golang structure for table balance_log_game.
type BalanceLogGame struct {
	Id                uint        `json:"id"                orm:"id"                   description:""`                    //
	SiteId            int         `json:"siteId"            orm:"site_id"              description:"站点ID"`                // 站点ID
	ChannelId         int         `json:"channelId"         orm:"channel_id"           description:"渠道id"`                // 渠道id
	UserId            int         `json:"userId"            orm:"user_id"              description:"会员ID"`                // 会员ID
	GameId            int         `json:"gameId"            orm:"game_id"              description:"游戏ID"`                // 游戏ID
	InOrOut           int         `json:"inOrOut"           orm:"in_or_out"            description:"转入还是转出；0=转出；1=转入"`    // 转入还是转出；0=转出；1=转入
	TradeNo           string      `json:"tradeNo"           orm:"trade_no"             description:""`                    //
	Money             float64     `json:"money"             orm:"money"                description:"操作金额"`                // 操作金额
	Status            int         `json:"status"            orm:"status"               description:"状态. 1=待确认，2=成功，3=失败"` // 状态. 1=待确认，2=成功，3=失败
	Currency          string      `json:"currency"          orm:"currency"             description:"币类型"`                 // 币类型
	Remark            string      `json:"remark"            orm:"remark"               description:"备注"`                  // 备注
	ThirdPartyTradeNo string      `json:"thirdPartyTradeNo" orm:"third_party_trade_no" description:"第三方订单号"`              // 第三方订单号
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"           description:""`                    //
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"           description:""`                    //
}
