// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SiteApkSourceDao is the data access object for the table site_apk_source.
type SiteApkSourceDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SiteApkSourceColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SiteApkSourceColumns defines and stores column names for the table site_apk_source.
type SiteApkSourceColumns struct {
	Id           string // ID
	SiteId       string // 商户id
	ChannelId    string // 渠道id
	ApkName      string // apk包名
	Version      string // 版本号
	DownloadUrl  string // 下载地址
	Status       string // 状态1开0关
	ForceUpdate  string // 是否强更1是0否
	UploadStatus string // 上传状态1成功2失败
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
}

// siteApkSourceColumns holds the columns for the table site_apk_source.
var siteApkSourceColumns = SiteApkSourceColumns{
	Id:           "id",
	SiteId:       "site_id",
	ChannelId:    "channel_id",
	ApkName:      "apk_name",
	Version:      "version",
	DownloadUrl:  "download_url",
	Status:       "status",
	ForceUpdate:  "force_update",
	UploadStatus: "upload_status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewSiteApkSourceDao creates and returns a new DAO object for table data access.
func NewSiteApkSourceDao(handlers ...gdb.ModelHandler) *SiteApkSourceDao {
	return &SiteApkSourceDao{
		group:    "by_site",
		table:    "site_apk_source",
		columns:  siteApkSourceColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SiteApkSourceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SiteApkSourceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SiteApkSourceDao) Columns() SiteApkSourceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SiteApkSourceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SiteApkSourceDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SiteApkSourceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
