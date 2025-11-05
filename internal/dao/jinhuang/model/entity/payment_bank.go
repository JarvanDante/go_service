// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PaymentBank is the golang structure for table payment_bank.
type PaymentBank struct {
	Id        uint        `json:"id"         orm:"id"         description:""`             //
	PaymentId int         `json:"payment_id" orm:"payment_id" description:"支付ID"`         // 支付ID
	BankId    int         `json:"bank_id"    orm:"bank_id"    description:"银行ID"`         // 银行ID
	BankName  string      `json:"bank_name"  orm:"bank_name"  description:"银行名称"`         // 银行名称
	BankCode  string      `json:"bank_code"  orm:"bank_code"  description:"银行标识"`         // 银行标识
	BankValue string      `json:"bank_value" orm:"bank_value" description:"支付接口银行code"`   // 支付接口银行code
	Status    int         `json:"status"     orm:"status"     description:"状态。1=可用；0=禁用"` // 状态。1=可用；0=禁用
	Sort      int         `json:"sort"       orm:"sort"       description:"排序"`           // 排序
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`             //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`             //
}
