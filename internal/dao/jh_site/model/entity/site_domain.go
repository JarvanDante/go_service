// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteDomain is the golang structure for table site_domain.
type SiteDomain struct {
	Id        uint        `json:"id"         orm:"id"         description:""` //
	SiteId    int         `json:"site_id"    orm:"site_id"    description:""` //
	Domain    string      `json:"domain"     orm:"domain"     description:""` //
	Status    int         `json:"status"     orm:"status"     description:""` //
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""` //
}
