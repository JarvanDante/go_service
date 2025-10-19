// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityWaterConfig is the golang structure for table activity_water_config.
type ActivityWaterConfig struct {
	Id         int         `json:"id"         orm:"id"          description:"ID"`     // ID
	TradeType  int         `json:"tradeType"  orm:"trade_type"  description:"活动交易类型"` // 活动交易类型
	WaterTimes int         `json:"waterTimes" orm:"water_times" description:"流水倍数"`   // 流水倍数
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`   // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`   // 更新时间
}
