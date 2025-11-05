// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogBonusTransfer is the golang structure for table balance_log_bonus_transfer.
type BalanceLogBonusTransfer struct {
	Id                 uint64      `json:"id"                   orm:"id"                   description:""`         //
	SiteId             int         `json:"site_id"              orm:"site_id"              description:"站点ID"`     // 站点ID
	UserId             int         `json:"user_id"              orm:"user_id"              description:"会员ID"`     // 会员ID
	TradeNo            string      `json:"trade_no"             orm:"trade_no"             description:"流水号"`      // 流水号
	TransferLogId      int64       `json:"transfer_log_id"      orm:"transfer_log_id"      description:"转账汇款日志ID"` // 转账汇款日志ID
	ActivityRechargeId int         `json:"activity_recharge_id" orm:"activity_recharge_id" description:"充值活动ID"`   // 充值活动ID
	Money              float64     `json:"money"                orm:"money"                description:"金额"`       // 金额
	CreatedAt          *gtime.Time `json:"created_at"           orm:"created_at"           description:""`         //
	UpdatedAt          *gtime.Time `json:"updated_at"           orm:"updated_at"           description:""`         //
}
