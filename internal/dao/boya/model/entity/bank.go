// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Bank is the golang structure for table bank.
type Bank struct {
	Id        uint        `json:"id"        orm:"id"         description:""` //
	Type      int         `json:"type"      orm:"type"       description:""` //
	Name      string      `json:"name"      orm:"name"       description:""` //
	Code      string      `json:"code"      orm:"code"       description:""` //
	Status    int         `json:"status"    orm:"status"     description:""` //
	Sort      int         `json:"sort"      orm:"sort"       description:""` //
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""` //
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""` //
}
