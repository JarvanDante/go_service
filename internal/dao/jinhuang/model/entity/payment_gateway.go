// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PaymentGateway is the golang structure for table payment_gateway.
type PaymentGateway struct {
	Id            uint        `json:"id"             orm:"id"             description:""`         //
	PaymentId     int         `json:"payment_id"     orm:"payment_id"     description:""`         //
	Gateway       int         `json:"gateway"        orm:"gateway"        description:""`         //
	Status        int         `json:"status"         orm:"status"         description:""`         //
	MobileSupport int         `json:"mobile_support" orm:"mobile_support" description:""`         //
	UrlOrder      string      `json:"url_order"      orm:"url_order"      description:""`         //
	UrlQuery      string      `json:"url_query"      orm:"url_query"      description:""`         //
	IsSelectBank  int         `json:"is_select_bank" orm:"is_select_bank" description:"是否需要选择银行"` // 是否需要选择银行
	CreatedAt     *gtime.Time `json:"created_at"     orm:"created_at"     description:""`         //
	UpdatedAt     *gtime.Time `json:"updated_at"     orm:"updated_at"     description:""`         //
}
