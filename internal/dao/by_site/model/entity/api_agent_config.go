// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ApiAgentConfig is the golang structure for table api_agent_config.
type ApiAgentConfig struct {
	Id            int         `json:"id"            orm:"id"             description:""`    //
	AgentName     string      `json:"agentName"     orm:"agent_name"     description:""`    //
	AppId         string      `json:"appId"         orm:"app_id"         description:""`    //
	AppKey        string      `json:"appKey"        orm:"app_key"        description:""`    //
	PayloadSecret string      `json:"payloadSecret" orm:"payload_secret" description:""`    //
	ApiUrl        string      `json:"apiUrl"        orm:"api_url"        description:""`    //
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:""`    //
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:""`    //
	GameId        int         `json:"gameId"        orm:"game_id"        description:"厅ID"` // 厅ID
	Currency      string      `json:"currency"      orm:"currency"       description:"币种"`  // 币种
}
