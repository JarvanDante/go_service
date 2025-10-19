// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberDayReport is the golang structure for table member_day_report.
type MemberDayReport struct {
	Id            uint64      `json:"id"            orm:"id"             description:"主键ID"`   // 主键ID
	SiteId        uint64      `json:"siteId"        orm:"site_id"        description:"商户ID"`   // 商户ID
	UserId        uint64      `json:"userId"        orm:"user_id"        description:"会员ID"`   // 会员ID
	Date          *gtime.Time `json:"date"          orm:"date"           description:"统计日期"`   // 统计日期
	DepositAmount float64     `json:"depositAmount" orm:"deposit_amount" description:"当日充值总额"` // 当日充值总额
	BetAmount     float64     `json:"betAmount"     orm:"bet_amount"     description:"当日投注总额"` // 当日投注总额
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"记录创建时间"` // 记录创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"记录更新时间"` // 记录更新时间
}
