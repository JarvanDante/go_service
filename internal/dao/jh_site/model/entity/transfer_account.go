// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TransferAccount is the golang structure for table transfer_account.
type TransferAccount struct {
	Id          uint        `json:"id"           orm:"id"           description:""`               //
	SiteId      int         `json:"site_id"      orm:"site_id"      description:"站点ID"`           // 站点ID
	BankType    int         `json:"bank_type"    orm:"bank_type"    description:"银行类型"`           // 银行类型
	BankName    string      `json:"bank_name"    orm:"bank_name"    description:"银行名称"`           // 银行名称
	Name        string      `json:"name"         orm:"name"         description:"转账接口名称"`         // 转账接口名称
	BankUrl     string      `json:"bank_url"     orm:"bank_url"     description:"银行链接"`           // 银行链接
	Qrcode      string      `json:"qrcode"       orm:"qrcode"       description:"二维码图片地址"`        // 二维码图片地址
	CardAccount string      `json:"card_account" orm:"card_account" description:"银行户名或者第三方收款人"`   // 银行户名或者第三方收款人
	CardNo      string      `json:"card_no"      orm:"card_no"      description:"银行卡号或者第三方账号"`    // 银行卡号或者第三方账号
	DepositBank string      `json:"deposit_bank" orm:"deposit_bank" description:"开户行"`            // 开户行
	EachMin     float64     `json:"each_min"     orm:"each_min"     description:"单笔最低。默认为1元"`     // 单笔最低。默认为1元
	EachMax     float64     `json:"each_max"     orm:"each_max"     description:"单笔最高。默认为5000元"`  // 单笔最高。默认为5000元
	DailyMax    float64     `json:"daily_max"    orm:"daily_max"    description:"单日上限。默认为0。0=不限"` // 单日上限。默认为0。0=不限
	TodayCount  int         `json:"today_count"  orm:"today_count"  description:"今日入款次数"`         // 今日入款次数
	TodayAmount float64     `json:"today_amount" orm:"today_amount" description:"今日转账总额"`         // 今日转账总额
	Status      int         `json:"status"       orm:"status"       description:"状态。1=可用；0=禁用"`   // 状态。1=可用；0=禁用
	Sort        int         `json:"sort"         orm:"sort"         description:"排序。值越小排名越靠前"`    // 排序。值越小排名越靠前
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   description:""`               //
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   description:""`               //
	Remark      string      `json:"remark"       orm:"remark"       description:""`               //
}
