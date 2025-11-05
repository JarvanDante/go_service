// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLevelTransfer is the golang structure for table user_level_transfer.
type UserLevelTransfer struct {
	Id                uint        `json:"id"                  orm:"id"                  description:""`                       //
	SiteId            int         `json:"site_id"             orm:"site_id"             description:"站点ID"`                   // 站点ID
	UserLevelId       int         `json:"user_level_id"       orm:"user_level_id"       description:"会员层级ID"`                 // 会员层级ID
	BankType          int         `json:"bank_type"           orm:"bank_type"           description:"转账类型。1=网银转账；2=微信；3=支付宝"` // 转账类型。1=网银转账；2=微信；3=支付宝
	TransferAccountId int         `json:"transfer_account_id" orm:"transfer_account_id" description:"转账接口ID"`                 // 转账接口ID
	CreatedAt         *gtime.Time `json:"created_at"          orm:"created_at"          description:""`                       //
	UpdatedAt         *gtime.Time `json:"updated_at"          orm:"updated_at"          description:""`                       //
}
