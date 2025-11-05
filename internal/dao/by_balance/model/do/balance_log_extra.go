// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogExtra is the golang structure of table balance_log_extra for DAO operations like Where/Data.
type BalanceLogExtra struct {
	g.Meta    `orm:"table:balance_log_extra, do:true"`
	Id        any         // ID
	SiteId    any         // 应用id
	ChannelId any         // 渠道id
	UserId    any         // 用户id
	Username  any         // 用户名
	Money     any         // 金额
	AdminId   any         // 后台管理员id
	Type      any         // 类型1加2减
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
