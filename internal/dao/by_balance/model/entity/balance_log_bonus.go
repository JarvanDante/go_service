// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogBonus is the golang structure for table balance_log_bonus.
type BalanceLogBonus struct {
	Id                 uint64      `json:"id"                 orm:"id"                   description:""`                                      //
	SiteId             int         `json:"siteId"             orm:"site_id"              description:"站点ID"`                                  // 站点ID
	ChannelId          int         `json:"channelId"          orm:"channel_id"           description:"渠道id"`                                  // 渠道id
	TradeType          int         `json:"tradeType"          orm:"trade_type"           description:"交易类型。1=在线充值红利；2=转账汇款红利；3=生日红利；4=签到红利；"` // 交易类型。1=在线充值红利；2=转账汇款红利；3=生日红利；4=签到红利；
	UserId             int         `json:"userId"             orm:"user_id"              description:"会员ID"`                                  // 会员ID
	Username           string      `json:"username"           orm:"username"             description:"会员用户名"`                                 // 会员用户名
	TradeNo            string      `json:"tradeNo"            orm:"trade_no"             description:"流水号"`                                   // 流水号
	ActivityRechargeId int         `json:"activityRechargeId" orm:"activity_recharge_id" description:"充值活动ID"`                                // 充值活动ID
	Money              float64     `json:"money"              orm:"money"                description:"金额"`                                    // 金额
	Fee                float64     `json:"fee"                orm:"fee"                  description:"手续费"`                                   // 手续费
	Status             int         `json:"status"             orm:"status"               description:"状态。1=待处理；2=成功；3=失败"`                    // 状态。1=待处理；2=成功；3=失败
	AdminId            int         `json:"adminId"            orm:"admin_id"             description:"后台管理员ID。默认为0"`                          // 后台管理员ID。默认为0
	Remark             string      `json:"remark"             orm:"remark"               description:"备注"`                                    // 备注
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"           description:""`                                      //
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"           description:""`                                      //
}
