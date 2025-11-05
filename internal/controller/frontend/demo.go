package frontend

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	frontend2 "go-service/api/frontend"
	"golang.org/x/net/context"
)

type DemoController struct{}

func NewDemo() *DemoController {
	return &DemoController{}
}

func (c *DemoController) Index(ctx context.Context, req *frontend2.DemoReq) (res *frontend2.DemoRes, err error) {

	fmt.Println(g.Model("site"))
	site := g.Model("site")
	siteOne, err := site.One()
	//siteJson, _ := json.Marshal(siteOne) // 把 map 转成 JSON 字符串
	//msg := string(siteJson)
	////if err == nil {
	//	req.R
	//}

	//msg, err := frontend.Demo().Index(ctx)
	if err != nil {
		return &frontend2.DemoRes{Message: ""}, nil
	}
	return &frontend2.DemoRes{Message: siteOne["name"].String()}, nil
}
