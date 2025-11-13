package global

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"go-service/utility/helpers"
)

type Global struct {
	Ctx    context.Context
	Config *gcfg.Config
	Log    *glog.Logger
	DB     *gdb.DB
	Cache  *gcache.Cache
	Server *ghttp.Server
}

var G *Global

var (
	SiteCode string
	Site     interface{}
)

// Init 初始化全局变量
func Init() error {
	G = &Global{
		Ctx:    gctx.New(),
		Config: g.Cfg(),
		Log:    g.Log(),
		Cache:  gcache.New(),
	}

	// 初始化服务器
	G.Server = g.Server()

	// 定义全局变量
	SiteCode = helpers.GetEnv("SITE_CODE", "jh")

	//获取当前站点信息
	//Site, _ = dao.GetSiteObject()

	//设置日志格式为json
	glog.SetDefaultHandler(glog.HandlerJson)

	return nil
}

// GetDB 获取数据库实例
func GetDB() *gdb.DB {
	return G.DB
}

// GetConfig 获取配置实例
func GetConfig() *gcfg.Config {
	return G.Config
}
