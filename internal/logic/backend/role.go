package backend

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"go-service/api/backendRoute"
	"go-service/internal/service/backend"
)

type sRole struct {
}

func init() {
	backend.RegisterRole(&sRole{})
}

func (s *sRole) LIndex(ctx context.Context, req *backendRoute.RoleReq) (string, error) {

	// 从 context 中获取 admin 信息
	adminValue := g.RequestFromCtx(ctx).GetCtxVar("admin")
	if adminValue.IsNil() {
		return "", gerror.New("未找到用户信息")
	}

	adminInfo := adminValue.Map()

	// 使用管理员信息
	userID := adminInfo["id"]
	username := adminInfo["username"]
	siteID := adminInfo["site_id"]

	fmt.Printf("当前用户: %s (ID: %v, SiteID: %v)\n", username, userID, siteID)

	return "操作成功", nil

}
