// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLoginLog is the golang structure for table user_login_log.
type UserLoginLog struct {
	Id           uint        `json:"id"            orm:"id"            description:""`                   //
	SiteId       int         `json:"site_id"       orm:"site_id"       description:""`                   //
	UserId       int         `json:"user_id"       orm:"user_id"       description:"会员ID"`               // 会员ID
	Username     string      `json:"username"      orm:"username"      description:"会员用户名"`              // 会员用户名
	RefererUrl   string      `json:"referer_url"   orm:"referer_url"   description:"来源网址"`               // 来源网址
	LoginUrl     string      `json:"login_url"     orm:"login_url"     description:"登录网址"`               // 登录网址
	LoginTime    *gtime.Time `json:"login_time"    orm:"login_time"    description:"登录时间"`               // 登录时间
	LoginIp      string      `json:"login_ip"      orm:"login_ip"      description:"登录IP"`               // 登录IP
	LoginAddress string      `json:"login_address" orm:"login_address" description:"登录地址"`               // 登录地址
	Os           string      `json:"os"            orm:"os"            description:"操作系统"`               // 操作系统
	Network      string      `json:"network"       orm:"network"       description:"网络"`                 // 网络
	Screen       string      `json:"screen"        orm:"screen"        description:"分辨率"`                // 分辨率
	Browser      string      `json:"browser"       orm:"browser"       description:"浏览器"`                // 浏览器
	Device       int         `json:"device"        orm:"device"        description:"终端。1=电脑；2=手机；3=平板"`  // 终端。1=电脑；2=手机；3=平板
	IsRobot      int         `json:"is_robot"      orm:"is_robot"      description:"判断是否是机器人登录。1=是；0=否"` // 判断是否是机器人登录。1=是；0=否
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:"创建时间"`               // 创建时间
}
