// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityWaterConfig is the golang structure of table activity_water_config for DAO operations like Where/Data.
type ActivityWaterConfig struct {
	g.Meta     `orm:"table:activity_water_config, do:true"`
	Id         any         // ID
	TradeType  any         // 活动交易类型
	WaterTimes any         // 流水倍数
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
