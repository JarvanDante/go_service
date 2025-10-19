// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GameList is the golang structure for table game_list.
type GameList struct {
	Id        int         `json:"id"        orm:"id"         description:"ID"`       // ID
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"商户id"`     // 商户id
	Platform  string      `json:"platform"  orm:"platform"   description:"平台"`       // 平台
	Name      string      `json:"name"      orm:"name"       description:"游戏名称"`     // 游戏名称
	AliasName string      `json:"aliasName" orm:"alias_name" description:"别名"`       // 别名
	Style     string      `json:"style"     orm:"style"      description:"游戏类型"`     // 游戏类型
	Img       string      `json:"img"       orm:"img"        description:"游戏图片"`     // 游戏图片
	Img1      string      `json:"img1"      orm:"img1"       description:"游戏图片1"`    // 游戏图片1
	Device    string      `json:"device"    orm:"device"     description:"设备"`       // 设备
	Vendor    string      `json:"vendor"    orm:"vendor"     description:"提供方"`      // 提供方
	Code      string      `json:"code"      orm:"code"       description:"游戏代码"`     // 游戏代码
	Language  string      `json:"language"  orm:"language"   description:"语言"`       // 语言
	Hot       int         `json:"hot"       orm:"hot"        description:"热门"`       // 热门
	Recommend int         `json:"recommend" orm:"recommend"  description:"推荐"`       // 推荐
	LikeNum   int         `json:"likeNum"   orm:"like_num"   description:"喜欢数量"`     // 喜欢数量
	IsNew     int         `json:"isNew"     orm:"is_new"     description:"新的"`       // 新的
	Sort      int         `json:"sort"      orm:"sort"       description:"排序大的靠前"`   // 排序大的靠前
	Status    int         `json:"status"    orm:"status"     description:"状态0关闭1开启"` // 状态0关闭1开启
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"新增时间"`     // 新增时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`     // 更新时间
	GameId    int         `json:"gameId"    orm:"game_id"    description:""`         //
	Currency  string      `json:"currency"  orm:"currency"   description:"币种"`       // 币种
}
