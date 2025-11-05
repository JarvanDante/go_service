// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBank is the golang structure for table user_bank.
type UserBank struct {
	Id          uint        `json:"id"           orm:"id"           description:""`                  //
	SiteId      int         `json:"site_id"      orm:"site_id"      description:"站点ID"`              // 站点ID
	UserId      int         `json:"user_id"      orm:"user_id"      description:"会员ID"`              // 会员ID
	BankName    string      `json:"bank_name"    orm:"bank_name"    description:"银行名称"`              // 银行名称
	CardAccount string      `json:"card_account" orm:"card_account" description:"银行账户名.这个应该与会员实名一致"` // 银行账户名.这个应该与会员实名一致
	CardNo      string      `json:"card_no"      orm:"card_no"      description:"银行卡号"`              // 银行卡号
	DepositBank string      `json:"deposit_bank" orm:"deposit_bank" description:"开户行"`               // 开户行
	IsDefault   int         `json:"is_default"   orm:"is_default"   description:"默认地址。1=默认，0=非默认"`   // 默认地址。1=默认，0=非默认
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   description:""`                  //
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   description:""`                  //
}
