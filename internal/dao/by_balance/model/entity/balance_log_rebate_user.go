// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogRebateUser is the golang structure for table balance_log_rebate_user.
type BalanceLogRebateUser struct {
	SiteId          int         `json:"siteId"          orm:"site_id"           description:"站点ID"`           // 站点ID
	ChannelId       int         `json:"channelId"       orm:"channel_id"        description:"渠道id"`           // 渠道id
	Id              uint64      `json:"id"              orm:"id"                description:""`               //
	UserId          int         `json:"userId"          orm:"user_id"           description:"用户ID"`           // 用户ID
	Username        string      `json:"username"        orm:"username"          description:"会员用户名"`          // 会员用户名
	TradeNo         string      `json:"tradeNo"         orm:"trade_no"          description:"流水号"`            // 流水号
	RebateStartTime *gtime.Time `json:"rebateStartTime" orm:"rebate_start_time" description:"洗码开始时间"`         // 洗码开始时间
	RebateEndTime   *gtime.Time `json:"rebateEndTime"   orm:"rebate_end_time"   description:"洗码结束时间"`         // 洗码结束时间
	ValidBetAmount  float64     `json:"validBetAmount"  orm:"valid_bet_amount"  description:"有效投注总金额"`        // 有效投注总金额
	Money           float64     `json:"money"           orm:"money"             description:"返水金额"`           // 返水金额
	Fee             float64     `json:"fee"             orm:"fee"               description:"手续费"`            // 手续费
	Status          int         `json:"status"          orm:"status"            description:"状态。0=未确认；1=已确认"` // 状态。0=未确认；1=已确认
	AdminId         int         `json:"adminId"         orm:"admin_id"          description:"后台管理员ID"`        // 后台管理员ID
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"        description:""`               //
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"        description:""`               //
}
