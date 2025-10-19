// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserMgAccount is the golang structure for table user_mg_account.
type UserMgAccount struct {
	UserId      int         `json:"userId"      orm:"user_id"       description:"用户id"`     // 用户id
	MgAccountId int         `json:"mgAccountId" orm:"mg_account_id" description:"mg平台账号id"` // mg平台账号id
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:"创建时间"`     // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:"更新时间"`     // 更新时间
}
