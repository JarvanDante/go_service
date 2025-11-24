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

type sUserLevel struct {
}

func init() {
	backend.RegisterUserLevel(&sUserLevel{})
}

// LIndex
/**
 * @desc：获取会员层级列表
 * @param ctx
 * @param req
 * @return res
 * @return err
 * @author : Carson
 */
func (s *sUserLevel) LIndex(ctx context.Context, req *backendRoute.UserLevelsReq) (interface{}, error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 查询会员层级列表
	var levels []entitysite.UserLevel
	err := daosite.UserLevel.Ctx(ctx).
		Where("site_id", siteId).
		Where("status", 1). // STATUS_ACTIVE = 1
		Order("id ASC").
		Scan(&levels)

	if err != nil {
		return nil, err
	}

	// 组装返回数据
	data := make([]g.Map, 0)
	for _, level := range levels {
		// 获取该层级的会员数量
		userCount, _ := daosite.User.Ctx(ctx).
			Where("site_id", siteId).
			Where("level_id", level.Id).
			Count()

		item := g.Map{
			"id":                   level.Id,
			"name":                 level.Name,
			"is_rebate":            level.IsRebate,
			"rebate_rule_id":       level.RebateRuleId,
			"daily_withdraw_times": level.DailyWithdrawTimes,
			"login_url":            level.LoginUrl,
			"status":               level.Status,
			"user_count":           userCount,
		}

		data = append(data, item)
	}

	return data, nil
}
