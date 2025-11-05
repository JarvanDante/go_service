// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SignLog is the golang structure for table sign_log.
type SignLog struct {
	Id        uint        `json:"id"         orm:"id"         description:""`     //
	SiteId    int         `json:"site_id"    orm:"site_id"    description:"站点ID"` // 站点ID
	UserId    int         `json:"user_id"    orm:"user_id"    description:"会员ID"` // 会员ID
	SignDate  *gtime.Time `json:"sign_date"  orm:"sign_date"  description:"签到日期"` // 签到日期
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`     //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`     //
}
