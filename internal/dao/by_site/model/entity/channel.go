// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Channel is the golang structure for table channel.
type Channel struct {
	Id        int         `json:"id"        orm:"id"         description:"ID"`     // ID
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"商户id"`   // 商户id
	Code      string      `json:"code"      orm:"code"       description:"渠道编码"`   // 渠道编码
	Name      string      `json:"name"      orm:"name"       description:"渠道名称"`   // 渠道名称
	Type      int         `json:"type"      orm:"type"       description:"渠道类型"`   // 渠道类型
	Status    int         `json:"status"    orm:"status"     description:"状态0关1开"` // 状态0关1开
	Remark    string      `json:"remark"    orm:"remark"     description:"备注"`     // 备注
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`   // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`   // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`       //
}
