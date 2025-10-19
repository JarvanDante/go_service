// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Favorites is the golang structure for table favorites.
type Favorites struct {
	Id         int         `json:"id"         orm:"id"          description:"ID"`      // ID
	SiteId     int         `json:"siteId"     orm:"site_id"     description:"商户id"`    // 商户id
	PlatformId int         `json:"platformId" orm:"platform_id" description:"游戏平台id"`  // 游戏平台id
	UserId     int         `json:"userId"     orm:"user_id"     description:"用户id"`    // 用户id
	GameId     int         `json:"gameId"     orm:"game_id"     description:"收藏的游戏id"` // 收藏的游戏id
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:""`        //
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:""`        //
}
