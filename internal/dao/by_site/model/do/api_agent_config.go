// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ApiAgentConfig is the golang structure of table api_agent_config for DAO operations like Where/Data.
type ApiAgentConfig struct {
	g.Meta        `orm:"table:api_agent_config, do:true"`
	Id            any         //
	AgentName     any         //
	AppId         any         //
	AppKey        any         //
	PayloadSecret any         //
	ApiUrl        any         //
	CreatedAt     *gtime.Time //
	UpdatedAt     *gtime.Time //
	GameId        any         // 厅ID
	Currency      any         // 币种
}
