// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Channel is the golang structure of table channel for DAO operations like Where/Data.
type Channel struct {
	g.Meta    `orm:"table:channel, do:true"`
	Id        any         // ID
	SiteId    any         // 商户id
	Code      any         // 渠道编码
	Name      any         // 渠道名称
	Type      any         // 渠道类型
	Status    any         // 状态0关1开
	Remark    any         // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time //
}
