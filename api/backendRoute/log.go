package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type AdminLogsReq struct {
	g.Meta   `path:"/admin-logs" method:"get" summary:"获取后台操作日志列表"`
	Username *string `json:"username" dc:"用户名"`
	Start    *string `json:"start" dc:"开始时间"`
	End      *string `json:"end" dc:"结束时间"`
	Page     *int    `json:"page" d:"1" dc:"页数"`
	Size     *int    `json:"size" d:"50" dc:"页码"`
}
