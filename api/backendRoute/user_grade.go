package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type UserGradesReq struct {
	g.Meta `path:"/user-grades" method:"get" summary:"获取会员等级列表"`
}
