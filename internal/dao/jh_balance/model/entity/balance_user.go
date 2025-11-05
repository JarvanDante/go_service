// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceUser is the golang structure for table balance_user.
type BalanceUser struct {
	Id            uint        `json:"id"             orm:"id"             description:""`                                                                                                                     //
	SiteId        int         `json:"site_id"        orm:"site_id"        description:"站点ID"`                                                                                                                 // 站点ID
	UserId        int         `json:"user_id"        orm:"user_id"        description:"会员ID"`                                                                                                                 // 会员ID
	Balance       float64     `json:"balance"        orm:"balance"        description:"可用金额"`                                                                                                                 // 可用金额
	BalanceFrozen float64     `json:"balance_frozen" orm:"balance_frozen" description:"冻结金额"`                                                                                                                 // 冻结金额
	BalanceDebt   float64     `json:"balance_debt"   orm:"balance_debt"   description:"负债金额。防止账户余额出现负数的情况。例如：注单结算错误，需要扣掉会员账户金额。如果会员账户金额足够，则直接扣除会员账户金额；如果会员账户金额不足，则不扣除会员账户金额，直接存入负债金额。在下次会员有资金变动的时候，先来计算这笔金额"` // 负债金额。防止账户余额出现负数的情况。例如：注单结算错误，需要扣掉会员账户金额。如果会员账户金额足够，则直接扣除会员账户金额；如果会员账户金额不足，则不扣除会员账户金额，直接存入负债金额。在下次会员有资金变动的时候，先来计算这笔金额
	Points        int         `json:"points"         orm:"points"         description:"会员积分"`                                                                                                                 // 会员积分
	BalanceStatus int         `json:"balance_status" orm:"balance_status" description:"资金状态。1=正常；0=锁定"`                                                                                                       // 资金状态。1=正常；0=锁定
	CreatedAt     *gtime.Time `json:"created_at"     orm:"created_at"     description:""`                                                                                                                     //
	UpdatedAt     *gtime.Time `json:"updated_at"     orm:"updated_at"     description:""`                                                                                                                     //
}
