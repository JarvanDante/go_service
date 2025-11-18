package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type BasicSettingReq struct {
	g.Meta `path:"/basic-setting" method:"get" summary:"获取站点基本信息"`
}
