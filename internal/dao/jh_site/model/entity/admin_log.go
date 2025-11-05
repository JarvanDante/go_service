// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminLog is the golang structure for table admin_log.
type AdminLog struct {
	Id            uint        `json:"id"             orm:"id"             description:""` //
	SiteId        int         `json:"site_id"        orm:"site_id"        description:""` //
	AdminId       int         `json:"admin_id"       orm:"admin_id"       description:""` //
	AdminUsername string      `json:"admin_username" orm:"admin_username" description:""` //
	Ip            string      `json:"ip"             orm:"ip"             description:""` //
	CreatedAt     *gtime.Time `json:"created_at"     orm:"created_at"     description:""` //
	Remark        string      `json:"remark"         orm:"remark"         description:""` //
}
