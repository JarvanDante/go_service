// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Game is the golang structure for table game.
type Game struct {
	Id        uint        `json:"id"         orm:"id"         description:""`                         //
	Type      int         `json:"type"       orm:"type"       description:"类型。1=体育；2=彩票；3=真人；4=电子游戏"` // 类型。1=体育；2=彩票；3=真人；4=电子游戏
	Name      string      `json:"name"       orm:"name"       description:"游戏名称"`                     // 游戏名称
	Platform  string      `json:"platform"   orm:"platform"   description:"平台标识"`                     // 平台标识
	Code      string      `json:"code"       orm:"code"       description:"游戏标识"`                     // 游戏标识
	ImageCode string      `json:"image_code" orm:"image_code" description:"图片标识"`                     // 图片标识
	Status    int         `json:"status"     orm:"status"     description:"状态。1=启用；0=禁用"`             // 状态。1=启用；0=禁用
	DbName    string      `json:"db_name"    orm:"db_name"    description:"数据表名称"`                    // 数据表名称
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`                       // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`                         //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`                         //
}
