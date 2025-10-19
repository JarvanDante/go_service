// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Favorites is the golang structure of table favorites for DAO operations like Where/Data.
type Favorites struct {
	g.Meta     `orm:"table:favorites, do:true"`
	Id         any         // ID
	SiteId     any         // 商户id
	PlatformId any         // 游戏平台id
	UserId     any         // 用户id
	GameId     any         // 收藏的游戏id
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
