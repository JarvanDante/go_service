// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SignInRecords is the golang structure for table sign_in_records.
type SignInRecords struct {
	Id           int64       `json:"id"           orm:"id"            description:"主键ID"`           // 主键ID
	SiteId       int         `json:"siteId"       orm:"site_id"       description:"网站id"`           // 网站id
	UserId       int64       `json:"userId"       orm:"user_id"       description:"用户ID"`           // 用户ID
	SignInDate   *gtime.Time `json:"signInDate"   orm:"sign_in_date"  description:"签到日期（每天最多签到一次）"` // 签到日期（每天最多签到一次）
	RewardAmount float64     `json:"rewardAmount" orm:"reward_amount" description:"签到获得的奖励金额"`      // 签到获得的奖励金额
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"签到时间"`           // 签到时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"updated_at"`     // updated_at
}
