// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserSignInStatus is the golang structure of table user_sign_in_status for DAO operations like Where/Data.
type UserSignInStatus struct {
	g.Meta         `orm:"table:user_sign_in_status, do:true"`
	Id             any         // 用户ID，主键，一对一对应用户表
	UserId         any         // uid
	SiteId         any         // 网站id
	LastSignInDate *gtime.Time // 最后一次签到日期
	ContinuousDays any         // 连续签到天数（断签后重置）
	TotalDays      any         // 累计签到总天数（不断签）
	CreatedAt      *gtime.Time // created_at
	UpdatedAt      *gtime.Time // 记录更新时间
}
