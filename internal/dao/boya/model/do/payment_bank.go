// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PaymentBank is the golang structure of table payment_bank for DAO operations like Where/Data.
type PaymentBank struct {
	g.Meta    `orm:"table:payment_bank, do:true"`
	Id        any         //
	PaymentId any         // 支付ID
	BankId    any         // 银行ID
	BankName  any         // 银行名称
	BankCode  any         // 银行标识
	BankValue any         // 支付接口银行code
	Status    any         // 状态。1=可用；0=禁用
	Sort      any         // 排序
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
