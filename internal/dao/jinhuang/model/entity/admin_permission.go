// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminPermission is the golang structure for table admin_permission.
type AdminPermission struct {
	Id          uint        `json:"id"           orm:"id"           description:""`                 //
	ParentId    int         `json:"parent_id"    orm:"parent_id"    description:"父级id"`             // 父级id
	Type        int         `json:"type"         orm:"type"         description:"权限类型；1=菜单；2=操作权限"` // 权限类型；1=菜单；2=操作权限
	Name        string      `json:"name"         orm:"name"         description:"权限名称"`             // 权限名称
	BackendUrl  string      `json:"backend_url"  orm:"backend_url"  description:"后端url"`            // 后端url
	FrontendUrl string      `json:"frontend_url" orm:"frontend_url" description:"前端url"`            // 前端url
	Status      int         `json:"status"       orm:"status"       description:"状态。1=可用；0=禁用"`     // 状态。1=可用；0=禁用
	Sort        int         `json:"sort"         orm:"sort"         description:"排序。值越小，越靠前"`       // 排序。值越小，越靠前
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   description:""`                 //
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   description:""`                 //
}
