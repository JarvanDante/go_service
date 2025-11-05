// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogBonusUpgrade is the golang structure for table balance_log_bonus_upgrade.
type BalanceLogBonusUpgrade struct {
	Id        uint64      `json:"id"         orm:"id"         description:""`       //
	SiteId    int         `json:"site_id"    orm:"site_id"    description:"站点ID"`   // 站点ID
	UserId    int         `json:"user_id"    orm:"user_id"    description:"会员ID"`   // 会员ID
	TradeNo   string      `json:"trade_no"   orm:"trade_no"   description:"流水号"`    // 流水号
	GradeId   int         `json:"grade_id"   orm:"grade_id"   description:"会员等级ID"` // 会员等级ID
	Money     float64     `json:"money"      orm:"money"      description:"金额"`     // 金额
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`       //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`       //
}
