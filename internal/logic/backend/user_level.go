package backend

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
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

// LIndex 获取会员层级列表
func (s *sUserLevel) LIndex(ctx context.Context, req *backendRoute.UserLevelsReq) (interface{}, error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 查询会员层级列表
	var levels []entitysite.UserLevel
	err := daosite.UserLevel.Ctx(ctx).
		Where("site_id", siteId).
		Where("status", 1).
		Order("id ASC").
		Scan(&levels)

	if err != nil {
		return nil, err
	}

	// 获取所有返水规则
	rebateRules, _ := s.getRebateRuleOptions(ctx, gconv.Int(siteId))

	// 组装返回数据
	data := make([]g.Map, 0)
	for i, level := range levels {
		item := g.Map{
			"id":                   level.Id,
			"name":                 level.Name,
			"is_rebate":            level.IsRebate,
			"rebate_rule_name":     s.getRebateRuleName(rebateRules, level.IsRebate, level.RebateRuleId),
			"daily_withdraw_times": level.DailyWithdrawTimes,
			"is_default":           0,
		}

		// 第一个为默认层级
		if i == 0 {
			item["is_default"] = 1
		}

		// 获取该层级的会员数量
		userCount, _ := daosite.User.Ctx(ctx).
			Where("site_id", siteId).
			Where("level_id", level.Id).
			Count()
		item["user_count"] = userCount

		data = append(data, item)
	}

	return data, nil
}

// LCreateUserLevel 新增会员层级
func (s *sUserLevel) LCreateUserLevel(ctx context.Context, req *backendRoute.CreateUserLevelReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 检查层级名称是否已存在
	count, err := daosite.UserLevel.Ctx(ctx).
		Where("site_id", siteId).
		Where("name", req.Name).
		Count()

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("会员层级名称已经存在")
	}

	// 开启事务
	err = daosite.UserLevel.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		now := gtime.Now()

		// 添加会员层级
		levelData := g.Map{
			"site_id":              siteId,
			"name":                 req.Name,
			"is_rebate":            req.IsRebate,
			"rebate_rule_id":       req.RebateRuleId,
			"daily_withdraw_times": req.DailyWithdrawTimes,
			"status":               1,
			"created_at":           now,
			"updated_at":           now,
		}

		_, err := tx.Model("user_level").Data(levelData).Insert()
		if err != nil {
			return err
		}

		// TODO: 保存转账接口列表和支付接口列表
		// 这部分需要根据实际的表结构来实现

		return nil
	})

	return err
}

// LGetUpdateUserLevel 获取会员层级信息用于编辑
func (s *sUserLevel) LGetUpdateUserLevel(ctx context.Context, req *backendRoute.GetUpdateUserLevelReq) (interface{}, error) {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 查询层级信息
	var level entitysite.UserLevel
	err := daosite.UserLevel.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Scan(&level)

	if err != nil || level.Id == 0 {
		return nil, errors.New("没有找到用户层级")
	}

	// 获取返水规则列表
	rebateRuleList, _ := s.getRebateRuleOptions(ctx, gconv.Int(siteId))

	data := g.Map{
		"id":                   level.Id,
		"name":                 level.Name,
		"is_rebate":            level.IsRebate,
		"rebate_rule_id":       level.RebateRuleId,
		"rebate_rule_list":     rebateRuleList,
		"daily_withdraw_times": level.DailyWithdrawTimes,
		// TODO: 获取转账接口列表和支付接口列表
		"transfer_account_list": []interface{}{},
		"payment_account_list":  []interface{}{},
	}

	return data, nil
}

// LUpdateUserLevel 编辑会员层级
func (s *sUserLevel) LUpdateUserLevel(ctx context.Context, req *backendRoute.UpdateUserLevelReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 检查层级是否存在
	var level entitysite.UserLevel
	err := daosite.UserLevel.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Scan(&level)

	if err != nil || level.Id == 0 {
		return errors.New("没有找到用户层级")
	}

	// 开启事务
	err = daosite.UserLevel.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 更新层级
		_, err := tx.Model("user_level").
			Where("id", req.Id).
			Where("site_id", siteId).
			Data(g.Map{
				"name":                 req.Name,
				"is_rebate":            req.IsRebate,
				"rebate_rule_id":       req.RebateRuleId,
				"daily_withdraw_times": req.DailyWithdrawTimes,
				"updated_at":           gtime.Now(),
			}).
			Update()

		if err != nil {
			return err
		}

		// TODO: 更新转账接口列表和支付接口列表

		return nil
	})

	return err
}

// LDeleteUserLevel 删除会员层级
func (s *sUserLevel) LDeleteUserLevel(ctx context.Context, req *backendRoute.DeleteUserLevelReq) error {
	site := g.RequestFromCtx(ctx).GetCtxVar("site").Val().(*entity.Site)
	siteId := site.Id

	// 查询层级信息
	var level entitysite.UserLevel
	err := daosite.UserLevel.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Scan(&level)

	if err != nil || level.Id == 0 {
		return errors.New("没有找到用户层级")
	}

	// 检查是否有会员使用该层级
	userCount, err := daosite.User.Ctx(ctx).
		Where("site_id", siteId).
		Where("level_id", req.Id).
		Count()

	if err != nil {
		return err
	}

	if userCount > 0 {
		return errors.New("该层级下面有会员，请先将会员变更为其它层级，才能删除该层级")
	}

	// 删除层级
	_, err = daosite.UserLevel.Ctx(ctx).
		Where("id", req.Id).
		Where("site_id", siteId).
		Delete()

	if err != nil {
		return errors.New("删除失败")
	}

	return nil
}

// getRebateRuleOptions 获取返水规则选项列表
func (s *sUserLevel) getRebateRuleOptions(ctx context.Context, siteId int) (map[int]string, error) {
	var rules []entitysite.RebateRule
	err := daosite.RebateRule.Ctx(ctx).
		Where("site_id", siteId).
		Where("status", 1).
		Fields("id", "name").
		Scan(&rules)

	if err != nil {
		return nil, err
	}

	ruleMap := make(map[int]string)
	for _, rule := range rules {
		ruleMap[int(rule.Id)] = rule.Name
	}

	return ruleMap, nil
}

// getRebateRuleName 获取返水规则名称
func (s *sUserLevel) getRebateRuleName(rebateRules map[int]string, isRebate int, ruleId int) string {
	if isRebate == 0 {
		return "关闭"
	}

	if name, ok := rebateRules[ruleId]; ok {
		return name
	}

	return "未设置"
}
