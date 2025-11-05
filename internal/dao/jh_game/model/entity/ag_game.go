// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AgGame is the golang structure for table ag_game.
type AgGame struct {
	Id          int         `json:"id"           orm:"id"           description:"自增长id"`                 // 自增长id
	Username    string      `json:"username"     orm:"username"     description:"下注会员"`                  // 下注会员
	GameNo      string      `json:"game_no"      orm:"game_no"      description:"唯一单号"`                  // 唯一单号
	GameName    string      `json:"game_name"    orm:"game_name"    description:"游戏名"`                   // 游戏名
	TableNo     string      `json:"table_no"     orm:"table_no"     description:"游戏桌号"`                  // 游戏桌号
	InningNo    string      `json:"inning_no"    orm:"inning_no"    description:"游戏局号"`                  // 游戏局号
	BetAmount   float64     `json:"bet_amount"   orm:"bet_amount"   description:"下注金额"`                  // 下注金额
	ValidAmount float64     `json:"valid_amount" orm:"valid_amount" description:"有效金额"`                  // 有效金额
	PlayType    string      `json:"play_type"    orm:"play_type"    description:"游戏玩法"`                  // 游戏玩法
	BetContent  string      `json:"bet_content"  orm:"bet_content"  description:"游戏内容"`                  // 游戏内容
	GameResult  string      `json:"game_result"  orm:"game_result"  description:"游戏结果"`                  // 游戏结果
	BetTime     *gtime.Time `json:"bet_time"     orm:"bet_time"     description:"下注时间"`                  // 下注时间
	Win         float64     `json:"win"          orm:"win"          description:"输赢金额、派彩金额、退还金额"`        // 输赢金额、派彩金额、退还金额
	OrderStatus int         `json:"order_status" orm:"order_status" description:"注单结算状态"`                // 注单结算状态
	GameType    string      `json:"game_type"    orm:"game_type"    description:"游戏类型"`                  // 游戏类型
	GameHall    string      `json:"game_hall"    orm:"game_hall"    description:"所属厅"`                   // 所属厅
	Devicetype  int         `json:"devicetype"   orm:"devicetype"   description:"设备类型(pc=0 mobile=1)"`   // 设备类型(pc=0 mobile=1)
	BetStatus   int         `json:"bet_status"   orm:"bet_status"   description:"注单状态；0=未结算；1=已结算；默认=0"` // 注单状态；0=未结算；1=已结算；默认=0
}
