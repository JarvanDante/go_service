// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogPayment is the golang structure for table balance_log_payment.
type BalanceLogPayment struct {
	Id                 uint64      `json:"id"                 orm:"id"                   description:""`                  //
	SiteId             int         `json:"siteId"             orm:"site_id"              description:"站点ID"`              // 站点ID
	ChannelId          int         `json:"channelId"          orm:"channel_id"           description:"渠道id"`              // 渠道id
	UserId             int         `json:"userId"             orm:"user_id"              description:"用户ID"`              // 用户ID
	Username           string      `json:"username"           orm:"username"             description:"会员用户名"`             // 会员用户名
	ActivityRechargeId int         `json:"activityRechargeId" orm:"activity_recharge_id" description:"充值活动ID"`            // 充值活动ID
	Gateway            int         `json:"gateway"            orm:"gateway"              description:"网关类型"`              // 网关类型
	PaymentId          int         `json:"paymentId"          orm:"payment_id"           description:"支付ID"`              // 支付ID
	PaymentAccountId   int         `json:"paymentAccountId"   orm:"payment_account_id"   description:"支付接口账号ID"`          // 支付接口账号ID
	BankValue          string      `json:"bankValue"          orm:"bank_value"           description:"银行代码"`              // 银行代码
	TradeNo            string      `json:"tradeNo"            orm:"trade_no"             description:"流水号"`               // 流水号
	Money              float64     `json:"money"              orm:"money"                description:"充值金额"`              // 充值金额
	Fee                float64     `json:"fee"                orm:"fee"                  description:"手续费"`               // 手续费
	Status             int         `json:"status"             orm:"status"               description:"状态。1=未付；2=成功；3=失败"` // 状态。1=未付；2=成功；3=失败
	AdminId            int         `json:"adminId"            orm:"admin_id"             description:"后台管理员ID。默认为0"`      // 后台管理员ID。默认为0
	Remark             string      `json:"remark"             orm:"remark"               description:"备注"`                // 备注
	Channel            string      `json:"channel"            orm:"channel"              description:"支付渠道"`              // 支付渠道
	CardNo             string      `json:"cardNo"             orm:"card_no"              description:"支付卡号"`              // 支付卡号
	ImageUrl           string      `json:"imageUrl"           orm:"image_url"            description:"渠道引导图片"`            // 渠道引导图片
	SysTradeNo         string      `json:"sysTradeNo"         orm:"sys_trade_no"         description:"支付系统订单流水号"`         // 支付系统订单流水号
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:""`                  //
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:""`                  //
	Currency           string      `json:"currency"           orm:"currency"             description:"币种"`                // 币种
}
