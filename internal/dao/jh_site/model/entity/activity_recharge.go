// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityRecharge is the golang structure for table activity_recharge.
type ActivityRecharge struct {
	Id            uint        `json:"id"             orm:"id"             description:""`                           //
	SiteId        int         `json:"site_id"        orm:"site_id"        description:"站点ID"`                       // 站点ID
	ActivityId    int         `json:"activity_id"    orm:"activity_id"    description:"活动ID"`                       // 活动ID
	PcContent     string      `json:"pc_content"     orm:"pc_content"     description:"pc端活动内容"`                    // pc端活动内容
	MobileContent string      `json:"mobile_content" orm:"mobile_content" description:"手机端活动内容"`                    // 手机端活动内容
	PcLink        string      `json:"pc_link"        orm:"pc_link"        description:"pc端活动内容链接"`                  // pc端活动内容链接
	MobileLink    string      `json:"mobile_link"    orm:"mobile_link"    description:"手机端活动内容链接"`                  // 手机端活动内容链接
	PcType        int         `json:"pc_type"        orm:"pc_type"        description:"内容类型: 1=pc端编辑框内容 2=pc端内容链接"` // 内容类型: 1=pc端编辑框内容 2=pc端内容链接
	MobileType    int         `json:"mobile_type"    orm:"mobile_type"    description:"内容类型: 1=手机端编辑框内容 2=手机端内容链接"` // 内容类型: 1=手机端编辑框内容 2=手机端内容链接
	CreatedAt     *gtime.Time `json:"created_at"     orm:"created_at"     description:""`                           //
	UpdatedAt     *gtime.Time `json:"updated_at"     orm:"updated_at"     description:""`                           //
}
