// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityModule is the golang structure for table activity_module.
type ActivityModule struct {
	Id         uint        `json:"id"          orm:"id"          description:""`                                    //
	SiteId     int         `json:"site_id"     orm:"site_id"     description:"站点ID"`                                // 站点ID
	ActivityId int         `json:"activity_id" orm:"activity_id" description:"活动ID"`                                // 活动ID
	ModuleType int         `json:"module_type" orm:"module_type" description:"活动模块类型。1=充值；2=大转盘；3=开宝箱；4=砸金蛋；5=抢红包"` // 活动模块类型。1=充值；2=大转盘；3=开宝箱；4=砸金蛋；5=抢红包
	ModuleId   int         `json:"module_id"   orm:"module_id"   description:"活动模块ID"`                              // 活动模块ID
	ModuleName string      `json:"module_name" orm:"module_name" description:"活动模块名称"`                              // 活动模块名称
	StartTime  *gtime.Time `json:"start_time"  orm:"start_time"  description:"开始时间"`                                // 开始时间
	EndTime    *gtime.Time `json:"end_time"    orm:"end_time"    description:"结束时间"`                                // 结束时间
	Status     int         `json:"status"      orm:"status"      description:"活动状态。1=开启；0=关闭"`                      // 活动状态。1=开启；0=关闭
	CreatedAt  *gtime.Time `json:"created_at"  orm:"created_at"  description:""`                                    //
	UpdatedAt  *gtime.Time `json:"updated_at"  orm:"updated_at"  description:""`                                    //
}
