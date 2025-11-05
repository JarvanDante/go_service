// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GameConfig is the golang structure for table game_config.
type GameConfig struct {
	Id        int         `json:"id"        orm:"id"         description:"ID"`         // ID
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"应用id"`       // 应用id
	ChannelId int         `json:"channelId" orm:"channel_id" description:"渠道id"`       // 渠道id
	GameId    int         `json:"gameId"    orm:"game_id"    description:"游戏id"`       // 游戏id
	Vendor    string      `json:"vendor"    orm:"vendor"     description:"游戏厂商"`       // 游戏厂商
	Status    int         `json:"status"    orm:"status"     description:"状态,1开0关"`    // 状态,1开0关
	Available int         `json:"available" orm:"available"  description:"可用状态1可用2维护"` // 可用状态1可用2维护
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`       // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`       // 更新时间
}
