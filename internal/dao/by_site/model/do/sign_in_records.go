// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SignInRecords is the golang structure of table sign_in_records for DAO operations like Where/Data.
type SignInRecords struct {
	g.Meta       `orm:"table:sign_in_records, do:true"`
	Id           any         // 主键ID
	SiteId       any         // 网站id
	UserId       any         // 用户ID
	SignInDate   *gtime.Time // 签到日期（每天最多签到一次）
	RewardAmount any         // 签到获得的奖励金额
	CreatedAt    *gtime.Time // 签到时间
	UpdatedAt    *gtime.Time // updated_at
}
