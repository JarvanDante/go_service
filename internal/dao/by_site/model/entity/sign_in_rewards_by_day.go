// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SignInRewardsByDay is the golang structure for table sign_in_rewards_by_day.
type SignInRewardsByDay struct {
	Id           int64       `json:"id"           orm:"id"            description:"主键ID"`           // 主键ID
	SiteId       int         `json:"siteId"       orm:"site_id"       description:"站点id"`           // 站点id
	DayNumber    int         `json:"dayNumber"    orm:"day_number"    description:"连续签到的第几天（从1开始）"` // 连续签到的第几天（从1开始）
	RewardAmount float64     `json:"rewardAmount" orm:"reward_amount" description:"第N天签到的奖励金额"`     // 第N天签到的奖励金额
	Remarks      string      `json:"remarks"      orm:"remarks"       description:"备注说明（可选）"`       // 备注说明（可选）
	PaymentLimit float64     `json:"paymentLimit" orm:"payment_limit" description:"存款要求"`           // 存款要求
	BetLimit     float64     `json:"betLimit"     orm:"bet_limit"     description:"投注要求"`           // 投注要求
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`           // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`           // 更新时间
}
