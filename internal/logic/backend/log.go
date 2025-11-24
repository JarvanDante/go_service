package backend

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"go-service/api/backendRoute"
	daosite "go-service/internal/dao/jh_site/dao"
	entitysite "go-service/internal/dao/jh_site/model/entity"
	"go-service/internal/dao/jinhuang/model/entity"
	"go-service/internal/service/backend"
)

type sLog struct {
}

func init() {
	backend.RegisterLog(&sLog{})
}

//LAdminLogs
/**
 * @desc：获取后台操作日志列表
 * @param ctx
 * @param req
 * @return res
 * @return err
 * @author : Carson
 */
func (s *sLog) LAdminLogs(ctx context.Context, req *backendRoute.AdminLogsReq) (interface{}, error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	where := map[string]interface{}{}
	where["site_id"] = siteId

	// 用户名筛选
	if req.Username != nil && *req.Username != "" {
		where["admin_username"] = *req.Username
	}

	m := daosite.AdminLog.Ctx(ctx).Where(where)

	// 时间范围筛选
	if req.Start != nil && *req.Start != "" && req.End != nil && *req.End != "" {
		m = m.WhereBetween("created_at", *req.Start, *req.End)
	}

	// 获取总数
	total, _ := m.Count()

	// 分页
	page := *req.Page
	size := *req.Size

	var list []*entitysite.AdminLog
	err := m.Fields("admin_username", "ip", "remark", "created_at").
		OrderDesc("created_at").
		Page(page, size).
		Scan(&list)

	if err != nil {
		return nil, err
	}

	// 构建返回数据
	var result []map[string]interface{}
	for _, log := range list {
		result = append(result, g.Map{
			"username":   log.AdminUsername,
			"ip":         log.Ip,
			"remark":     log.Remark,
			"created_at": log.CreatedAt,
		})
	}

	return g.Map{
		"list":  result,
		"count": total,
	}, nil
}
