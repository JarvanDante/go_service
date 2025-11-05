// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// GameConfig is the golang structure of table game_config for DAO operations like Where/Data.
type GameConfig struct {
	g.Meta    `orm:"table:game_config, do:true"`
	Id        any         // ID
	SiteId    any         // 应用id
	ChannelId any         // 渠道id
	GameId    any         // 游戏id
	Vendor    any         // 游戏厂商
	Status    any         // 状态,1开0关
	Available any         // 可用状态1可用2维护
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
