// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Site is the golang structure for table site.
type Site struct {
	Id        uint        `json:"id"         orm:"id"         description:""`             //
	Code      string      `json:"code"       orm:"code"       description:"站点标识"`         // 站点标识
	Name      string      `json:"name"       orm:"name"       description:"站点名称"`         // 站点名称
	Status    int         `json:"status"     orm:"status"     description:"状态。1=正常；0=禁用"` // 状态。1=正常；0=禁用
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`             //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`             //
}
