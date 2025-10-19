// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PaymentGateway is the golang structure for table payment_gateway.
type PaymentGateway struct {
	Id            uint        `json:"id"            orm:"id"             description:""`         //
	PaymentId     int         `json:"paymentId"     orm:"payment_id"     description:""`         //
	Gateway       int         `json:"gateway"       orm:"gateway"        description:""`         //
	Status        int         `json:"status"        orm:"status"         description:""`         //
	MobileSupport int         `json:"mobileSupport" orm:"mobile_support" description:""`         //
	UrlOrder      string      `json:"urlOrder"      orm:"url_order"      description:""`         //
	UrlQuery      string      `json:"urlQuery"      orm:"url_query"      description:""`         //
	IsSelectBank  int         `json:"isSelectBank"  orm:"is_select_bank" description:"是否需要选择银行"` // 是否需要选择银行
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""`         //
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:""`         //
}
