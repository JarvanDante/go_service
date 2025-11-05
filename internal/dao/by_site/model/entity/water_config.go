// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WaterConfig is the golang structure for table water_config.
type WaterConfig struct {
	Id         int         `json:"id"         orm:"id"          description:"ID"`                            // ID
	Name       string      `json:"name"       orm:"name"        description:"配置项名称"`                         // 配置项名称
	Style      int         `json:"style"      orm:"style"       description:"主类型:充值/加款/活动/返水"`               // 主类型:充值/加款/活动/返水
	Type       int         `json:"type"       orm:"type"        description:"子类型:签到/升级/注册送/首充送/大转盘/周礼金/月礼金"` // 子类型:签到/升级/注册送/首充送/大转盘/周礼金/月礼金
	WaterTimes int         `json:"waterTimes" orm:"water_times" description:"倍水数"`                           // 倍水数
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`                          // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`                          // 更新时间
}
