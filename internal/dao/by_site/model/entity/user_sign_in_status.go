// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserSignInStatus is the golang structure for table user_sign_in_status.
type UserSignInStatus struct {
	Id             int64       `json:"id"             orm:"id"                description:"用户ID，主键，一对一对应用户表"` // 用户ID，主键，一对一对应用户表
	UserId         int         `json:"userId"         orm:"user_id"           description:"uid"`              // uid
	SiteId         int         `json:"siteId"         orm:"site_id"           description:"网站id"`             // 网站id
	LastSignInDate *gtime.Time `json:"lastSignInDate" orm:"last_sign_in_date" description:"最后一次签到日期"`         // 最后一次签到日期
	ContinuousDays int         `json:"continuousDays" orm:"continuous_days"   description:"连续签到天数（断签后重置）"`    // 连续签到天数（断签后重置）
	TotalDays      int         `json:"totalDays"      orm:"total_days"        description:"累计签到总天数（不断签）"`     // 累计签到总天数（不断签）
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"        description:"created_at"`       // created_at
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"        description:"记录更新时间"`           // 记录更新时间
}
