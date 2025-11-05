// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notice is the golang structure for table notice.
type Notice struct {
	Id          uint        `json:"id"           orm:"id"           description:""` //
	SiteId      int         `json:"site_id"      orm:"site_id"      description:""` //
	Type        int         `json:"type"         orm:"type"         description:""` //
	Title       string      `json:"title"        orm:"title"        description:""` //
	Url         string      `json:"url"          orm:"url"          description:""` //
	StartTime   *gtime.Time `json:"start_time"   orm:"start_time"   description:""` //
	ExpiredTime *gtime.Time `json:"expired_time" orm:"expired_time" description:""` //
	Sort        int         `json:"sort"         orm:"sort"         description:""` //
	Status      int         `json:"status"       orm:"status"       description:""` //
	Platform    int         `json:"platform"     orm:"platform"     description:""` //
	Content     string      `json:"content"      orm:"content"      description:""` //
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   description:""` //
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   description:""` //
}
