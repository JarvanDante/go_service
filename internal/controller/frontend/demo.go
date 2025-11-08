package frontend

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"go-service/api/frontendRoute"
	"golang.org/x/net/context"
)

type DemoController struct{}

func NewDemo() *DemoController {
	return &DemoController{}
}

func (c *DemoController) Index(ctx context.Context, req *frontendRoute.DemoReq) (res *frontendRoute.DemoRes, err error) {

	fmt.Println(g.Model("site"))
	site := g.Model("site")
	siteOne, err := site.One()
	//siteJson, _ := json.Marshal(siteOne) // 把 map 转成 JSON 字符串
	//msg := string(siteJson)
	////if err == nil {
	//	req.R
	//}

	//msg, err := frontendRoute.Demo().Index(ctx)
	if err != nil {
		return &frontendRoute.DemoRes{Message: ""}, nil
	}
	return &frontendRoute.DemoRes{Message: siteOne["name"].String()}, nil
}
