// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogPayment is the golang structure for table balance_log_payment.
type BalanceLogPayment struct {
	Id                 uint64      `json:"id"                   orm:"id"                   description:""`                  //
	SiteId             int         `json:"site_id"              orm:"site_id"              description:"站点ID"`              // 站点ID
	UserId             int         `json:"user_id"              orm:"user_id"              description:"用户ID"`              // 用户ID
	Username           string      `json:"username"             orm:"username"             description:"会员用户名"`             // 会员用户名
	ActivityRechargeId int         `json:"activity_recharge_id" orm:"activity_recharge_id" description:"充值活动ID"`            // 充值活动ID
	Gateway            int         `json:"gateway"              orm:"gateway"              description:"网关类型"`              // 网关类型
	PaymentId          int         `json:"payment_id"           orm:"payment_id"           description:"支付ID"`              // 支付ID
	PaymentAccountId   int         `json:"payment_account_id"   orm:"payment_account_id"   description:"支付接口账号ID"`          // 支付接口账号ID
	BankValue          string      `json:"bank_value"           orm:"bank_value"           description:"银行代码"`              // 银行代码
	TradeNo            string      `json:"trade_no"             orm:"trade_no"             description:"流水号"`               // 流水号
	Money              float64     `json:"money"                orm:"money"                description:"充值金额"`              // 充值金额
	Fee                float64     `json:"fee"                  orm:"fee"                  description:"手续费"`               // 手续费
	Status             int         `json:"status"               orm:"status"               description:"状态。1=未付；2=成功；3=失败"` // 状态。1=未付；2=成功；3=失败
	AdminId            int         `json:"admin_id"             orm:"admin_id"             description:"后台管理员ID。默认为0"`      // 后台管理员ID。默认为0
	Remark             string      `json:"remark"               orm:"remark"               description:"备注"`                // 备注
	CreatedAt          *gtime.Time `json:"created_at"           orm:"created_at"           description:""`                  //
	UpdatedAt          *gtime.Time `json:"updated_at"           orm:"updated_at"           description:""`                  //
}
