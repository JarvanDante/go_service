package frontend

import "github.com/gogf/gf/v2/frame/g"

type DemoReq struct {
	g.Meta `path:"/" method:"get" tags:"Demo" summary:"根目前"`
}

type DemoRes struct {
	Message string `json:"message"`
}
