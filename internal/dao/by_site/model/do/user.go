// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta            `orm:"table:user, do:true"`
	Id                any         //
	SiteId            any         //
	GradeId           any         //
	LevelId           any         //
	AgentId           any         //
	ChannelId         any         // 渠道id
	Username          any         //
	Password          any         //
	PayPassword       any         //
	Status            any         //
	RegisterIp        any         //
	RegisterTime      *gtime.Time //
	RegisterUrl       any         // 注册来源
	RegisterDevice    any         // 1;电脑2; 手机3;平台
	DeviceCode        any         // 设备码
	LastLoginIp       any         //
	LastLoginTime     *gtime.Time //
	LastLoginAddress  any         // 最后登录地址
	Realname          any         //
	Mobile            any         //
	Email             any         // 邮箱
	Qq                any         // QQ号
	Birthday          *gtime.Time //
	Sex               any         // 性别。0=未知；1=男；2=女
	IsOnline          any         //
	FocusLevel        any         // 会员关注级别。1=正常；2=可疑；3=危险
	BalanceStatus     any         // 1=0=
	SafeQuestion      any         // 密保问题
	SafeAnswer        any         // 密保答案
	ShowBeginnerGuide any         // 是否显示新手引导。1=显示；0=不显示
	DeleteAt          any         // 是否删除。0=未删除；其他为删除时间戳
	Remark            any         // 备注
	SiteCode          any         // 应用码
	ChannelCode       any         // 渠道码
	CreatedAt         *gtime.Time //
	UpdatedAt         *gtime.Time //
	PayTimes          any         // 充值次数
	HasPass           any         // 是否设置密码
	Currency          any         // 用户币种
	WithdrawTimes     any         // 提现次数
}
