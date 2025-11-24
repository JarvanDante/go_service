package backendRoute

import "github.com/gogf/gf/v2/frame/g"

type UserLevelsReq struct {
	g.Meta `path:"/user-levels" method:"get" summary:"获取会员层级列表"`
}
