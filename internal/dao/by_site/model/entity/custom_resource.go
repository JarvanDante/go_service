// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CustomResource is the golang structure for table custom_resource.
type CustomResource struct {
	Id             int         `json:"id"             orm:"id"              description:""`     //
	SiteId         int         `json:"siteId"         orm:"site_id"         description:"站点id"` // 站点id
	UserId         int         `json:"userId"         orm:"user_id"         description:"用户ID"` // 用户ID
	Ip             string      `json:"ip"             orm:"ip"              description:"ip"`   // ip
	ResourceDomain string      `json:"resourceDomain" orm:"resource_domain" description:"来源域名"` // 来源域名
	ResourceLink   string      `json:"resourceLink"   orm:"resource_link"   description:"原来链接"` // 原来链接
	Remark         string      `json:"remark"         orm:"remark"          description:"备注"`   // 备注
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""`     //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""`     //
}
