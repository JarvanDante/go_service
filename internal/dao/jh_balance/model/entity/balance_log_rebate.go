// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogRebate is the golang structure for table balance_log_rebate.
type BalanceLogRebate struct {
	Id             uint64      `json:"id"               orm:"id"               description:""`               //
	SiteId         int         `json:"site_id"          orm:"site_id"          description:"站点ID"`           // 站点ID
	UserId         int         `json:"user_id"          orm:"user_id"          description:"用户ID"`           // 用户ID
	Username       string      `json:"username"         orm:"username"         description:"会员用户名"`          // 会员用户名
	TradeNo        string      `json:"trade_no"         orm:"trade_no"         description:"流水号"`            // 流水号
	RebateDate     *gtime.Time `json:"rebate_date"      orm:"rebate_date"      description:"返水日期"`           // 返水日期
	ValidBetAmount float64     `json:"valid_bet_amount" orm:"valid_bet_amount" description:"有效投注总金额"`        // 有效投注总金额
	Money          float64     `json:"money"            orm:"money"            description:"返水金额"`           // 返水金额
	Fee            float64     `json:"fee"              orm:"fee"              description:"手续费"`            // 手续费
	Status         int         `json:"status"           orm:"status"           description:"状态。0=未确认；1=已确认"` // 状态。0=未确认；1=已确认
	AdminId        int         `json:"admin_id"         orm:"admin_id"         description:"后台管理员ID"`        // 后台管理员ID
	CreatedAt      *gtime.Time `json:"created_at"       orm:"created_at"       description:""`               //
	UpdatedAt      *gtime.Time `json:"updated_at"       orm:"updated_at"       description:""`               //
}
