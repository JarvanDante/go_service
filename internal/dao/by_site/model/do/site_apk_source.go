// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteApkSource is the golang structure of table site_apk_source for DAO operations like Where/Data.
type SiteApkSource struct {
	g.Meta       `orm:"table:site_apk_source, do:true"`
	Id           any         // ID
	SiteId       any         // 商户id
	ChannelId    any         // 渠道id
	ApkName      any         // apk包名
	Version      any         // 版本号
	DownloadUrl  any         // 下载地址
	Status       any         // 状态1开0关
	ForceUpdate  any         // 是否强更1是0否
	UploadStatus any         // 上传状态1成功2失败
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
	DeletedAt    *gtime.Time //
}
