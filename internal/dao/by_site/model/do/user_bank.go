// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserBank is the golang structure of table user_bank for DAO operations like Where/Data.
type UserBank struct {
	g.Meta      `orm:"table:user_bank, do:true"`
	Id          any         //
	SiteId      any         // 站点ID
	UserId      any         // 会员ID
	BankName    any         // 银行名称
	CardAccount any         // 银行账户名.这个应该与会员实名一致
	CardNo      any         // 银行卡号
	DepositBank any         // 开户行
	IsDefault   any         // 默认地址。1=默认，0=非默认
	Province    any         // 省
	City        any         // 市
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
