// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PaymentGateway is the golang structure of table payment_gateway for DAO operations like Where/Data.
type PaymentGateway struct {
	g.Meta        `orm:"table:payment_gateway, do:true"`
	Id            any         //
	PaymentId     any         //
	Gateway       any         //
	Status        any         //
	MobileSupport any         //
	UrlOrder      any         //
	UrlQuery      any         //
	IsSelectBank  any         // 是否需要选择银行
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
}
