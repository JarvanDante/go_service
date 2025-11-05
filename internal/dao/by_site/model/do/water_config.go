// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WaterConfig is the golang structure of table water_config for DAO operations like Where/Data.
type WaterConfig struct {
	g.Meta     `orm:"table:water_config, do:true"`
	Id         any         // ID
	Name       any         // 配置项名称
	Style      any         // 主类型:充值/加款/活动/返水
	Type       any         // 子类型:签到/升级/注册送/首充送/大转盘/周礼金/月礼金
	WaterTimes any         // 倍水数
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
