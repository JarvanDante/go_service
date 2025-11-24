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

type sUser struct {
}

func init() {
	backend.RegisterUser(&sUser{})
}

//LIndex
/**
 * @desc：获取会员列表
 * @param ctx
 * @param req
 * @return res
 * @return err
 * @author : Carson
 */
func (s *sUser) LIndex(ctx context.Context, req *backendRoute.UsersReq) (interface{}, error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 构建查询条件
	where := g.Map{}
	where["site_id"] = siteId

	// 会员等级
	if req.Grade != nil && *req.Grade > 0 {
		where["grade_id"] = *req.Grade
	}

	// 会员层级
	if req.Level != nil && *req.Level > 0 {
		where["level_id"] = *req.Level
	}

	// 状态
	if req.Status != nil {
		where["status"] = *req.Status
	}

	// 用户名
	if req.Username != nil && *req.Username != "" {
		where["username like ?"] = "%" + *req.Username + "%"
	}

	// 真实姓名
	if req.Realname != nil && *req.Realname != "" {
		where["realname like ?"] = "%" + *req.Realname + "%"
	}

	// 手机号
	if req.Mobile != nil && *req.Mobile != "" {
		where["mobile like ?"] = "%" + *req.Mobile + "%"
	}

	// 银行卡号
	if req.CardNo != nil && *req.CardNo != "" {
		where["card_no like ?"] = "%" + *req.CardNo + "%"
	}

	// 代理用户名（需要关联查询 agent 表）
	if req.Agent != nil && *req.Agent != "" {
		// 先查询代理ID
		var agent *entitysite.Agent
		err := daosite.Agent.Ctx(ctx).
			Where("site_id", siteId).
			Where("username", *req.Agent).
			Fields("id").
			Scan(&agent)

		if err != nil {
			return g.Map{
				"list":               []interface{}{},
				"count":              0,
				"total_users":        0,
				"total_charge_users": 0,
			}, nil
		}

		// 如果找不到代理，返回空列表
		if agent == nil || agent.Id == 0 {
			return g.Map{
				"list":               []interface{}{},
				"count":              0,
				"total_users":        0,
				"total_charge_users": 0,
			}, nil
		}

		// 添加代理ID筛选条件
		where["agent_id"] = agent.Id
	}

	// 域名
	if req.Domain != nil && *req.Domain != "" {
		where["domain"] = *req.Domain
	}

	// 注册时间范围
	// TODO: 实现时间范围筛选

	// 是否首存
	// TODO: 实现首存筛选

	// 查询数据库（这里需要根据实际的 DAO 实现）
	// 由于没有看到具体的 User DAO，这里提供一个基础框架

	page := 1
	if req.Page != nil {
		page = *req.Page
	}

	size := 20
	if req.Num != nil {
		size = *req.Num
	}

	// TODO: 实现实际的数据库查询
	// 这里返回一个示例结构
	return g.Map{
		"list":               []interface{}{},
		"count":              0,
		"total_users":        0,
		"total_charge_users": 0,
	}, nil
}
