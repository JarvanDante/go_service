// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteApkSource is the golang structure for table site_apk_source.
type SiteApkSource struct {
	Id           int         `json:"id"           orm:"id"            description:"ID"`         // ID
	SiteId       int         `json:"siteId"       orm:"site_id"       description:"商户id"`       // 商户id
	ChannelId    int         `json:"channelId"    orm:"channel_id"    description:"渠道id"`       // 渠道id
	ApkName      string      `json:"apkName"      orm:"apk_name"      description:"apk包名"`      // apk包名
	Version      string      `json:"version"      orm:"version"       description:"版本号"`        // 版本号
	DownloadUrl  string      `json:"downloadUrl"  orm:"download_url"  description:"下载地址"`       // 下载地址
	Status       int         `json:"status"       orm:"status"        description:"状态1开0关"`     // 状态1开0关
	ForceUpdate  int         `json:"forceUpdate"  orm:"force_update"  description:"是否强更1是0否"`   // 是否强更1是0否
	UploadStatus int         `json:"uploadStatus" orm:"upload_status" description:"上传状态1成功2失败"` // 上传状态1成功2失败
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:""`           //
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:""`           //
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:""`           //
}
